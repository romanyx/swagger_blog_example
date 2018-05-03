package blog

import (
	"strings"
	"testing"
)

func Test_generateBlogQuery(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "2 fields",
			args: args{
				fields: []string{"title", "markdown"},
			},
			want: "UPDATE blogs SET title=?, markdown=? WHERE id=?;",
		},
		{
			name: "3 fields",
			args: args{
				fields: []string{"title", "markdown", "published_at"},
			},
			want: "UPDATE blogs SET title=?, markdown=?, published_at=? WHERE id=?;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateBlogQuery(tt.args.fields); strings.Compare(got, tt.want) != 0 {
				t.Errorf("generateBlogQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
