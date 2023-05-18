package ad

import (
	"errors"
	"fmt"
)

func AddError(err error, value interface{}) error {
	return fmt.Errorf("An error ocurred: %w\nValue: %s\n", err, value)
}

func CreateDomainError(text string) error {
	return errors.New(text)
}
