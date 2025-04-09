package flux

import (
	"fmt"

	"github.com/eznd-go/flux/internal/domain"
	"github.com/eznd-go/flux/pkg/htmlqx"
)

func (s *service) OnlineStats() (domain.OnlineStats, error) {
	data, err := s.client.Get(vendingBasePath)
	if err != nil {
		return domain.OnlineStats{}, err
	}

	online, err := htmlqx.
		Parse(data).
		Int(`//section[@class="status"]//span[text()="ONLINE"]/following-sibling::span`)
	if err != nil {
		return domain.OnlineStats{}, fmt.Errorf("parse online stats: %w", err)
	}

	peak, err := htmlqx.
		Parse(data).
		Int(`//section[@class="status"]//span[text()="PEAK"]/following-sibling::span`)
	if err != nil {
		return domain.OnlineStats{}, fmt.Errorf("parse peak stats: %w", err)
	}

	return domain.OnlineStats{
		Current: online,
		Peak:    peak,
	}, err
}
