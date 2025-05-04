package clerr

import "errors"

const (
	errorServer = "server error"
)

// Auth error
const (
	ErrorInvalidLoginOrPassword       = "invalid login or password"
	ErrorExistAccount                 = "an account with this email already exists"
	ErrorAccountIsConfirmed           = "account verified, re-verification is not required"
	ErrorConfirmAccountTimeoutExpired = "time allotted for account verification has expired"
	ErrorConfirmCodeNotMatched        = "expected code did not match actual code"
	ErrorAccountIsDeleted             = "account deleted"
	ErrorConfirmCodeNotValid          = "verification code is not valid"
)

// Account error
const (
	ErrorUnRegisteredAccount    = "account with this email not found"
	ErrorInvalidAccountStatus   = "invalid account status"
	ErrorOrganizationNotMatches = "the user is in a different organization"
	ErrorChangingOnesStatus     = "it is forbidden to change your status"
)

// Employee error
const (
	ErrorAccountStatusDenied   = "this action is prohibited for the current account status"
	ErrorAccountExpectsConfirm = "an account with this email has already been registered and is awaiting confirmation"
)

// File error
const (
	ErrorInvalidFile = "invalid file"
	ErrorSaveFile    = "the file was not saved due to an unknown error"
)

var ErrorServer = errors.New(errorServer)
