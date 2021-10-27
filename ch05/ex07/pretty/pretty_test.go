package pretty

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrint(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty html",
			args: args{html: ""},
			want: `<html>
  <head />
  <body />
</html>
`,
		},
		{
			name: "simple html",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <h1>test header</h1>
  <p>test body.</p>
</body>
</html>`},
			want: `<!doctype html>
<html>
  <head>
    <title>
      test title
    </title>
  </head>
  <body>
    <h1>
      test header
    </h1>
    <p>
      test body.
    </p>
  </body>
</html>
`,
		},
		{
			name: "html4 document type declaration",
			args: args{html: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head><title>test title</title></head>
<body>
  <h1>test header</h1>
  <p>test body.</p>
</body>
</html>`},
			want: `<!doctype html public "-//W3C//DTD HTML 4.01 Transitional//EN" system "http://www.w3.org/TR/html4/loose.dtd">
<html>
  <head>
    <title>
      test title
    </title>
  </head>
  <body>
    <h1>
      test header
    </h1>
    <p>
      test body.
    </p>
  </body>
</html>
`,
		},
		{
			name: "an element that has no child element",
			args: args{html: `<!doctype html>
<html>
<head><title>test title</title></head>
<body>
  <img src="/test.jpg">
</body>
</html>`},
			want: `<!doctype html>
<html>
  <head>
    <title>
      test title
    </title>
  </head>
  <body>
    <img src="/test.jpg" />
  </body>
</html>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(tt.args.html))
			if err != nil {
				t.Fatalf("html parse failed: %v", err)
			}
			buf := &bytes.Buffer{}
			forEachNode(buf, doc, startElement, endElement)
			if got := buf.String(); got != tt.want {
				t.Errorf("Print(%v) = %v, want %v", tt.args.html, got, tt.want)
			}
		})
	}
}
