package service

import (
	"context"
	"testing"

	"github.com/Jhooomn/file-transaction-processor/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func Test_processorService_process(t *testing.T) {
	repoMock := mocks.NewProcessorRepository(t)
	emailMock := mocks.NewEmailService(t)

	logger, _ := zap.NewProduction()

	ps := &processorService{
		logger:              logger,
		processorRepository: repoMock,
		emailService:        emailMock,
	}

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
		ps    *processorService
		setUp func(repo *mocks.ProcessorRepository, email *mocks.EmailService)
		args  struct {
			ctx      context.Context
			data     []map[string]string
			fileName string
		}
		wantErr bool
	}{
		{
			name: "Successful Process",
			ps:   ps,
			setUp: func(repo *mocks.ProcessorRepository, email *mocks.EmailService) {
				repoMock.On("Save", ctx, transactionsPerMonth, totalBalance, avgCredit, avgDebit, creditCount, debitCount, "", "").Return(nil)
				emailMock.On("Send", ctx, "", "Stori Total Balance 7/9/2024", mock.Anything).Return(nil)

			},
			args: struct {
				ctx      context.Context
				data     []map[string]string
				fileName string
			}{
				ctx:      context.Background(),
				data:     data,
				fileName: "transactions.csv",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setUp(tt.ps.processorRepository.(*mocks.ProcessorRepository), tt.ps.emailService.(*mocks.EmailService))

			if err := tt.ps.process(tt.args.ctx, tt.args.data, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("processorService.process() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
