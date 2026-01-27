package internals

import (
	"errors"
	"github.com/sqweek/dialog"
)

func SelectPDF() (string, error) {
	path, err := dialog.
		File().
		Filter("PDF files", "pdf").
		Load()

	if err != nil {
		return "", errors.New("Failed to select file: " + err.Error() + "in selectPDF.go")
	}
	return path, nil
	
}
