default: help

help:
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

gen-api: # generate api file
	cd ./cmd/commands/api/providers && ~/go/bin/wire; cd ../..

run-api: # run docker container
	docker run -d --name poseidon-api poseidon-api

stop-api: # stop docker container
	docker stop poseidon-api

build-api: #build api docker image
	docker build -t poseidon-api -f Dockerfile .