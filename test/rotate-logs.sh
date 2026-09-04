 #!/bin/bash

while true; do
    logrotate -s ./logrotate.status ./logrotate.conf
    sleep 5
done
