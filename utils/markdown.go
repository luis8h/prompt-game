package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
)

func RenderMarkdown(content string) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(content), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
