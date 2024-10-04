#!/bin/bash
rm -rf /tmp/kobrad-temp

NUM_CLIENTS=128
kobrad --devnet --appdir=/tmp/kobrad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
kobrad_PID=$!
kobrad_KILLED=0
function killkobradIfNotKilled() {
  if [ $kobrad_KILLED -eq 0 ]; then
    kill $kobrad_PID
  fi
}
trap "killkobradIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $kobrad_PID

wait $kobrad_PID
kobrad_EXIT_CODE=$?
kobrad_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "kobrad exit code: $kobrad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kobrad_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
