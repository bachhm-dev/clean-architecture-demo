package open_meteo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/forecast?latitude=10.000000&longitude=20.000000&current_weather=true", r.URL.String())
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	api := New()
	weather, err := api.GetWeather(context.Background(), 10.0, 20.0)

	assert.NoError(t, err)
	assert.NotEmpty(t, weather)
}
