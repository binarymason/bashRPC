#!/usr/bin/env bash

service_manager=$(ps --no-headers -o comm 1)

# stop the bashrpc service if it's running
service bashrpc stop &>/dev/null

if [ "$service_manager" = init ]; then
  echo "+ removing init script"
  rm -f /etc/init.d/bashrpc
else
  echo "+ removing systemd service"
  rm -f /etc/systemd/system/bashrpc.service
fi

echo "+ DONE"
