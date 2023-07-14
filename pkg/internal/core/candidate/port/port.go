package port

import (
	"context"
)

type CandidateService interface {
	SwipeAction(ctx context.Context)
}

type CandidateRepository interface {
}
