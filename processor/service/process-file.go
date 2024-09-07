package service

import (
	"context"
	"strconv"
	"time"

	"github.com/Jhooomn/file-transaction-processor/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	dataHeader = []string{"Id", "Date", "Transaction"}
)

func (ps *processorService) Execute() {
	ps.execute()
}

func (ps *processorService) execute() {

	// fetch different files names
	fileNames, err := utils.GetFileNames(ps.opts.dataPath)
	if err != nil {
		ps.logger.WithFields(logrus.Fields{
			"details": err,
			"path":    ps.opts.dataPath,
		}).Error("Failed to read data path")
		return
	}

	if len(fileNames) == 0 {
		ps.logger.WithFields(logrus.Fields{
			"details": err,
			"path":    ps.opts.dataPath,
		}).Error("Could not fetch files from data path")
		return
	}

	eg, _ := errgroup.WithContext(context.Background())
	eg.SetLimit(ps.opts.workerPool)

	for _, fileName := range fileNames {
		eg.Go(func() error {
			// Fetch data by path location
			data, err := utils.ReadCSV(fileName, dataHeader)
			if err != nil {
				ps.logger.WithFields(logrus.Fields{
					"details": err,
					"file":    fileName,
				}).Error("Failed to read csv file")
				return err
			}

			if len(data) == 0 {
				ps.logger.WithFields(logrus.Fields{
					"details": err,
					"file":    fileName,
				}).Error("No data found to process")
				return err
			}

			// process the data
			err = ps.process(data, fileName)
			if err != nil {
				ps.logger.WithFields(logrus.Fields{
					"details": err,
					"file":    fileName,
				}).Error("It was no possible to process the data")
			}
			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		ps.logger.WithFields(logrus.Fields{
			"details": err,
			"path":    ps.opts.dataPath,
		}).Error("There was a problem while processing files")
	}

}

func (ps *processorService) process(data []map[string]string, fileName string) error {
	// transform data
	transactions, err := ps.parseCSV(data)
	if err != nil {
		ps.logger.WithFields(logrus.Fields{
			"details": err,
			"file":    fileName,
		}).Error("It was no possible to parse the data")
		return err
	}

	if len(transactions) == 0 {
		ps.logger.WithFields(logrus.Fields{
			"details": nil,
			"file":    fileName,
		}).Error("No data found into this file")
		return nil // TODO: return business error
	}

	// invoke func with the logic
	summary := UserSummary{}
	summary.CalculateSummary(transactions)

	// persist in repository

	// send notifications concurrently

	return nil
}

func (ps *processorService) calculateTransactions(transactions []TransactionRecord) error {
	return nil
}

func (ps *processorService) parseCSV(data []map[string]string) ([]TransactionRecord, error) {
	var records []TransactionRecord

	for _, record := range data {
		id, err := strconv.Atoi(record[dataHeader[0]]) // ID
		if err != nil {
			return nil, err
		}

		date, err := time.Parse("1/2", record[dataHeader[1]]) // Date
		if err != nil {
			return nil, err
		}

		transaction, err := strconv.ParseFloat(record[dataHeader[2]], 64) // Transaction
		if err != nil {
			return nil, err
		}

		tRecord := TransactionRecord{
			Id:          id,
			Date:        date,
			Transaction: transaction,
		}

		tRecord.ParseTransactionType()

		records = append(records, tRecord)
	}

	return records, nil
}
