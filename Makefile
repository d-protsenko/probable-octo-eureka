MONGO_DB_CONTAINER_NAME=awesome_mongo
DB_NAME=awesome_db
DB_USER=awesome_user
DB_USER_PWD=awesome_user_password
AUTH_DB=admin

.PHONY: start-db

start-db:
	docker-compose up -d

.PHONY: connect-to-db

connect-to-db:
	docker exec -it ${MONGO_DB_CONTAINER_NAME} bash -c "mongo --username ${DB_USER} --password ${DB_USER_PWD} --authenticationDatabase ${AUTH_DB} mongodb://localhost:27017/${DB_NAME}"
