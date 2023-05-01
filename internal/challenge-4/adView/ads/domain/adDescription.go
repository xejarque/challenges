package Ad

import (
	domainError "github.com/javier-tw/learning-go/internal/challenge-2/adView/shared/domain"
	"github.com/javier-tw/learning-go/internal/challenge-4/shared/domain/valueObjects"
)

type Description struct {
	valueObjects.StringVO
}

var InvalidAdDescriptionError = domainError.Create("the ad Description is longer than 50 characters")

func CreateDescription(description string) (Description, error) {
	if len(description) > 50 {
		return Description{}, domainError.AddValue(InvalidAdDescriptionError, description)
	}
	return Description{valueObjects.CreateStringVO(description)}, nil
}

func DescriptionFromPrimitives(value string) Description {
	return Description{valueObjects.CreateStringVO(value)}
}
