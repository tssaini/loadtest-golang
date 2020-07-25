package main

import (
	"fmt"
	"time"

	"github.com/tssaini/loadtest-golang/connections"
)

// CreateConns creates remote connections
func CreateConns(host string, port string, connType string, activeConnections int) ([]connections.RemoteConn, error) {
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
		return nil, fmt.Errorf("Incorrect connection type: %v", connType)
	}
	return result, nil
}

// GenerateRate sends logs at rate to each connection
func GenerateRate(log string, rate int, interval int32, remoteConns []connections.RemoteConn) error {
	var doneChans []chan struct{}
	errChan := make(chan error)
	for _, conn := range remoteConns {
		doneChan := make(chan struct{})
		doneChans = append(doneChans, doneChan)
		go func(conn connections.RemoteConn, doneChan <-chan struct{}) {
			err := sendEPS(log, rate, conn, doneChan)
			if err != nil {
				errChan <- err
			}
		}(conn, doneChan)
	}
	select {
	case err := <-errChan:
		// close all the go routines launched above
		for _, doneChan := range doneChans {
			close(doneChan)
		}
		return err
	case <-time.After(time.Duration(interval) * time.Second):
		for _, doneChan := range doneChans {
			close(doneChan)
		}
		return nil
	}
}

// GenerateN sends n number of logs to each remoteConns
func GenerateN(log string, n int, remoteConns []connections.RemoteConn) error {
	for _, conn := range remoteConns {
		for i := 0; i < n; i++ {
			err := conn.Send(log)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// sendEPS send logs to destination at specified eps
func sendEPS(log string, rate int, conn connections.RemoteConn, done <-chan struct{}) error {
	var start time.Time
	var timeElap time.Duration
	var sleepTime time.Duration
	for {
		select {
		case <-done:
			return nil
		default:
		}
		start = time.Now()
		for i := 0; i < rate; i++ {
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
