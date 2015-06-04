package far

import (
	"bufio"
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	} else {
		return true
	}
}

func FindAndReplace(path, current, update string) (int, error) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	write_file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer write_file.Close()

	writer := bufio.NewWriter(file)

	replaceCount := 0

	for _, line := range lines {
		if strings.Contains(line, current) {
			line = strings.Replace(line, current, update, -1)
			replaceCount++
		}
		write_file.WriteString(line + "\n")
	}

	writer.Flush()

	return replaceCount, nil
}
