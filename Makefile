test:
	docker-compose -p blog_test up -d --build
	go test -v --race `go list ./... | grep -v /vendor/ | grep -v /internal/swagger/models | grep -v /internal/swagger/restapi`
start:
	docker-compose -p blog_dev up -d --build
	go run cmd/serverd/main.go