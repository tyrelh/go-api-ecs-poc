#! /usr/bin/env bash

url="https://devcatalog.giftbit.com/go/system/version"
delay=1

echo ""; echo "Monitoring ${url} every ${delay} seconds..."; echo ""

while true; do
    echo "GET ${url}"
    curl https://devcatalog.giftbit.com/go/system/version
    echo ""  # Add a newline for better readability
    sleep ${delay}
done