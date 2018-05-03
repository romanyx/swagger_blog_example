package swagger

// newPatchBlogIDHandler returns initialized pages.PatchBlogIDHandler.
import (
	"database/sql"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/romanyx/swagger_blog_example/internal/swagger/models"
	"github.com/romanyx/swagger_blog_example/internal/swagger/restapi/operations/blog"
)

const (
	blogNotFoundMessage = "Blog not found"
	blogUpdatedMessage  = "Blog updated"
)

func newPatchBlogIDHandler(updateFunc func(int64, *models.BlogUpdateRequest) error, logError func(error)) blog.PatchBlogIDHandler {
	a := &blogPatchID{
		updateFunc: updateFunc,
		logError:   logError,
	}

	return blog.PatchBlogIDHandlerFunc(a.blogPatchID)
}

type blogPatchID struct {
	updateFunc func(int64, *models.BlogUpdateRequest) error
	logError   func(error)
}

func (a *blogPatchID) blogPatchID(params blog.PatchBlogIDParams) middleware.Responder {
	if err := a.updateFunc(params.ID, params.Body); err != nil {
		switch errors.Cause(err) {
		case sql.ErrNoRows:
			return blog.NewPatchBlogIDDefault(http.StatusNotFound).WithPayload(&models.ErrorResponse{
				Message: blogNotFoundMessage,
			})
		default:
			a.logError(errors.Wrapf(err, "patch blog id: %d", params.ID))
			return blog.NewPatchBlogIDDefault(http.StatusInternalServerError).WithPayload(&models.ErrorResponse{
				Message: internalServerErrorMessage,
			})
		}
	}

	return blog.NewPatchBlogIDOK().WithPayload(&models.MessageResponse{
		Message: blogUpdatedMessage,
	})
}
