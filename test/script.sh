 #!/bin/bash

./rotate-logs.sh &
rotate_pid=$!

cleanup() {
    kill "$rotate_pid" 2>/dev/null
}

trap cleanup EXIT

apache_url="http://localhost:8080"

while true; do
    curl -s -o /dev/null "$apache_url"
done
