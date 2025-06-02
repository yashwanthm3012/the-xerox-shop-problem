package utils

import (
	"fmt"

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

// Function to extract the specified page ranges and return it
func ExtractPages(inputPath, outputPath, pageRange string) error {
	// Validate the page range
	if pageRange == "" {
		return fmt.Errorf("please specifiy the page ranges")
	}

	err := api.TrimFile(inputPath, outputPath, []string{pageRange}, nil)
	if err != nil {
		return fmt.Errorf("failed to extract pages: %w", err)
	}

	return err

}
