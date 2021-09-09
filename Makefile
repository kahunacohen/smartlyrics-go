dev_up :
	docker-compose -f docker-compose.dev.yaml up 
dev_up_build :
	docker-compose -f docker-compose.dev.yaml up --build 
dev_down	:
	docker-compose -f docker-compose.dev.yaml down
go_build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin
watch :
	fswatch -e bin/ . | xargs -n1 -I{} make handle_watch
handle_watch	: go_build
	docker-compose -f docker-compose.dev.yaml restart web
ls_db_data	:
	docker exec $(CT_ID) ls /var/lib/postgresql/data
rm_db_data	:
	docker exec $(CT_ID) sh -c "rm -rf /var/lib/postgresql/data/*"

