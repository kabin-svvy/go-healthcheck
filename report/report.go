package report

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request .
type Request struct {
	TotalWebsites int   `json:"total_websites"`
	Success       int   `json:"success"`
	Failure       int   `json:"failure"`
	TotalTime     int64 `json:"total_time"`
}

// Response .
type Response struct {
	Message string `json:"message"`
}

// Create .
func Create(c echo.Context) error {
	req := Request{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, nil)
}

func (r *Request) validate() error {
	if r.TotalWebsites < 0 {
		return errors.New("total websites bad request")
	}
	if r.Success < 0 || r.Success > r.TotalWebsites {
		return errors.New("success bad request")
	}
	if r.Failure < 0 || r.Failure > r.TotalWebsites {
		return errors.New("failure bad request")
	}
	if r.Success+r.Failure > r.TotalWebsites {
		return errors.New("success and failure bad request")
	}
	if r.TotalTime < 0 {
		return errors.New("total time bad request")
	}
	return nil
}
