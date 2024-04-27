package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"scale-x/dto"
	"strconv"
)

func ReadBooksFromFile(filename string) ([]dto.Book, error) {
	var books []dto.Book

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return books, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read and ignore the first line (header)
	if _, err := reader.Read(); err != nil {
		fmt.Println("Error reading header:", err)
		return books, err
	}

	// Read the remaining lines
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return books, err
	}

	for _, line := range lines {
		pubYear, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Error converting publication year to int:", err)
			continue
		}

		book := dto.Book{
			Name:            line[0],
			Author:          line[1],
			PublicationYear: pubYear,
		}
		books = append(books, book)
	}

	return books, nil
}
