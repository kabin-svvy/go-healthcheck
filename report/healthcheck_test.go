package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Run("health check should be success", func(t *testing.T) {
		expected := 1
		url := "https://www.google.co.th"
		req := Request{}
		_ = req.HealthCheck(url)
		actual := req.Success
		assert.Equal(t, expected, actual, "health check success should be %v but get %v", expected, actual)
	})

	t.Run("health check should be failure", func(t *testing.T) {
		expected := 1
		url := "https://www.google.co.th/healthcheck"
		req := Request{}
		_ = req.HealthCheck(url)
		actual := req.Failure
		assert.Equal(t, expected, actual, "health check success should be %v but get %v", expected, actual)
	})

	t.Run("health check time should be great than 0", func(t *testing.T) {
		var expected int64
		url := "https://www.google.co.th"
		req := Request{}
		_ = req.HealthCheck(url)
		actual := req.TotalTime
		if actual <= expected {
			t.Errorf("health check time should be great than %v but get %v", expected, actual)
		}
	})
}
