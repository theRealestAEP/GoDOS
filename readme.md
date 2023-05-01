Default settings (10 workers, 100 requests per worker, http mode, and target localhost:8080):
go run godos.go

Customizing the number of workers and requests per worker:

go run godos.go -workers 20 -requests 50

Running the test in TCP mode with a custom target IP and port:

go run godos.go -mode tcp -target 192.168.1.100:9090 -workers 15 -requests 200

Running the test in UDP mode with a custom target IP and port:

go run godos.go -mode udp -target 192.168.1.100:9090 -workers 15 -requests 200


These commands demonstrate various combinations of flags that can be used to customize the load test. Make sure to replace the target IP and port with the actual values you want to test.