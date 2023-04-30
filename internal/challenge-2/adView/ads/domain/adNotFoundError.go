package Ad

import (
	domainError "github.com/javier-tw/learning-go/internal/challenge-2/adView/shared/domain"
)

var NotFound = domainError.Create("the ad was not found")
