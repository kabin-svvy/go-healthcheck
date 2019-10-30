package report

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler .
type Handler struct {
	URL         string
	AccessToken string
}

// SendReport .
func (h Handler) SendReport(req Request) (Response, int, error) {
	res := Response{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	newResp, err := http.NewRequest(http.MethodPost, h.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return res, http.StatusInternalServerError, err
	}
	newResp.Header.Add("Content-Type", "application/json")
	newResp.Header.Add("Authorization", fmt.Sprintf("Bearer %v", h.AccessToken))

	resp, err := http.DefaultClient.Do(newResp)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	// log.Printf("%v", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	if resp.StatusCode != http.StatusOK {
		return res, resp.StatusCode, errors.New(resp.Status)
	}
	return res, resp.StatusCode, nil
}
