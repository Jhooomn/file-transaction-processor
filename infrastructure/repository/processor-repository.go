package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProcessorRepository interface {
	Save(ctx context.Context, transactionsPerMonth map[string]uint, totalBalance, avgCredit, avgDebit float64, creditCount, debitCount uint, name, email string) error
	findByNameAndEmail(ctx context.Context, name, email string) (*TransactionSummary, error)
}

type processorRepository struct {
	collection *mongo.Collection
}

func NewProcessorService(dbClient *mongo.Client, dbName string) ProcessorRepository {
	collection := dbClient.Database(dbName).Collection("user-summary")

	return &processorRepository{
		collection: collection,
	}
}

func (pr *processorRepository) Save(ctx context.Context,
	transactionsPerMonth map[string]uint,
	totalBalance,
	avgCredit,
	avgDebit float64,
	creditCount,
	debitCount uint,
	name,
	email string) error {

	document := bson.D{
		{Key: "transactionsPerMonth", Value: transactionsPerMonth},
		{Key: "totalBalance", Value: totalBalance},
		{Key: "avgCredit", Value: avgCredit},
		{Key: "avgDebit", Value: avgDebit},
		{Key: "creditCount", Value: creditCount},
		{Key: "debitCount", Value: debitCount},
		{Key: "name", Value: name},
		{Key: "email", Value: email},
	}

	existingDocument, err := pr.findByNameAndEmail(ctx, name, email)
	if err != nil {
		return err
	}

	if existingDocument != nil && !existingDocument.ID.IsZero() {
		// Update existing document
		filter := bson.M{"_id": existingDocument.ID}
		update := bson.M{
			"$set": document,
		}
		_, err = pr.collection.UpdateOne(ctx, filter, update)
		if err != nil {

			return err
		}
	} else {
		// Insert new document
		_, err := pr.collection.InsertOne(ctx, document)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pr *processorRepository) findByNameAndEmail(ctx context.Context, name, email string) (*TransactionSummary, error) {

	filter := bson.D{
		{Key: "name", Value: name},
		{Key: "email", Value: email},
	}

	var result bson.M
	err := pr.collection.FindOne(ctx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// Get not data structure fields
	summary := &TransactionSummary{
		ID:           result["_id"].(primitive.ObjectID),
		Name:         result["name"].(string),
		Email:        result["email"].(string),
		TotalBalance: result["totalBalance"].(float64),
	}

	// safe way to handle map
	if tpm, ok := result["transactionsPerMonth"].(map[string]interface{}); ok {
		summary.TransactionsPerMonth = make(map[string]uint)
		for k, v := range tpm {
			if val, ok := v.(int32); ok {
				summary.TransactionsPerMonth[k] = uint(val)
			} else {
				return nil, fmt.Errorf("type assertion failed for transactionsPerMonth")
			}
		}
	}

	// safe way to handle map
	if avgCredit, ok := result["avgCredit"].(map[string]interface{}); ok {
		summary.AvgCredit = make(map[string]float64)
		for k, v := range avgCredit {
			if val, ok := v.(float64); ok {
				summary.AvgCredit[k] = val
			} else {
				return nil, fmt.Errorf("type assertion failed for avgCredit")
			}
		}
	}

	// safe way to handle map
	if avgDebit, ok := result["avgDebit"].(map[string]interface{}); ok {
		summary.AvgDebit = make(map[string]float64)
		for k, v := range avgDebit {
			if val, ok := v.(float64); ok {
				summary.AvgDebit[k] = val
			} else {
				return nil, fmt.Errorf("type assertion failed for avgDebit")
			}
		}
	}

	// safe way to handle map
	if creditCount, ok := result["creditCount"].(map[string]interface{}); ok {
		summary.CreditCount = make(map[string]uint)
		for k, v := range creditCount {
			if val, ok := v.(int32); ok {
				summary.CreditCount[k] = uint(val)
			} else {
				return nil, fmt.Errorf("type assertion failed for creditCount")
			}
		}
	}

	// safe way to handle map
	if debitCount, ok := result["debitCount"].(map[string]interface{}); ok {
		summary.DebitCount = make(map[string]uint)
		for k, v := range debitCount {
			if val, ok := v.(int32); ok {
				summary.DebitCount[k] = uint(val)
			} else {
				return nil, fmt.Errorf("type assertion failed for debitCount")
			}
		}
	}

	return summary, nil
}
