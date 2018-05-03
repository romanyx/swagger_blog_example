package swagger

import (
	"log"

	"github.com/romanyx/swagger_blog_example/internal/swagger/restapi/operations"
)

const (
	internalServerErrorMessage = "Internal server error"
)

// handler contents swagger API
// handlers implementations.
type handler struct {
	blogService BlogService

	logError func(err error)
}

// newHandler returns initialized Handler.
func newHandler(api *operations.SwaggerBlogExampleAPI, blogService BlogService) *handler {
	h := handler{
		blogService: blogService,
		logError: func(err error) {
			log.Printf("%s", err)
		},
	}

	h.configureAPI(api)

	return &h
}

// configureAPI sets handlers for API
func (h *handler) configureAPI(api *operations.SwaggerBlogExampleAPI) {
	//  PATCH /blog/:id
	api.BlogPatchBlogIDHandler = newPatchBlogIDHandler(h.blogService.Update, h.logError)
}
