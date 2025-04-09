package flux

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/antchfx/htmlquery"

	"github.com/eznd-go/flux/internal/domain"
)

const (
	vendingBasePath = "?module=vending&action=items&p="
)

var (
	statsRE = regexp.MustCompile(
		`Found a total of (\d+) record\(s\) across (\d+) page\(s\)\.  Displaying result\(s\) 1-(\d+)\.`)
)

func (s *service) Market() ([]domain.Shop, error) {
	return nil, nil
}

func (s *service) MarketStats() (domain.MarketStats, error) {
	data, err := s.client.Get(vendingBasePath)
	if err != nil {
		return domain.MarketStats{}, err
	}

	doc, err := htmlquery.Parse(bytes.NewReader(data))
	if err != nil {
		return domain.MarketStats{}, err
	}

	list := htmlquery.Find(doc, `//p[contains(text(), "Found a total of ")]`)

	if len(list) != 1 {
		return domain.MarketStats{}, domain.ErrParsingFailed
	}

	msg := htmlquery.InnerText(list[0])

	nns := statsRE.FindStringSubmatch(msg)
	if len(nns) != 4 {
		return domain.MarketStats{}, domain.ErrParsingFailed
	}

	var res domain.MarketStats

	res.TotalOrders, err = strconv.Atoi(nns[1])
	if err != nil {
		return domain.MarketStats{}, err
	}

	res.TotalPages, err = strconv.Atoi(nns[2])
	if err != nil {
		return domain.MarketStats{}, err
	}

	res.OrdersPerPage, err = strconv.Atoi(nns[3])
	if err != nil {
		return domain.MarketStats{}, err
	}

	// чиним потенциально кривую калькуляцию при единственной неполной странице
	if res.TotalPages == 1 && res.OrdersPerPage < 20 {
		res.OrdersPerPage = 20
	}

	return res, err
}
