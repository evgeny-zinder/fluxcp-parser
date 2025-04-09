package services

import (
	"github.com/eznd-go/flux/internal/domain"
)

type FluxCP interface {
	MarketStats() (domain.MarketStats, error)
	Market() ([]domain.Shop, error)
	OnlineStats() (domain.OnlineStats, error)
}
