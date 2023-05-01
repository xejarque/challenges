package ads

import (
	"github.com/javier-tw/learning-go/internal/challenge-4/shared/domain/errors"
)

var NotFound = domainError.Create("the ad was not found")
