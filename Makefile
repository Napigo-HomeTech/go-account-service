APP_NAME=account-service
CURDIR=$(PWD)

# for locally build only
local-build: 
	go build -o main

local-run:
	go run ./main.go

# use this to rebuild the service docker container when update code changes	
rebuild:
	echo "Rebuilding container : $(APP_NAME)"
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml kill $(APP_NAME)
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml up -d --build $(APP_NAME)


# Same as "restart" but will rebuild the image before running on docker-container
clean-rebuild:
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml kill $(APP_NAME)
	docker rmi $(APP_NAME) -f
	docker build -t $(APP_NAME) .
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml up -d --build $(APP_NAME)





## Please use below commands wisely as it will clean up most resources such
## as DB caching and redis etc... restarting the environment will be slower

## For restaring the entire cluster right from this repo itself
start-all:
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml up -d
stop-all:
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml down

restart-all:
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml down
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml up -d



clean-all:
	docker-compose -f $(CURDIR)/../npg-dev-tool/docker-compose.yaml down
	docker image prune --all -f
	docker volume prune -f

clean-restart-all:
	make clean-all
	make start-all