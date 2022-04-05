package config

type MockConfiguration struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Name     string `json:"name"`
	Mocks    []struct {
		Comment    string              `json:"_comment"`
		Predicates map[string]struct{} `json:"predicates"`
		Responses  []struct {
			StatusCode int               `json:"statusCode"`
			Headers    map[string]string `json:"headers"`
			Body       struct {
			} `json:"body"`
		} `json:"responses"`
	} `json:"mocks"`
}
