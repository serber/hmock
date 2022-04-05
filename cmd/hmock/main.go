package main

import (
	"encoding/json"
	"fmt"

	"github.com/serber/hmock/internal/config"
	"github.com/serber/hmock/internal/server"

	"github.com/serber/hmock/internal/server/runner"
)

func main() {
	var portNumber uint16 = 443
	s := server.New(portNumber)

	runner.Run(s)

	config.Print()

	c := config.MockConfiguration{
		Port:     440,
		Protocol: "HTTP",
		Name:     "CardPay OpenAPI mock",
		Mocks: []struct {
			Comment    string              "json:\"_comment\""
			Predicates map[string]struct{} "json:\"predicates\""
			Responses  []struct {
				StatusCode int               "json:\"statusCode\""
				Headers    map[string]string "json:\"headers\""
				Body       struct{}          "json:\"body\""
			} "json:\"responses\""
		}{
			{
				Comment:    "Comment",
				Predicates: make(map[string]struct{}),
			},
		},
	}

	c.Mocks[0].Predicates["A"] = struct{}{}
	c.Mocks[0].Predicates["B"] = struct{}{}

	human_enc, err := json.Marshal(c)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	fmt.Println(string(human_enc))
}
