package find

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	type args struct {
		html string
		name []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty html",
			args: args{html: "", name: []string{"div"}},
			want: ``,
		},
		{
			name: "not exist target tag name",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <h1>test header</h1>
  <p>test body.</p>
</body>
</html>`,
				name: []string{"div"}},
			want: ``,
		},
		{
			name: "exist target tag name",
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
				name: []string{"div"}},
			want: `<div target="">
    <p>test body.</p>
  </div>`,
		},
		{
			name: "multiple tag name",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <h1>h1 header</h1>
  <h2>h2 header1</h2>
	<h3>h3 header1</h3>
	<h2>h2 header2</h2>
	<h3>h3 header2</h3>
</body>
</html>`,
				name: []string{"h1", "h2", "h3"}},
			want: `<h1>h1 header</h1><h2>h2 header1</h2><h3>h3 header1</h3><h2>h2 header2</h2><h3>h3 header2</h3>`,
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
			nodes := ElementsByTagName(doc, tt.args.name...)

			// 戻り値の同値チェック
			buf := &bytes.Buffer{}
			for _, n := range nodes {
				html.Render(buf, n)
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("ElementByID(%v, %v) = %v, want %v", tt.args.html, tt.args.name, got, tt.want)
			}
		})
	}
}
