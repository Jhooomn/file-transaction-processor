package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Jhooomn/file-transaction-processor/utils"
	"golang.org/x/sync/errgroup"
)

var (
	dataHeader   = []string{"Id", "Date", "Transaction"}
	defaultName  = os.Getenv("DEFAULT_NAME")
	defaultEmail = os.Getenv("DEFAULT_EMAIL")
)

func (ps *processorService) Execute() {
	ps.execute()
}

func (ps *processorService) execute() {

	eg, ctx := errgroup.WithContext(context.Background())
	eg.SetLimit(ps.opts.workerPool)

	// fetch different files names
	fileNames, err := utils.GetFileNames(ps.opts.dataPath)
	if err != nil {
		ps.logger.Error(fmt.Sprintf("Failed to read data path: [%s]", err))
		return
	}

	if len(fileNames) == 0 {
		ps.logger.Error(fmt.Sprintf("Could not fetch files from data path: [%s]", err))
		return
	}

	for _, fileName := range fileNames {
		eg.Go(func() error {
			// Fetch data by path location
			data, err := utils.ReadCSV(fileName, dataHeader)
			if err != nil {
				ps.logger.Error(fmt.Sprintf("Failed to read csv file: [%s]", err))
				return err
			}

			if len(data) == 0 {
				ps.logger.Error(fmt.Sprintf("No data found to process: [%s]", err))
				return err
			}

			// process the data
			err = ps.process(ctx, data, fileName)
			if err != nil {
				ps.logger.Error(fmt.Sprintf("It was no possible to process the data: [%s]", err))
				return err
			}

			if ctx.Err() != nil {
				return ctx.Err()
			}

			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		ps.logger.Error(fmt.Sprintf("There was a problem while processing files: [%s]", err))
		return
	}

	ps.logger.Info("last line")

}

func (ps *processorService) process(ctx context.Context, data []map[string]string, fileName string) error {
	// transform data
	transactions, err := ps.parseCSV(data)
	if err != nil {
		ps.logger.Error(fmt.Sprintf("It was no possible to parse the data: [%s]", err))
		return err
	}

	if len(transactions) == 0 {
		ps.logger.Error(fmt.Sprintf("No data found into this file: [%s]", err))
		return nil // TODO: return business error
	}

	// invoke func with the logic
	summary := UserSummary{}
	summary.CalculateSummary(transactions)

	// persist in repository
	err = ps.processorRepository.Save(ctx,
		summary.TransactionsPerMonth,
		summary.TotalBalance,
		summary.AvgCredit,
		summary.AvgDebit,
		summary.CreditCount,
		summary.DebitCount,
		defaultName,
		defaultEmail)
	if err != nil {
		ps.logger.Error(fmt.Sprintf("It was no possible to save the record: [%s]", err))
		return err
	}

	ps.logger.Info(fmt.Sprintf("record has been saved into db [%s] - [%s]", defaultName, defaultEmail))

	// send notifications concurrently

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

// CalculateSummary calculates the summary of transactions for a user.
func (us *UserSummary) CalculateSummary(transactions []TransactionRecord) {
	us.TransactionsPerMonth = make(map[string]uint)

	for _, tr := range transactions {
		month := tr.Date.Format("January 2006")

		us.TransactionsPerMonth[month]++

		us.TotalBalance += tr.Transaction

		if tr.Type == Credit {
			us.AvgCredit += tr.Transaction // counter
			us.CreditCount++
		} else {
			us.AvgDebit += tr.Transaction // counter
			us.DebitCount++
		}
	}

	noMonths := float64(len(us.TransactionsPerMonth))

	us.AvgDebit = us.AvgDebit / noMonths
	us.AvgCredit = us.AvgCredit / noMonths
}

// ParseTransactionType determines if the transaction is a credit or debit.
func (tr *TransactionRecord) ParseTransactionType() {
	if tr.Transaction >= 0 {
		tr.Type = Credit
	} else {
		tr.Type = Debit
	}
}
