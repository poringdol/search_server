DB_CONTAINER = yandex-food-mongodb
SEARCH_SERVER = search-server
PROJECT_PATH = /home/cddoma/GolandProjects/yandexFood

all: db search-server web
#back front

db:
	echo $(PROJECT_PATH)
	( docker ps | grep $(DB_CONTAINER) ) || \
	(( docker ps -a | grep $(DB_CONTAINER) ) && docker start $(DB_CONTAINER) ) || \
	( docker run -d -p 27017:27017 --volume ~/db_volume:/data/db --name $(DB_CONTAINER) mongo:4.4 )

search-server:
	( ls | grep $(SEARCH_SERVER) && gnome-terminal --window -- ./$(SEARCH_SERVER) http ) || \
	( echo "Building search server..." && \
	go build -o ./$(SEARCH_SERVER) ./server && pwd && ls && gnome-terminal --window -- ./$(SEARCH_SERVER) http )

web:
	cd web && npm start


stop: stop.web stop.search-server stop.db

stop.web:
	pkill start || true

stop.search-server:
	pkill $(SEARCH_SERVER) || true

stop.db:
	docker stop $(DB_CONTAINER) || true


rm: rm.search.server rm.db

rm.search.server:
	rm $(SEARCH_SERVER) || true

rm.db:
	docker stop $(DB_CONTAINER) || true
	docker rm $(DB_CONTAINER) || true

re: rm all

.PHONY: web db search-server
