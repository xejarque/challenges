package domainError

import (
	"errors"
	"fmt"
)

func AddValue(err error, value interface{}) error {
	return fmt.Errorf("Error: %w\nValue: %s\n", err, value)
}

func Create(text string) error {
	return errors.New(text)
}
