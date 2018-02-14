package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
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
	return &def, nil
}
