package swagger

import (
	"github.com/romanyx/swagger_blog_example/internal/swagger/models"
)

// BlogService service inteface for handling blog requests.
type BlogService interface {
	Update(int64, *models.BlogUpdateRequest) error
}
