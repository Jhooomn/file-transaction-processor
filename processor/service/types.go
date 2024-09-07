package service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// UserSummary stores the latest summary of the user's transactions.
type UserSummary struct {
	Id                   primitive.ObjectID
	Contact              Contact
	TotalBalance         float64
	TransactionsPerMonth map[string]uint
	AvgCredit            float64
	AvgDebit             float64
	CreditCount          uint
	DebitCount           uint
}

// Contact store common info of the user
type Contact struct {
	Name  string
	Email string
}

// Historic contains logging values for the record
type Historic struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
}
