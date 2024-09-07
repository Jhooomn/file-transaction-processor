package main

import (
	"os"

	"github.com/Jhooomn/file-transaction-processor/processor/service"
	"github.com/Jhooomn/file-transaction-processor/utils"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {

	logger := utils.NewLogger()

	if os.Getenv("_LAMBDA_SERVER_PORT") == "" || os.Getenv("AWS_LAMBDA_RUNTIME_API") == "" {
		if err := godotenv.Load(); err != nil {
			logger.Error("Error loading .env file:", err)
			return
		}
	}

	service := service.NewProcessorService(os.Getenv("DATA_PATH"), 1, logger) // TODO: configure go-r-pool
	service.Execute()
	logger.Info("Running the lambda!")
	lambda.Start(service.Execute)
}
