package main

import (
	"os"

	"github.com/Jhooomn/file-transaction-processor/infrastructure/database"
	"github.com/Jhooomn/file-transaction-processor/infrastructure/email"
	"github.com/Jhooomn/file-transaction-processor/infrastructure/repository"
	"github.com/Jhooomn/file-transaction-processor/processor/service"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if os.Getenv("_LAMBDA_SERVER_PORT") == "" || os.Getenv("AWS_LAMBDA_RUNTIME_API") == "" {
		if err := godotenv.Load(); err != nil {
			logger.Error("Error loading .env file:",
				zap.String("error", err.Error()),
			)
			return
		}
	}

	dbName := os.Getenv("DB_NAME")

	dbClient, close := database.NewConnection()
	defer close()

	processorRepository := repository.NewProcessorService(dbClient, dbName)
	emailService := email.NewEmailService(os.Getenv("SMTP_FROM"), os.Getenv("SMTP_PSW"), os.Getenv("SMTP_SERVER"))

	service := service.NewProcessorService(os.Getenv("DATA_PATH"), 4, logger, processorRepository, emailService) // TODO: configure go-r-pool
	service.Execute()
	logger.Info("Running the lambda!")
	lambda.Start(service.Execute)
}
