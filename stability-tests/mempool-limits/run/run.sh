#!/bin/bash

APPDIR=/tmp/kobrad-temp
kobrad_RPC_PORT=29587

rm -rf "${APPDIR}"

kobrad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${kobrad_RPC_PORT}" --profile=6061 &
kobrad_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${kobrad_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $kobrad_PID

wait $kobrad_PID
kobrad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kobrad exit code: $kobrad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kobrad_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
