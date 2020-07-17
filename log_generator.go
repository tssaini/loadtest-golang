package main

import (
	"time"

	"github.com/tssaini/syslog-ng-config-testing/destinations"
)

// generate send logs to destination at specified eps
func generate(log string, eps int, conn destinations.RemoteConn) error {
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
