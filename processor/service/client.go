package service

import "github.com/sirupsen/logrus"

type processorService struct {
	opts   opts
	logger *logrus.Logger
	// email service client
	// file  service client
	// user  repo    client
}

type opts struct {
	dataPath   string
	workerPool int
}

func NewProcessorService(
	dataPath string,
	workerPool int,
	logger *logrus.Logger,
) ProcessorService {
	return &processorService{
		opts: opts{
			dataPath:   dataPath,
			workerPool: workerPool,
		},
		logger: logger,
	}
}
