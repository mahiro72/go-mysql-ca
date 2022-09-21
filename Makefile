ENV_LOCAL_FILE := .env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))

GO := go run
GO_MAIN_PATH := ./src

DB_CONTAINER_NAME:=db_go-mysql-ca

DATA_DIR:=./db/data

RM:=rm -rf

.PHONY: serve
serve:
	$(ENV_LOCAL) $(GO) $(GO_MAIN_PATH)

up:
	$(ENV_LOCAL) docker-compose -f ./docker/docker-compose.base.yml -f ./docker/docker-compose.dev.db.yml -f ./docker/docker-compose.dev.server.yml up


# docker-compose down
# imageやvolumeも削除
.PHONY: down
down:
	docker-compose -f ./docker/docker-compose.base.yml -f ./docker/docker-compose.dev.db.yml -f ./docker/docker-compose.dev.server.yml down --rmi all --volumes --remove-orphans


.PHONY: fclean
fclean:down del-volumes

.PHONY: re
re:fclean up


del-volumes:del-data

.PHONY: del-data
del-data:
	$(RM) $(DATA_DIR)


# コンテナへのアクセスをします
attach-db:
	docker exec -it $(DB_CONTAINER_NAME) sh