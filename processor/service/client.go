package service

type processorService struct {
	opts                opts
	logger              Logger
	processorRepository processorRepository
	emailService        emailService
}

type opts struct {
	dataPath   string
	workerPool int
}

func NewProcessorService(
	dataPath string,
	workerPool int,
	logger Logger,
	processorRepository processorRepository,
	emailService emailService,
) ProcessorServiceClient {
	return &processorService{
		opts: opts{
			dataPath:   dataPath,
			workerPool: workerPool,
		},
		logger:              logger,
		processorRepository: processorRepository,
		emailService:        emailService,
	}
}
