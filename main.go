package scill_errors

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	// Common errors
	GenericErr = errors.New("generic_err")

	// GDPR Errors
	DataPrivacyNotApproved   = errors.New("data_privacy_not_approved")
	PrivacyPolicyNotAccepted = errors.New("privacy_policy_not_accepted")

	// Wargaming platform errors
	WargamingAccountTaken     = errors.New("wargaming_account_taken")
	WargamingAccountIDEmpty   = errors.New("wargaming_account_id_empty")
	WargamingAccessTokenEmpty = errors.New("wargaming_access_token_empty")

	// User errors
	UserBlocked  = errors.New("user_blocked")
	UserInactive = errors.New("user_inactive")
	InvalidEmail = errors.New("invalid_email")
	EmailTaken   = errors.New("email_taken")

	// DB Specific errors
	RecordNotFound = errors.New("record_not_found")

	CouldNotAddGamesToTheList = errors.New("could_not_add_games_to_the_list")
)
