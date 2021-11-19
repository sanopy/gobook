package xmlnode

import (
	"fmt"
	"strings"
	"testing"
)

func TestXmlNodeFail(t *testing.T) {
	for i, test := range []string{
		"",
		"test",
		"<html>",
		"<html></body>",
	} {
		t.Run(fmt.Sprintf("case #%d", i), func(t *testing.T) {
			_, err := Parse(strings.NewReader(test))
			if err == nil {
				t.Errorf("unexpected success: %s", test)
				return
			}
			fmt.Printf("parse error: %v\n", err)
		})
	}
}
