package main

import (
	"sync"
	"time"

	"github.com/tssaini/syslog-ng-config-testing/connections"
)

func createConns(host string, port string, connType string, activeConnections int) ([]connections.RemoteConn, error) {
	var result []connections.RemoteConn
	if connType == "udp" {
		for i := 0; i < activeConnections; i++ {
			conn, err := connections.NewUDPConn(host, port)
			if err != nil {
				return nil, err
			}
			result = append(result, conn)
		}
	} else if connType == "tcp" {
		for i := 0; i < activeConnections; i++ {
			conn, err := connections.NewTCPConn(host, port)
			if err != nil {
				return nil, err
			}
			result = append(result, conn)
		}
	} else {
		panic("Incorrect connection type")
	}
	return result, nil
}

func generateLogs(remoteConns []connections.RemoteConn, eps int, log string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(remoteConns))
	for _, conn := range remoteConns {
		go sendEPS(log, eps, conn, wg)
	}
	wg.Wait()
}

// sendEPS send logs to destination at specified eps
func sendEPS(log string, eps int, conn connections.RemoteConn, wg *sync.WaitGroup) error {
	defer wg.Done()
	var start time.Time
	var timeElap time.Duration
	var sleepTime time.Duration
	for {
		start = time.Now()
		for i := 0; i < eps; i++ {
			err := conn.Send(log)
			if err != nil {
				return err
			}
		}
		timeElap = time.Now().Sub(start)
		sleepTime = 1*time.Second - timeElap
		//took more than a second to run
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}