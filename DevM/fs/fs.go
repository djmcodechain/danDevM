package fs

import "os"

func FileReader(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
