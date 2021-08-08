#!/bin/bash

set -eo pipefail

appVersion=$1
repoUser=$2
repoHost=$3

if [ ! $appVersion ]; then
  appVersion="latest" 
fi

if [ ! $repoUser ]; then
  repoUser="roandocker" 
fi

if [ ! $repoHost ]; then
  repoHost="docker.io" 
fi

# build image
function build() {
	version=$1
	repoUser=$2
	repoHost=$3
	docker build -t $repoHost/$repoUser/gin-scaffold:$version -f docker/projects/golang/gin-scaffold/Dockerfile  --no-cache ./
}

# prune image
function prune() {
	sleep 10 && docker system prune -f
}

# push image
function push() {
	version=$1
	repoUser=$2
	repoHost=$3
	docker push $repoHost/$repoUser/gin-scaffold:$version
}



#main 
function main() {
	build $1 $2 $3
	prune
	push $1 $2 $3
}

# run main
main $appVersion $repoUser $repoHost

export TICKET_VERSION=$appVersion