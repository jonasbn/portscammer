#!/bin/bash

# Test the port scanner with debug logging and lower threshold

echo "Starting portscammer with debug logging and threshold=2..."
echo "This will make it easier to trigger scan detection."
echo ""
echo "In another terminal, run: ./test_scan.sh"
echo "Or manually: nc -z localhost 8080 (run this 3+ times quickly)"
echo ""

./portscammer -d -t 2 -l debug.log
