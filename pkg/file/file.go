package file

import (
	"bufio"
	"os"
)

func ReadInput(inputFile string) (content []string, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return content, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	err = scanner.Err()

	return content, err
}
