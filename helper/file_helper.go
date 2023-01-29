package helper

import "strings"

type FileHelper struct {

}

// Function to get line number that contains a string
func (fh *FileHelper) GetLineNumberThatContainsString(fileContent string, stringToFind string) int {
	lines := strings.Split(fileContent, "\n")

	for i, line := range lines {
		if strings.Contains(line, stringToFind) {
			return i
		}
	}

	return -1
}

// Function to remove line from file content
func (fh *FileHelper) RemoveLineFromFileContent(fileContent string, lineNumber int) string {
	lines := strings.Split(fileContent, "\n")
	lines = append(lines[:lineNumber], lines[lineNumber+1:]...)
	return strings.Join(lines, "\n")
}