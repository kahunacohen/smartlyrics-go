build :
	docker build --tag smartlyrics-go .
	
start :
	docker run -d -p 8080:8080 smartlyrics-go
