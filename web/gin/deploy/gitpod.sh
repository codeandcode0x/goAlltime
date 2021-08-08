#!/bin/bash

set -e

sudo apt-get install docker-compose -y

docker-compose -f docker/docker-compose.yml -p ticket  up
