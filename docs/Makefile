generate:
	$(GOPATH)/bin/swagger generate server \
	-f $(GOPATH)/src/github.com/romanyx/swagger_blog_example/docs/swagger.json \
	-t $(GOPATH)/src/github.com/romanyx/swagger_blog_example/internal/swagger \
	-A swagger_blog_example \
	--tags=blog \
	-M BlogUpdateRequest \
	-M MessageResponse \
	-M ErrorResponse \
	--exclude-main
validate:
	$(GOPATH)/bin/swagger validate $(GOPATH)/src/github.com/romanyx/swagger_blog_example/docs/swagger.json
start:
	docker-compose up -d
restart:
	docker-compose restart
