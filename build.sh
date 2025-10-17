#!/bin/bash

set -e

dt=`date +%Y%m%d`
tag="heybox/servers:${dt}"
archive="heybox-servers-${dt}.zip"

echo "build image tag is: ${tag}"

docker build --no-cache -t ${tag} .

docker save ${tag} -o ~/Downloads/${archive}