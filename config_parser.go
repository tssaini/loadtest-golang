package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// TestConfig represents the json config
type TestConfig struct {
	IntegrationTests []IntegrationTest
	PerformanceTests []PerformanceTest
}

// IntegrationTest testcase info
type IntegrationTest struct {
	Host            string
	SourceType      string
	Port            int
	InputMessage    string
	ExpectedMessage string
	DestinationName string
}

// PerformanceTest testcase info
type PerformanceTest struct {
	Host              string
	SourceType        string
	Port              int
	Rate              int
	ActiveConnections int
	Interval          time.Duration
}

// ParseConfig parses the given json file path into the Testconfig
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
