package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/tssaini/syslog-ng-config-testing/util"
	"github.com/urfave/cli/v2"
)

func main() {
	var filePath string

	app := &cli.App{
		Name:  "LoadTest",
		Usage: "Used to to test log forwarders",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "config.json",
				Destination: &filePath,
				Usage:       "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			executeTests(filePath)
			return nil
		},
	}

	app.Run(os.Args)
}

func executeTests(filename string) {
	config, err := ParseConfig(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(config.PerformanceTests))
	for _, perfTest := range config.PerformanceTests {

		go func(test PerformanceTest) {
			util.InfoLogger.Printf("Starting test: %v\n", test)
			defer wg.Done()
			err := executePerformanceTest(test)
			if err != nil {
				log.Fatalf("%v", err)
			}
			util.InfoLogger.Printf("Completed test: %v\n", test)
		}(perfTest)
	}
	wg.Wait()
}

func executeIntegrationTest(config IntegrationTest, wg *sync.WaitGroup) (bool, error) {
	defer wg.Done()
	conns, err := CreateConns(config.Host, fmt.Sprintf("%v", config.Port), config.SourceType, 1)
	if err != nil {
		return false, err
	}
	err = GenerateN(config.InputMessage, 1, conns)
	if err != nil {
		return false, err
	}
	return true, nil
}

func executePerformanceTest(config PerformanceTest) error {
	conns, err := CreateConns(config.Host, fmt.Sprintf("%v", config.Port), config.SourceType, config.ActiveConnections)
	if err != nil {
		return err
	}

	err = GenerateRate("Hello world chrome!", config.Rate, config.Interval, conns)
	if err != nil {
		return err
	}
	return nil
}
