#!/bin/bash
rm -rf /tmp/kobrad-temp

kobrad --devnet --appdir=/tmp/kobrad-temp --profile=6061 &
kobrad_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $kobrad_PID

wait $kobrad_PID
kobrad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kobrad exit code: $kobrad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kobrad_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
