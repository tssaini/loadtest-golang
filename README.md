# Load Test Golang
Used to load test TCP and UDP applications

## Usage
The load test can be run using the go command. 
```sh
go run . -c config.json
```

Sample config:
```json
{
    "performanceTests": [
        {
            "host": "localhost",
            "sourceType": "tcp",
            "port": 601,
            "rate": 200,
            "activeConnections": 10,
            "interval": 5
        },
        {
            "host": "localhost",
            "sourceType": "udp",
            "port": 514,
            "rate": 20,
            "activeConnections": 10,
            "interval": 30
        }
    ]
    
}
```
Each performance test requires the following parameters:
* <strong>host</strong>. - The address of the host to be tested 
* <strong>sourceType</strong>. - The sourceType can be "udp" or "tcp"
* <strong>port</strong>. - The port number
* <strong>rate</strong>. - The rate of data being sent per connection
* <strong>activeConnections</strong>.- The number of connections to open
* <strong>interval</strong> - The interval of the performance test in seconds
