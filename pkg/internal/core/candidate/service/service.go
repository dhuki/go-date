package service

import (
	"context"

	"github.com/dhuki/go-date/pkg/internal/core/candidate/port"
)

type candidateServiceImpl struct {
	repository port.CandidateRepository
}

func NewCandidateService(candidateRepository port.CandidateRepository) port.CandidateService {
	return candidateServiceImpl{
		repository: candidateRepository,
	}
}

func (u candidateServiceImpl) SwipeAction(ctx context.Context) {

}
