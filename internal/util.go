package internal

import "os"

func writeFileContent(path string, content []byte) error {
	return os.WriteFile(path, content, 0755)
}
