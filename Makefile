gen-api:
	cd ./cmd/commands/api && ~/go/bin/wire; cd ../..

run-api:
	docker run -d --name poseidon-api poseidon-api

stop-api:
	docker stop poseidon-api

build-api:
	docker build -t poseidon-api -f Dockerfile .