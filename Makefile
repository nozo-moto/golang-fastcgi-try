build:
	go build -o main
docker-build:
	docker build -t nozomi0966/fastcgi-golang-try -f ./docker/fcgi/Dockerfile .
clean:
	rm ./main
