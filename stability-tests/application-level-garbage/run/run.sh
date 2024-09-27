#!/bin/bash
rm -rf /tmp/kobrad-temp

kobrad --devnet --appdir=/tmp/kobrad-temp --profile=6061 --loglevel=debug &
kobrad_PID=$!
kobrad_KILLED=0
function killkobradIfNotKilled() {
    if [ $kobrad_KILLED -eq 0 ]; then
      kill $kobrad_PID
    fi
}
trap "killkobradIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $kobrad_PID

wait $kobrad_PID
kobrad_KILLED=1
kobrad_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "kobrad exit code: $kobrad_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $kobrad_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
