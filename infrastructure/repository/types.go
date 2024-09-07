package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type TransactionSummary struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	TransactionsPerMonth map[string]uint    `bson:"transactionsPerMonth"`
	TotalBalance         float64            `bson:"totalBalance"`
	AvgCredit            map[string]float64 `bson:"avgCredit"`
	AvgDebit             map[string]float64 `bson:"avgDebit"`
	CreditCount          map[string]uint    `bson:"creditCount"`
	DebitCount           map[string]uint    `bson:"debitCount"`
	Name                 string             `bson:"name"`
	Email                string             `bson:"email"`
}
