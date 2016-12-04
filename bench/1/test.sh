#!/usr/bin/env bash

set -x

# Should be set to project root dir
PROJECT_NAME='github.com/dlsniper/u/bench/1'

# Use the latest Go version
GO_VERSION='1.6.2-alpine'

# Do not edit below unless you know what you are doing

# This allows for execution from the current directory
# where the invocation happens. Should be running from
# the top directory
PROJECT_DIR="${GOPATH}/${PROJECT_NAME}"

CONTAINER_GOPATH='/go'
CONTAINER_PROJECT_DIR="${CONTAINER_GOPATH}/src/${PROJECT_NAME}"

echo "" > output.log
tail -f output.log &
TAIL_PID=$!

docker run --rm \
    --net="host" \
    -v ${PROJECT_DIR}:${CONTAINER_PROJECT_DIR} \
    -e CI=true \
    -e GODEBUG=netdns=go \
    -e GOPATH=${CONTAINER_GOPATH} \
    -w "${CONTAINER_PROJECT_DIR}" \
    golang:${GO_VERSION} \
    go test -v -race ./... 2> output.log

EXIT_CODE=$?
kill ${TAIL_PID}

if [ ${EXIT_CODE} != 0 ]; then
    exit ${EXIT_CODE}
fi

# Check for race conditions as we don't have a proper exit code for them from the tool
cat output.log | grep -v 'WARNING: DATA RACE'

EXIT_CODE=$?

# If we don't find a match then we don't have a race condition
if [ ${EXIT_CODE} == 1 ]; then
    # Remove the output file on success only
    rm -f output.log
    exit 0
fi

exit ${EXIT_CODE}
