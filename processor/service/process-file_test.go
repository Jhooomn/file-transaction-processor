package service

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"

	"github.com/Jhooomn/file-transaction-processor/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_processorService_process(t *testing.T) {

	data := []map[string]string{
		{"Id": "0", "Date": "7/15", "Transaction": "+60.5"},
		{"Id": "1", "Date": "7/28", "Transaction": "-10.3"},
		{"Id": "2", "Date": "8/2", "Transaction": "-20.46"},
		{"Id": "3", "Date": "8/13", "Transaction": "+10"},
	}

	ctx := context.Background()

	transactionsPerMonth := map[string]uint{
		"July":   2,
		"August": 2,
	}

	totalBalance := 39.74
	avgCredit := 35.25
	avgDebit := -15.38
	creditCount := uint(2)
	debitCount := uint(2)

	tests := []struct {
		name  string
		setUp func(repo *mocks.ProcessorRepository, email *mocks.EmailService)
		args  struct {
			data     []map[string]string
			fileName string
		}
		wantErr bool
	}{
		{
			name: "error email",
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {
				repo.On("Save", ctx, transactionsPerMonth, totalBalance, avgCredit, avgDebit, creditCount, debitCount, "", "").Return(nil)
				email.On("Send", ctx, "", "Stori Total Balance 7/9/2024", mock.Anything).Return(fmt.Errorf("not possible to send"))

			},
			args: struct {
				data     []map[string]string
				fileName string
			}{
				data:     data,
				fileName: "transactions.csv",
			},
			wantErr: true,
		},
		{
			name: "error Save",
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {
				repo.On("Save", ctx, transactionsPerMonth, totalBalance, avgCredit, avgDebit, creditCount, debitCount, "", "").Return(fmt.Errorf("could not save"))
			},
			args: struct {
				data     []map[string]string
				fileName string
			}{
				data:     data,
				fileName: "transactions.csv",
			},
			wantErr: true,
		},
		{
			name: "error empty csv",
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {

			},
			args: struct {
				data     []map[string]string
				fileName string
			}{
				data: []map[string]string{
					{"Id": "", "Date": "", "Transaction": ""},
				},
				fileName: "transactions.csv",
			},
			wantErr: true,
		},
		{
			name: "error parse csv",
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {

			},
			args: struct {
				data     []map[string]string
				fileName string
			}{
				data:     []map[string]string{},
				fileName: "transactions.csv",
			},
			wantErr: true,
		},
		{
			name: "Successful Process",
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {
				repo.On("Save", ctx, transactionsPerMonth, totalBalance, avgCredit, avgDebit, creditCount, debitCount, "", "").Return(nil)
				email.On("Send", ctx, "", "Stori Total Balance 7/9/2024", mock.Anything).Return(nil)

			},
			args: struct {
				data     []map[string]string
				fileName string
			}{
				data:     data,
				fileName: "transactions.csv",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			logger, _ := zap.NewProduction()
			repoMock := mocks.NewProcessorRepository(t)
			emailMock := mocks.NewEmailService(t)
			ps := &processorService{
				logger:              logger,
				processorRepository: repoMock,
				emailService:        emailMock,
			}

			tt.setUp(repoMock, emailMock)

			if err := ps.process(ctx, tt.args.data, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("processorService.process() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
