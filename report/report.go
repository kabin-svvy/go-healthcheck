package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Handler .
type Handler struct {
	URL         string
	AccessToken string
}

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

// SendReport .
func (h Handler) SendReport(req Request) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	newResp, err := http.NewRequest(http.MethodPost, h.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	newResp.Header.Add("Content-Type", "application/json")
	newResp.Header.Add("Authorization", fmt.Sprintf("Bearer %v", h.AccessToken))

	resp, err := http.DefaultClient.Do(newResp)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}
