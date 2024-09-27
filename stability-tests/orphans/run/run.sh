#!/bin/bash
rm -rf /tmp/kobrad-temp

kobrad --simnet --appdir=/tmp/kobrad-temp --profile=6061 &
kobrad_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $kobrad_PID

wait $kobrad_PID
kobrad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kobrad exit code: $kobrad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kobrad_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
