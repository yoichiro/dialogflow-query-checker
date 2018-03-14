package query

import (
	"github.com/yoichiro/dialogflow-query-checker/config"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io"
	"os"
)

func Execute(test *config.Test, def *config.Definition) (*Response, error) {
	for currentRetryCount := 0;; {
		res, err := doExecute(test, def)
		if err != nil {
			return nil, err
		}
		if res.Status.Code == 200 {
			return res, nil
		}
		currentRetryCount++
		if def.Environment.RetryCount < currentRetryCount {
			return res, nil
		}
		fmt.Print("R")
	}
}

func doExecute(test *config.Test, def *config.Definition) (*Response, error) {
	res, err := send(test, def)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if def.Environment.Debug {
		fmt.Printf("<Response Status> %s\n", res.Status)
	}

	var reader io.Reader = res.Body

	if def.Environment.Debug {
		fmt.Print("<Response body> ")
		reader = io.TeeReader(reader, os.Stdout)
	}

	var response Response
	err = json.NewDecoder(reader).Decode(&response)

	if def.Environment.Debug {
		fmt.Println()
	}

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func send(test *config.Test, def *config.Definition) (*http.Response, error) {
	requestBody := RequestBody{
		Contexts: test.Condition.Contexts,
		Language: test.Condition.Language,
		SessionId: test.Condition.SessionId,
	}
	if test.Condition.Query != "" {
		requestBody.Query = test.Condition.Query
	} else {
		requestBody.Event = Event{
			Name: test.Condition.EventName,
		}
	}
	requestBody.OriginalRequest.Source = "google"
	requestBody.OriginalRequest.Data.User.Locale = test.Condition.Locale
	if test.Condition.ServiceAccessToken != "" {
		requestBody.OriginalRequest.Data.User.AccessToken = test.Condition.ServiceAccessToken
	}

	body, err := json.MarshalIndent(&requestBody, "", "  ")
	if err != nil {
		return nil, err
	}

	if def.Environment.Debug {
		fmt.Printf("<Request body> %s\n", string(body))
	}

	req, err := http.NewRequest("POST", "https://api.dialogflow.com/v1/query?v=20150910", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + def.ClientAccessToken)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}