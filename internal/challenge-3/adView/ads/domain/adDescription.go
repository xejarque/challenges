package Ad

import (
	domainError "github.com/javier-tw/learning-go/internal/challenge-2/adView/shared/domain"
	"github.com/javier-tw/learning-go/internal/challenge-3/adView/shared/domain/valueObjects"
)

type Description struct {
	valueObjects.StringVO
}

var InvalidAdDescriptionError = domainError.Create("the ad Description is longer than 50 characters")

func CreateAdDescription(value string) (Description, error) {
	if len(value) > 50 {
		return Description{}, domainError.AddValue(InvalidAdDescriptionError, value)
	}
	return Description{valueObjects.StringVO{Value: value}}, nil
}
