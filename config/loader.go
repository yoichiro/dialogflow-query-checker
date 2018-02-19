package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"github.com/google/uuid"
)

func LoadConfigurationFile(path string) (*Definition, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var def Definition
	err = yaml.Unmarshal(buf, &def)
	if err != nil {
		return nil, err
	}

	determineClientAccessToken(&def)
	determineSessionId(&def)

	return &def, nil
}
func determineSessionId(def *Definition) {
	sessionId := issueSessionId()
	for i := range def.Tests {
		condition := &def.Tests[i].Condition
		if condition.SessionId == "" || condition.SessionId == "inherit" {
			condition.SessionId = sessionId
		} else if condition.SessionId == "new" {
			sessionId = issueSessionId()
			condition.SessionId = sessionId
		} else {
			// Use the specified value as a Session ID.
		}
	}
}

func issueSessionId() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func determineClientAccessToken(def *Definition) {
	if def.ClientAccessToken == "" {
		def.ClientAccessToken = os.Getenv("DIALOGFLOW_CLIENT_ACCESS_TOKEN")
	}
}