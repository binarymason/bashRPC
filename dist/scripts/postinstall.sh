#!/usr/bin/env bash

service_manager=$(ps --no-headers -o comm 1)

banner() {
  echo "##################################"
}

common_instructions() {
  echo "REQUIRED: add secret to /etc/bashrpc/bashrpc.yml"
}

if [ "$service_manager" = init ]; then
  echo "+ setting up init script"
  mv /etc/bashrpc/init/sysvinit/bashrpc /etc/init.d/bashrpc
  banner
  echo "run the following to get started:"
  common_instructions
  echo "$ chkconfig --add bashrpc"
  echo "$ service bashrpc start"
  banner
else
  echo "+ setting up systemd service"
  mv /etc/bashrpc/init/systemd/bashrpc.service /etc/systemd/system/bashrpc.service
  banner
  echo "run the following to get started:"
  common_instructions
  echo "$ systemctl daemon-reload"
  echo "$ systemctl enable bashrpc"
  echo "$ systemctl start bashrpc"
  banner
fi

echo "+ DONE"
