package flux_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fluxrepo "github.com/eznd-go/flux/internal/infra/flux"
	"github.com/eznd-go/flux/internal/services/flux"
	"github.com/eznd-go/flux/internal/services/flux/testdata"
)

func TestParsing_OnlineStats(t *testing.T) {
	client := fluxrepo.NewClient(testdata.ServerName, testdata.ServerURL, testdata.User, testdata.Password)
	svc := flux.NewService(client)

	s, err := svc.OnlineStats()
	if err != nil {
		t.Fail()
	}

	require.NoError(t, err)
	assert.Greater(t, s.Current, 0)
	assert.Greater(t, s.Peak, 0)
}
