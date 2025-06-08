package utils

import (
	"fmt"
	"strings"

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

	// Split comma seperated ranges into slices
	ranges := strings.Split(pageRange, ",")

	// Trim whitespace in page ranges
	for i, r := range ranges {
		ranges[i] = strings.TrimSpace(r)
	}

	err := api.TrimFile(inputPath, outputPath, ranges, nil)
	if err != nil {
		return fmt.Errorf("failed to extract pages : %w", err)
	}

	return nil
}
