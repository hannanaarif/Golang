package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		// log.Println("Error while reading the file", err)
		return 1000, errors.New("failed to find file")
	}
	//fileContent := fmt.Sprintf("%s", content)
	valueText := string(content)
	//fmt.Println(fileContent)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse the stored value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, filename string) {
	valueStr := fmt.Sprintf("%.2f", value)
	os.WriteFile(filename, []byte(valueStr), 0644)
}
