#!/bin/bash

# Test script to simulate port scanning
# This will trigger the detection threshold (default 5 connections)

echo "Starting port scan simulation..."
echo "Connecting to localhost:8080 multiple times to trigger detection..."

# Make 6 connections to trigger the threshold of 5
for i in {1..6}; do
    echo "Connection $i"
    timeout 1 nc -z localhost 8080
    sleep 0.5
done

echo "Port scan simulation completed."
echo "Check the portscammer UI and logs for detection results."
