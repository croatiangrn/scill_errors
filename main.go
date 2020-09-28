package scill_errors

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	// Common errors
	GenericErr        = errors.New("generic_err")
	InvalidToken      = errors.New("invalid_token")
	InvalidDataSent   = errors.New("invalid_data_sent")
	EmptyLanguageID   = errors.New("empty_language_id")
	EmptyDBPointer    = errors.New("empty_db_pointer")
	EmptyPrimaryKey   = errors.New("empty_primary_key")
	InvalidPrimaryKey = errors.New("invalid_primary_key")
	ParseID           = errors.New("could_not_parse_id")
	CategoryIDBlank   = errors.New("category_id_cannot_be_blank")

	// GDPR Errors
	DataPrivacyNotApproved   = errors.New("data_privacy_not_approved")
	PrivacyPolicyNotAccepted = errors.New("privacy_policy_not_accepted")

	// Wargaming platform errors
	WargamingAccountTaken     = errors.New("wargaming_account_taken")
	WargamingAccountIDEmpty   = errors.New("wargaming_account_id_empty")
	WargamingAccessTokenEmpty = errors.New("wargaming_access_token_empty")

	// User errors
	UserBlocked           = errors.New("user_blocked")
	UserInactive          = errors.New("user_inactive")
	InvalidEmail          = errors.New("invalid_email")
	EmailTaken            = errors.New("email_taken")
	CouldNotGetUserAvatar = errors.New("could_not_get_user_avatar")
	WrongCredentials      = errors.New("wrong_credentials")
	Unauthorized          = errors.New("unauthorized")

	// DB Specific errors
	RecordNotFound = errors.New("record_not_found")

	CouldNotAddGamesToTheList = errors.New("could_not_add_games_to_the_list")
	GameCannotBeEmpty         = errors.New("game_cannot_be_empty")

	StartDateEmpty              = errors.New("start_date_cannot_be_empty")
	EndDateEmpty                = errors.New("end_date_cannot_be_empty")
	StartDateGreaterThanEndDate = errors.New("start_date_cannot_be_greater_than_end_date")

	EventNameBlank     = errors.New("event_name_cannot_be_blank")
	ReadMoreLinkNotURL = errors.New("read_more_link_is_not_valid")

	ChallengeNameBlank              = errors.New("challenge_name_cannot_be_blank")
	ChallengeDurationLessThanSecond = errors.New("challenge_duration_cannot_be_less_than_second")
	ChallengeGoalLessThanOne        = errors.New("challenge_goal_cannot_be_less_than_one")

	ApplicationDoesNotExist = errors.New("application_does_not_exist")
	CategoryDoesNotExist    = errors.New("category_does_not_exist")
)
