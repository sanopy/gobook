package find

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrint(t *testing.T) {
	type args struct {
		html string
		id   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty html",
			args: args{html: "", id: "target"},
			want: ``,
		},
		{
			name: "not exist target id",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <h1>test header</h1>
  <p>test body.</p>
</body>
</html>`,
				id: "target"},
			want: ``,
		},
		{
			name: "exist target id",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <h1>test header</h1>
  <div target="">
    <p>test body.</p>
  </div>
</body>
</html>`,
				id: "target"},
			want: `<div target="">
    <p>test body.</p>
  </div>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のDOM準備
			doc, err := html.Parse(strings.NewReader(tt.args.html))
			if err != nil {
				t.Fatalf("html parse failed: %v", err)
			}

			// テスト対象関数呼び出し
			n := ElementByID(doc, tt.args.id)

			// 戻り値の同値チェック
			buf := &bytes.Buffer{}
			if n != nil {
				html.Render(buf, n)
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("ElementByID(%v, %v) = %v, want %v", tt.args.html, tt.args.id, got, tt.want)
			}
		})
	}
}
