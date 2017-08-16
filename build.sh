#! /usr/bin/env bash

rm -rf ./vendor
rm glide.lock
rm ./kubernetes-scaleio-prom
glide up

grep -R --exclude-dir vendor --exclude-dir .git --exclude-dir mesos --exclude build.sh TODO ./

GOOS=linux GOARCH=amd64 go build .
