#!/usr/bin/env bash

set -eu -o pipefail

docker-compose build

#cd terraform-provider-todo
#make testacc TEST=./todo