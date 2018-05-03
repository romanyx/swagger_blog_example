package swagger

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	blogPatchOKRequest = `{
		"title": "Handling JSON null or missing values with go-swagger",
		"markdown": "In JSON there are two ways to represent null values, first is when there is no value for given key and it’s value implicitly set to null, and second is when key is present and it’s value is explicitly set to null.",
		"published_at": 123
	}`
	blogPatchOKRequestWithNullPublishedAt = `{
		"title": "Handling JSON null or missing values with go-swagger",
		"markdown": "In JSON there are two ways to represent null values, first is when there is no value for given key and it’s value implicitly set to null, and second is when key is present and it’s value is explicitly set to null.",
		"published_at": null
	}`
	blogPatchOKRequestWithoutPublishedAt = `{
		"title": "Handling JSON null or missing values with go-swagger",
		"markdown": "In JSON there are two ways to represent null values, first is when there is no value for given key and it’s value implicitly set to null, and second is when key is present and it’s value is explicitly set to null."
	}`
	blogPatchUnprocessableEntityRequest = `{
		"title": "Handling JSON null or missing values with go-swagger",
		"published_at": 123
	}`
)

func patchBlogIDSuite() func() {
	dbCleaner.Acquire("blogs")

	db.Exec("INSERT INTO blogs (title, markdown, published_at) VALUES (?, ?, ?);", "Title", "Body", 1)

	return func() {
		dbCleaner.Clean("blogs")
	}
}

func TestPatchBlogID(t *testing.T) {
	handler := testHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	tests := []struct {
		name        string
		code        int
		request     string
		publishedAt sql.NullInt64
		matches     []string
	}{
		{
			name:        "200 case 1",
			code:        http.StatusOK,
			request:     blogPatchOKRequest,
			publishedAt: sql.NullInt64{Valid: true, Int64: 123},
			matches: []string{
				"Blog updated",
			},
		},
		{
			name:        "200 case 2",
			code:        http.StatusOK,
			request:     blogPatchOKRequestWithNullPublishedAt,
			publishedAt: sql.NullInt64{},
			matches: []string{
				"Blog updated",
			},
		},
		{
			name:        "200 case 3",
			code:        http.StatusOK,
			request:     blogPatchOKRequestWithoutPublishedAt,
			publishedAt: sql.NullInt64{Valid: true, Int64: 1},
			matches: []string{
				"Blog updated",
			},
		},
		{
			name:    "422",
			code:    http.StatusUnprocessableEntity,
			request: blogPatchUnprocessableEntityRequest,
			matches: []string{
				"code\":602",
				"message\":\"markdown in body is required",
			},
		},
		{
			name:    "404",
			code:    http.StatusNotFound,
			request: blogPatchOKRequest,
			matches: []string{
				"Blog not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer patchBlogIDSuite()()
			id := 1
			if tt.code == http.StatusNotFound {
				id = 333
			}
			res := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", fmt.Sprintf("%s/blog/%d", server.URL, id), strings.NewReader(tt.request))
			req.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(res, req)

			var showBody bool
			if res.Code != tt.code {
				showBody = true
				t.Errorf("expected code %d got %d", tt.code, res.Code)
			}

			var publishedAt sql.NullInt64
			if err := db.QueryRow("SELECT published_at FROM blogs WHERE id=?", id).Scan(&publishedAt); err != nil {
				if publishedAt != tt.publishedAt {
					t.Errorf("expected published at to be %v got %v", tt.publishedAt, publishedAt)
				}
			}

			body := res.Body.String()
			for _, match := range tt.matches {
				if !strings.Contains(body, match) {
					showBody = true
					t.Errorf("body does not contains match: %s", match)
				}
			}

			if showBody {
				t.Errorf("body: %s", body)
			}
		})
	}
}
