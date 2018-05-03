package blog

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/romanyx/swagger_blog_example/internal/swagger/models"
)

// Service implements swagger BlogService interface.
type Service struct {
	db *sql.DB
}

// NewService returns initialized service.
func NewService(db *sql.DB) *Service {
	s := Service{
		db: db,
	}

	return &s
}

// Update updates record in database for give request.
func (s *Service) Update(id int64, req *models.BlogUpdateRequest) error {
	columns := []string{"title", "markdown"}
	values := []interface{}{*req.Title, *req.Markdown}

	if req.PublishedAt.Present {
		if req.PublishedAt.Valid {
			values = append(values, req.PublishedAt.Value)
		} else {
			values = append(values, sql.NullInt64{})
		}
		columns = append(columns, "published_at")
	}

	values = append(values, id)

	var count uint
	if err := s.db.QueryRow("SELECT COUNT(*) FROM blogs WHERE id=?", id).Scan(&count); err != nil {
		return errors.Wrap(err, "query row")
	}

	if count == 0 {
		return sql.ErrNoRows
	}

	stm, err := s.db.Prepare(generateBlogQuery(columns))
	if err != nil {
		return errors.Wrap(err, "prepare statement failed")
	}
	_, err = stm.Exec(values...)
	if err != nil {
		return errors.Wrapf(err, "update blog %d failed", id)
	}
	return nil
}

func generateBlogQuery(fields []string) string {
	baseQuery := "UPDATE blogs SET%s WHERE id=?;"

	var setQuery string
	for i, field := range fields {
		if i+1 < len(fields) {
			setQuery = fmt.Sprintf("%s %s=?,", setQuery, field)
			continue
		}

		setQuery = fmt.Sprintf("%s %s=?", setQuery, field)
	}

	return fmt.Sprintf(baseQuery, setQuery)
}
