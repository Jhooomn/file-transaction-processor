package service

import (
	"context"

	"go.uber.org/zap"
)

type processorRepository interface {
	Save(ctx context.Context, transactionsPerMonth map[string]uint, totalBalance, avgCredit, avgDebit float64, creditCount, debitCount uint, name, email string) error
}

type ProcessorServiceClient interface {
	Execute()
}

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}
