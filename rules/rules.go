package rules

import (
	"encoding/json"
	"github.com/fallenstedt/twitter-stream/httpclient"
)

type (
	//IRules is the interface the rules struct implements.
	IRules interface {
		Create(rules CreateRulesRequest, dryRun bool) (*TwitterRuleResponse, error)
		Delete(req DeleteRulesRequest, dryRun bool) (*TwitterRuleResponse, error)
		GetRules() (*TwitterRuleResponse, error)
	}

	//AddRulesRequest

	//TwitterRuleResponse is what is returned from twitter when adding or deleting a rule.
	TwitterRuleResponse struct {
		Data   []DataRule
		Meta   MetaRule
		Errors []ErrorRule
	}

	//DataRule is what is returned as "Data" when adding or deleting a rule.
	DataRule struct {
		Value string `json:"Value"`
		Tag   string `json:"Tag"`
		Id    string `json:"id"`
	}

	//MetaRule is what is returned as "Meta" when adding or deleting a rule.
	MetaRule struct {
		Sent    string      `json:"sent"`
		Summary MetaSummary `json:"summary"`
	}

	//MetaSummary is what is returned as "Summary" in "Meta" when adding or deleting a rule.
	MetaSummary struct {
		Created    uint `json:"created"`
		NotCreated uint `json:"not_created"`
	}

	//ErrorRule is what is returned as "Errors" when adding or deleting a rule.
	ErrorRule struct {
		Value string `json:"Value"`
		Id    string `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`
	}

	rules struct {
		httpClient httpclient.IHttpClient
	}

)

//NewRules creates a "rules" instance. This is used to create Twitter Filtered Stream rules.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule.
func NewRules(httpClient httpclient.IHttpClient) IRules {
	return &rules{httpClient: httpClient}
}

func (t *rules) Create(rules CreateRulesRequest, dryRun bool) (*TwitterRuleResponse, error) {
	body, err := json.Marshal(rules)
	if err != nil {
		return nil, err
	}

	res, err := t.httpClient.AddRules(func() string {
		if dryRun {
			return "?dry_run=true"
		} else {
			return ""
		}
	}(), string(body))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data := new(TwitterRuleResponse)

	err = json.NewDecoder(res.Body).Decode(data)
	return data, err
}

func (t *rules) Delete(req DeleteRulesRequest, dryRun bool) (*TwitterRuleResponse, error) {

	body, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	res, err := t.httpClient.AddRules(func() string {
		if dryRun {
			return "?dry_run=true"
		} else {
			return ""
		}
	}(), string(body))


	defer res.Body.Close()
	data := new(TwitterRuleResponse)

	err = json.NewDecoder(res.Body).Decode(data)
	return data, err
}


// GetRules gets rules for a stream using twitter's GET GET /2/tweets/search/stream/rules endpoint.
func (t *rules) GetRules() (*TwitterRuleResponse, error) {
	res, err := t.httpClient.GetRules()

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data := new(TwitterRuleResponse)
	json.NewDecoder(res.Body).Decode(data)

	return data, nil
}

