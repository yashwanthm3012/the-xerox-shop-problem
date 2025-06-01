package utils

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// Count the number of pages in pdf and return it
func CountPages(path string) (int, error) {
	ctx, err := api.ReadContextFile(path)

	if err != nil {
		return 0, err
	}

	return ctx.PageCount, nil
}
