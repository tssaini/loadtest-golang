package main

import (
	"encoding/json"
	"io/ioutil"
)

type TestConfig struct {
	IntegrationTests []IntegrationTest
	PerformanceTests []PerformanceTest
}

type IntegrationTest struct {
	Host            string
	SourceType      string
	Port            int
	InputMessage    string
	ExpectedMessage string
	DestinationName string
}

type PerformanceTest struct {
	Host              string
	SourceType        string
	Port              int
	Rate              int
	ActiveConnections int
}

func ParseConfig(fileName string) (*TestConfig, error) {
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var config TestConfig
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
