#!/bin/bash

set -e

docker-compose --env-file deploy/config/.env -f docker/docker-compose.yml -p gin-scaffold  up