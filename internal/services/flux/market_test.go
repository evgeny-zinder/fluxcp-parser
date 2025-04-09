package flux_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fluxrepo "github.com/eznd-go/flux/internal/infra/flux"
	"github.com/eznd-go/flux/internal/services/flux"
	"github.com/eznd-go/flux/internal/services/flux/testdata"
)

func TestParsing_MarketStat(t *testing.T) {
	client := fluxrepo.NewClient(testdata.ServerName, testdata.ServerURL, testdata.User, testdata.Password)
	svc := flux.NewService(client)

	n, err := svc.MarketStats()
	if err != nil {
		t.Fail()
	}

	require.NoError(t, err)
	assert.Greater(t, n.TotalOrders, 0)
	assert.Greater(t, n.TotalPages, 0)
	assert.Equal(t, n.OrdersPerPage, 20)
}

//func TestParsing_MarketPage(t *testing.T) {
//	client := fluxrepo.NewClient("https://wayragnarok.com", user, password)
//	svc := flux.NewService(client)
//
//	n, err := svc.MarketStats()
//	if err != nil {
//		t.Fail()
//	}
//
//	require.NoError(t, err)
//	assert.Greater(t, n.TotalOrders, 0)
//	assert.Greater(t, n.TotalPages, 0)
//	assert.Equal(t, n.OrdersPerPage, 20)
//}
