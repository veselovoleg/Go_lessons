package github

import (
"net/http"
"net/url"
"strings"
"fmt"
"encoding/json"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" +q)
		if err != nil {
			return nil, err
		}
	//Close resp.Body
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("Query error: %s", resp.Status)
		}
	var result IssuesSearchResult

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // Streaming Decoder. Fill variable "result"
			resp.Body.Close()
			return nil, err
		}
	resp.Body.Close()
	return &result, nil
}