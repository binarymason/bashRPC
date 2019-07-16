#!/usr/bin/env bash

service_manager=$(ps --no-headers -o comm 1)

if [ "$service_manager" = init ]; then
  echo "+ removing init script"
  rm -f /etc/init.d/bashrpc
else
  echo "+ removing systemd service"
  rm -f /etc/systemd/system/bashrpc.service
fi

echo "+ DONE"
