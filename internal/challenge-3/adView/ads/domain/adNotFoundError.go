package Ad

import (
	domainError "github.com/javier-tw/learning-go/internal/challenge-3/adView/shared/domain/errors"
)

var NotFound = domainError.Create("the ad was not found")
