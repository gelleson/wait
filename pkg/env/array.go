package env

import (
	"fmt"
	"os"
	"strings"
)

func GetArrayString(key string, sep string) ([]string, error) {
	value := os.Getenv(key)

	if value == "" {
		return nil, fmt.Errorf("%s doesn't exist", key)
	}

	arr := strings.Split(value, sep)

	for i, element := range arr {
		arr[i] = strings.TrimSpace(element)
	}

	return arr, nil
}
