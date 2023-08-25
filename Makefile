# start the environment of douyin
.PHONY: start
start:
	docker-compose up -d

# stop the environment of douyin
.PHONY: stop
stop:
	docker-compose down


# run the api
.PHONY: api
api:
	go build -o ./api_gateway/api ./api_gateway
	go mod tidy
	./api_gateway/api

# run the user
.PHONY: user
user:
	go build -o ./rpc/user/userservice ./rpc/user
	go mod tidy
	./rpc/user/userservice