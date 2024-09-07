package service

import (
	"time"
)

// TransactionType represents the type of transaction.
type TransactionType string

const (
	Credit TransactionType = "Credit"
	Debit  TransactionType = "Debit"
)

// TransactionRecord represents a financial transaction.
type TransactionRecord struct {
	Id          int
	Date        time.Time
	Transaction float64
	Type        TransactionType
}

// ParseTransactionType determines if the transaction is a credit or debit.
func (tr *TransactionRecord) ParseTransactionType() {
	if tr.Transaction >= 0 {
		tr.Type = Credit
	} else {
		tr.Type = Debit
	}
}

// UserSummary stores the summary of the user's transactions.
type UserSummary struct {
	TotalBalance         float64
	TransactionsPerMonth map[string]uint
	AvgCredit            float64
	AvgDebit             float64
	CreditCount          uint
	DebitCount           uint
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
