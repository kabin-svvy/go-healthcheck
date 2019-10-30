package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendReportShouldGet200(t *testing.T) {
	t.Run("send report should get 200", func(t *testing.T) {
		expected := 200
		h := Handler{
			URL:         `http://localhost:1323/healthcheck/report`,
			AccessToken: `eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E`,
		}
		req := Request{
			TotalWebsites: 3,
			Success:       2,
			Failure:       1,
			TotalTime:     100000,
		}
		actual, _ := h.SendReport(req)
		assert.Equal(t, expected, actual, "send report should get %v but get %v", expected, actual)
	})
}

func TestSendReportShouldGet404WithFakeURL(t *testing.T) {
	t.Run("send report should get 404 with fake url", func(t *testing.T) {
		expected := 404
		h := Handler{
			URL:         `http://localhost:1323/healthcheck/fakereport`,
			AccessToken: `eyJhbGciOiJIUzI1NiJ9.gYp-4ieS6GdXCU3QOduonlmLAU7E1Tz-DdksGX0a6SinuXLOSqos2Q6zBJiSze7leabHwQQT6DE4qGWtX4dHyzr5CfZZ_HYvU9d7bYRAAIVlCZttEcR4B5q_2SN1KOjD7Ly_I5j6btG0VoXwcNqtZqfxXNUeL9kc4SVVx9g6tFI.CKFYwLEzx63vtmFcWJGIZa1nLllmlpZquIUfhZI7h4E`,
		}
		req := Request{
			TotalWebsites: 3,
			Success:       2,
			Failure:       1,
			TotalTime:     100000,
		}
		actual, _ := h.SendReport(req)
		assert.Equal(t, expected, actual, "send report should get %v but get %v", expected, actual)
	})
}

func TestSendReportShouldGet400WithFakeToken(t *testing.T) {
	t.Run("send report should get 400 with fake token", func(t *testing.T) {
		expected := 400
		h := Handler{
			URL:         `http://localhost:1323/healthcheck/report`,
			AccessToken: `fakeToken`,
		}
		req := Request{
			TotalWebsites: 3,
			Success:       2,
			Failure:       1,
			TotalTime:     100000,
		}
		actual, _ := h.SendReport(req)
		assert.Equal(t, expected, actual, "send report should get %v but get %v", expected, actual)
	})
}

func TestSendReportShouldGet400WithBadRequest(t *testing.T) {
	t.Run("send report should get 400 with bad request", func(t *testing.T) {
		expected := 400
		h := Handler{
			URL:         `http://localhost:1323/healthcheck/report`,
			AccessToken: `fakeToken`,
		}
		req := Request{
			TotalWebsites: 0,
			Success:       2,
			Failure:       1,
			TotalTime:     100000,
		}
		actual, _ := h.SendReport(req)
		assert.Equal(t, expected, actual, "send report should get %v but get %v", expected, actual)
	})
}
