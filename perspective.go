package goperspective

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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
func (c *Client) AnalyzeComment(r map[string]interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+":analyze", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	addQuery(req, "key", c.key)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return respToMap(res)
}

//SuggestCommentScore takes map representation of JSON-object and
//lets you give the API feedback by allowing you to suggest a score that you think the API should have returned.
//All reqired and optional parameters can be found here:
//https://developers.perspectiveapi.com/s/about-the-api-methods.
func (c *Client) SuggestCommentScore(r map[string]interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+":suggestscore", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	addQuery(req, "key", c.key)

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return respToMap(res)
}

//addQuery adds query parameter to the request.
func addQuery(req *http.Request, key, value string) {
	q := req.URL.Query()
	q.Add(key, value)
	req.URL.RawQuery = q.Encode()
}

//respToMap converts json body to map.
func respToMap(res *http.Response) (map[string]interface{}, error) {
	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var s map[string]interface{}
	err = json.Unmarshal(rb, &s)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		errj := s["error"].(map[string]interface{})
		return s, fmt.Errorf("code: %d. %s", res.StatusCode, errj["message"])
	}
	return s, nil
}
