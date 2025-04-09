package flux

import (
	"time"

	"github.com/eznd-go/flux/internal/infra"
)

type service struct {
	client   infra.FluxCP
	cookie   string
	loggedAt time.Time
}

func NewService(client infra.FluxCP) *service {
	return &service{
		client: client,
	}
}
