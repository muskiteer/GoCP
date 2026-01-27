package internals

import (
	"regexp"
	"strings"
	"github.com/ledongthuc/pdf"
	"bytes"
)

func CleanText(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = regexp.MustCompile(`\n{2,}`).ReplaceAllString(s, "\n")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}



func ExtractPDFText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)

	final := CleanText(buf.String())
	return final, nil
}