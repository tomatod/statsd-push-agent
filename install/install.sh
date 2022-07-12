#!/bin/bash
set -eu

base=/opt/statsd-push-agent/
mkdir -p $base
cp -f ./statsd-push-agent $base
cp -n ./config.yaml $base

cp -f ./statsd-push-agent.service /lib/systemd/system/
systemctl daemon-reload
systemctl enable statsd-push-agent.service

echo "
Finish installation!

Please rewirte /opt/statsd-push-agent/config.yaml according to your requirements.
Please run 'systemctl start statsd-push-agent' for starting.

For more information, please refer https://github.com/tomatod/statsd-push-agent
"