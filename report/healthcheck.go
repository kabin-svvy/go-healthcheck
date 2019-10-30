package report

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
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

// HealthCheck .
func (r *Request) HealthCheck(uri string) error {
	now := time.Now()
	nano := now.UnixNano()
	defer r.AddTotalTime(nano)
	r.TotalWebsites++
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		r.Failure++
		return errors.New(res.Status)
	}
	r.Success++
	return nil
}

// AddTotalTime .
func (r *Request) AddTotalTime(start int64) {
	now := time.Now()
	nano := now.UnixNano()
	r.TotalTime += (nano - start)
}
