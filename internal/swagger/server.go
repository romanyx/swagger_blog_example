package swagger

import (
	"database/sql"
	"log"

	"github.com/go-openapi/loads"
	"github.com/romanyx/swagger_blog_example/internal/blog"
	"github.com/romanyx/swagger_blog_example/internal/swagger/restapi"
	"github.com/romanyx/swagger_blog_example/internal/swagger/restapi/operations"
)

// NewServer returns initialized server to run.
func NewServer(db *sql.DB, host string, port int) *restapi.Server {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalf("analyze swagger json: %s", err)
	}

	blogService := blog.NewService(db)
	api := operations.NewSwaggerBlogExampleAPI(swaggerSpec)

	newHandler(api, blogService)

	server := restapi.NewServer(api)
	server.Host = host
	server.Port = port

	server.ConfigureAPI()
	return server
}
