package scill_errors

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	GenericErr                = errors.New("generic_err")
	EmailTaken                = errors.New("email_taken")
	WargamingAccountTaken     = errors.New("wargaming_account_taken")
	DataPrivacyNotApproved    = errors.New("data_privacy_not_approved")
	PrivacyPolicyNotAccepted  = errors.New("privacy_policy_not_accepted")
	UserBlocked               = errors.New("err_user_blocked")
	UserInactive              = errors.New("err_user_inactive")
	InvalidEmail              = errors.New("err_invalid_email")
	RecordNotFound            = errors.New("record_not_found")
	CouldNotAddGamesToTheList = errors.New("could_not_add_games_to_the_list")
)