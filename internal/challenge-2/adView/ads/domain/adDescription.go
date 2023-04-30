package Ad

import (
	domainError "github.com/javier-tw/learning-go/internal/challenge-2/adView/shared/domain"
)

type Description struct {
	value string
}

var InvalidAdDescriptionError = domainError.Create("the ad description is longer than 50 characters")

func CreateDescription(value string) (Description, error) {
	if len(value) > 50 {
		return Description{}, domainError.AddValue(InvalidAdDescriptionError, value)
	}
	return Description{value: value}, nil
}

func (adDescription Description) String() string {
	return adDescription.value
}
