package goperspective

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//Client is the Perspecive API HTTP client.
type Client struct {
	baseURL string
	key     string
	http    *http.Client
}

//NewClient creates Perspective API client.
func NewClient(apiKey string) *Client {

	c := &Client{
		baseURL: "https://commentanalyzer.googleapis.com/v1alpha1/comments",
		key:     apiKey,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return c
}

//AnalyzeComment takes map representation of JSON-object and returns scores map for the comment.
//All reqired and optional parameters can be found here:
//https://developers.perspectiveapi.com/s/about-the-api-methods.
func (c *Client) AnalyzeComment(data AnalyzeRequest) (AnalyzeResponse, error) {
	var ar AnalyzeResponse
	body, err := json.Marshal(data)
	if err != nil {
		return ar, err
	}

	req, err := http.NewRequest("POST", c.baseURL+":analyze", bytes.NewBuffer(body))
	if err != nil {
		return ar, err
	}
	addQuery(req, "key", c.key)

	res, err := c.http.Do(req)
	if err != nil {
		return ar, err
	}
	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ar, err
	}

	err = json.Unmarshal(rb, &ar)
	if err != nil {
		return ar, err
	}

	if res.StatusCode != http.StatusOK {
		return ar, fmt.Errorf("code: %d. %s", res.StatusCode, ar.Error.Message)
	}
	return ar, nil
}

//SuggestCommentScore takes map representation of JSON-object and
//lets you give the API feedback by allowing you to suggest a score that you think the API should have returned.
//All reqired and optional parameters can be found here:
//https://developers.perspectiveapi.com/s/about-the-api-methods.
func (c *Client) SuggestCommentScore(data SuggestRequest) (SuggestResposne, error) {
	var sr SuggestResposne
	body, err := json.Marshal(data)
	if err != nil {
		return sr, err
	}

	req, err := http.NewRequest("POST", c.baseURL+":suggestscore", bytes.NewBuffer(body))
	if err != nil {
		return sr, err
	}
	addQuery(req, "key", c.key)

	res, err := c.http.Do(req)
	if err != nil {
		return sr, err
	}
	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return sr, err
	}

	err = json.Unmarshal(rb, &sr)
	if err != nil {
		return sr, err
	}

	if res.StatusCode != http.StatusOK {
		return sr, fmt.Errorf("code: %d. %s", res.StatusCode, sr.Error.Message)
	}
	return sr, nil
}

//addQuery adds query parameter to the request.
func addQuery(req *http.Request, key, value string) {
	q := req.URL.Query()
	q.Add(key, value)
	req.URL.RawQuery = q.Encode()
}
