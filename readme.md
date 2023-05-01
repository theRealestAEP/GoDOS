> go build godos.go

For an HTTP stress test:
./godos -mode=http -target="http://localhost:8080"

For a TCP stress test:
./godos -mode=tcp -target="localhost:12345"

For a UDP stress test:
./godos -mode=udp -target="localhost:12345"

Repalce localhost for the target url and the port accordingly