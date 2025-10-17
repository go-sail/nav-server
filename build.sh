#!/bin/bash

set -e

dt=`date +%Y%m%d`
tag="go-sail/nav-server:${dt}"
archive="nav-server-${dt}.zip"

echo "build image tag is: ${tag}"

docker build --no-cache -t ${tag} .

docker save ${tag} -o ~/Downloads/${archive}