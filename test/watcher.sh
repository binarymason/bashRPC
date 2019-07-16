#!/usr/bin/env bash
#
# runs `go test` on file changes

while true
do
  inotifywait -qq -r -e create,close_write,modify,move,delete ./ && \
    sleep 1 && \
    clear && \
    go test ./... "$@" && ./test/integration_test.sh
done
