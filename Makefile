up_build_dev :
	docker-compose -f docker-compose.dev.yaml up --build 
	
go_build :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
