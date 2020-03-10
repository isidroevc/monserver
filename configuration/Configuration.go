package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	CommunityChain        string `json:"communityChain"`
	MysqlConnectionString string `json:"mysqlConnectionString"`
}

var configuration *Configuration

func readConfiguration() (*Configuration, error) {
	const path = "config.json"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	configuration := new(Configuration)
	parsingError := json.Unmarshal(content, configuration)
	if parsingError != nil {
		return nil, parsingError
	}
	fmt.Println(string(content))
	fmt.Printf("porcentaje: %s\n", configuration.CommunityChain)
	return configuration, nil
}

func GetConfiguration() (*Configuration, error) {
	if configuration == nil {
		config, err := readConfiguration()
		configuration = config
		if err != nil {
			return nil, err
		}
	}
	return configuration, nil
}
