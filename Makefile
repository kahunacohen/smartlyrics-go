up_dev :
	docker-compose -f docker-compose.dev.yaml up 
up_build_dev :
	docker-compose -f docker-compose.dev.yaml up --build 
down_dev	:
	docker-compose -f docker-compose.dev.yaml down
go_build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin
watch :
	fswatch -e bin -e git -o . | xargs -n1 -I{} make handle_watch
handle_watch	: go_build
	docker-compose -f docker-compose.dev.yaml restart web

