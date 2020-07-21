package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	// host := "127.0.0.1"
	// port := "601"
	// conns, err := createConns(host, port, "tcp", 10)
	// if err != nil {
	// 	log.Fatalf("Unable to create connections")
	// }
	// generateLogs(conns, 20, "Hello google world!")
	config, err := ParseConfig("examples/test-config.json")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	wg := sync.WaitGroup{}
	// wg.Add(len(config.IntegrationTests))
	// for _, integrationTest := range config.IntegrationTests {
	// 	_, err := executeIntegrationTest(integrationTest, &wg)
	// 	if err != nil {
	// 		log.Fatalf("unable to execute testcase %v", err)
	// 	}
	// }
	// wg.Wait()
	wg.Add(1)
	for _, perfTest := range config.PerformanceTests {

		go func(test PerformanceTest) {
			defer wg.Done()
			err := executePerformanceTest(test)
			if err != nil {
				log.Fatalf("%v", err)
			}
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
