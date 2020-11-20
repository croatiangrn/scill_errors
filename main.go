package scill_errors

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	// Common errors
	GenericErr                       = errors.New("generic_err")
	InvalidToken                     = errors.New("invalid_token")
	InvalidDataSent                  = errors.New("invalid_data_sent")
	InvalidDataReceived              = errors.New("invalid_data_received")
	InvalidMediaType                 = errors.New("invalid_media_type")
	ContentTypeMustBeJSON            = errors.New("content_type_must_be_json")
	EmptyLanguageID                  = errors.New("empty_language_id")
	EmptyUserID                      = errors.New("empty_user_id")
	EmptyDBPointer                   = errors.New("empty_db_pointer")
	EmptyPrimaryKey                  = errors.New("empty_primary_key")
	InvalidPrimaryKey                = errors.New("invalid_primary_key")
	ParseID                          = errors.New("could_not_parse_id")
	CategoryIDBlank                  = errors.New("category_id_cannot_be_blank")
	HTTPRequestErr                   = errors.New("http_request_error_occurred")
	RedisDelKeyErr                   = errors.New("redis_deleting_key_error_occurred")
	InvalidURL                       = errors.New("invalid_url_provided")
	InvalidServiceName               = errors.New("invalid_service_name_provided")
	BlankEndpoint                    = errors.New("endpoint_cannot_be_blank")
	UnknownApp                       = errors.New("unknown_app")
	FieldXCannotBeBlank              = errors.New("field_x_cannot_be_blank")
	GaaSUserIDAccessTokenIDMismatch  = errors.New("gaas_user_id_access_token_id_mismatch")
	GaaSExpectedAPIKeyGotAccessToken = errors.New("gaas_expected_api_key_got_access_token")
	GaaSNoAccessTokenProvided        = errors.New("gaas_no_access_token_provided")
	ConnectionClosed                 = errors.New("connection_closed")

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

	ChallengeNameBlank                   = errors.New("challenge_name_cannot_be_blank")
	ChallengeDurationLessThanSecond      = errors.New("challenge_duration_cannot_be_less_than_second")
	ChallengeGoalLessThanOne             = errors.New("challenge_goal_cannot_be_less_than_one")
	ChallengeTypeCannotBeBlank           = errors.New("challenge_type_cannot_be_blank")
	GenericChallengeMetadataMustBeUnique = errors.New("generic_challenge_metadata_must_be_unique")

	ApplicationDoesNotExist     = errors.New("application_does_not_exist")
	CategoryDoesNotExist        = errors.New("category_does_not_exist")
	CategoryNameBlank           = errors.New("category_name_cannot_be_blank")
	InvalidCategoryTypeProvided = errors.New("invalid_category_type_provided")

	SameChallengeTypeAlreadyActive = errors.New("same_challenge_type_already_active")

	MailjetReceiverEmailBlank = errors.New("mailjet_receiver_email_blank")
	MailjetSubjectEmailBlank  = errors.New("mailjet_subject_email_blank")
	MailjetTemplateIDZero     = errors.New("mailjet_template_id_zero")

	UnknownWebhookType          = errors.New("unknown_webhook_type")
	WebhookSecretKeyMustInclude = errors.New("webhook_secret_key_must_include")

	BattlePassHasAlreadyBeenPurchased = errors.New("battle_pass_has_already_been_purchased")

	ChallengeHasAlreadyBeenActivated = errors.New("challenge_has_already_been_activated")

	// Formatted errors

	FmtWebhookSecretKeyMinLength          = errors.New("webhook_secret_key_must_be_%d_characters_long")
	FMTInvalidURLForFieldProvided         = errors.New("invalid_url_for_%q_provided")
	FmtWebhookTypeAlreadyExists           = errors.New("webhook_type_%q_already_exists")
	FmtMaxUploadSizeMB                    = errors.New("max_upload_size_for_upload_is_%d_mb")
	FmtUnknownChallengeTypeEvent          = errors.New("unknown_challenge_type_%s")
	FmtUnknownGenericChallengeMetadataKey = errors.New("generic_challenge_unknown_metadata_key_position_%d_key_%s")

	// Formatted errors -- END

	BillingProfileNotFound = errors.New("billing_profile_not_found")

	// Generic challenges errors
	GenericChallengePayloadMissingRequiredMetadata        = errors.New("generic_challenge_payload_is_missing_required_metadata")
	GenericChallengePayloadMissingSessionID               = errors.New("generic_challenge_payload_is_missing_session_id")
	GenericChallengeDamageAmountCannotBeZero              = errors.New("generic_challenge_damage_amount_cannot_be_zero")
	GenericChallengeLapPositionCannotBeZero               = errors.New("generic_challenge_lap_position_cannot_be_zero")
	GenericChallengeWeaponIDOrWeaponUsedMustContainValues = errors.New("weapon_id_or_weapon_used_must_contain_values")
	GenericChallengeCharacterNameCannotBeBlank            = errors.New("generic_challenge_character_name_cannot_be_blank")
	GenericChallengeBattleStatusCannotBeBlank             = errors.New("generic_challenge_battle_status_cannot_be_blank")
	GenericChallengeRoundIDCannotBeBlank                  = errors.New("generic_challenge_round_id_cannot_be_blank")
	GenericChallengeResourceNameCannotBeBlank             = errors.New("generic_challenge_resource_name_cannot_be_blank")
	GenericChallengeItemNameCannotBeBlank                 = errors.New("generic_challenge_item_name_cannot_be_blank")
	GenericChallengeItemTypeCannotBeBlank                 = errors.New("generic_challenge_item_type_cannot_be_blank")
	GenericChallengeCrewNameCannotBeBlank                 = errors.New("generic_challenge_crew_name_cannot_be_blank")
	GenericChallengeCoordinateXCannotBeBlank              = errors.New("generic_challenge_coordinate_x_cannot_be_blank")
	GenericChallengeCoordinateYCannotBeBlank              = errors.New("generic_challenge_coordinate_y_cannot_be_blank")
	GenericChallengeGainedResourceCannotBeBlank           = errors.New("generic_challenge_gained_resource_cannot_be_blank")
	GenericChallengeGivenResourceCannotBeBlank            = errors.New("generic_challenge_given_resource_cannot_be_blank")
	GenericChallengeAmountGainedCannotBeBlank             = errors.New("generic_challenge_amount_gained_cannot_be_blank")
	GenericChallengeAmountGivenCannotBeBlank              = errors.New("generic_challenge_amount_given_cannot_be_blank")
	GenericChallengeAnimalNameCannotBeBlank               = errors.New("generic_challenge_animal_name_cannot_be_blank")
	GenericChallengeScoreCannotBeZero                     = errors.New("generic_challenge_score_cannot_be_zero")
	GenericChallengeAmountCannotBeZero                    = errors.New("generic_challenge_amount_cannot_be_zero")
	GenericChallengeAmmoUsedCannotBeZero                  = errors.New("generic_challenge_ammo_used_cannot_be_zero")
	GenericChallengeUnitTypeCannotBeBlank                 = errors.New("generic_challenge_unit_type_cannot_be_blank")
	GenericChallengePlayerCharacterCannotBeBlank          = errors.New("generic_challenge_player_character_cannot_be_blank")
	GenericChallengeRequiredTimeOrDurationMissing         = errors.New("generic_challenge_required_time_or_duration_missing")
	GenericChallengeRequiredTimeMissing                   = errors.New("generic_challenge_required_time_missing")
	GenericChallengeDistanceCannotBeZero                  = errors.New("generic_challenge_distance_cannot_be_zero")
	GenericChallengeCardTypeCannotBeZero                  = errors.New("generic_challenge_card_type_cannot_be_zero")
	GenericChallengeEffectIDCannotBeBlank                 = errors.New("generic_challenge_effect_id_cannot_be_blank")
	GenericChallengeWeaponTypeCannotBeBlank               = errors.New("generic_challenge_weapon_type_cannot_be_blank")
	GenericChallengeBountyNameCannotBeBlank               = errors.New("generic_challenge_bounty_name_cannot_be_blank")
	GenericChallengeBuildingIDCannotBeBlank               = errors.New("generic_challenge_building_id_cannot_be_blank")
	GenericChallengeTransportIDCannotBeBlank              = errors.New("generic_challenge_transport_id_cannot_be_blank")
	GenericChallengeHomeTeamIDCannotBeBlank               = errors.New("generic_challenge_home_team_id_cannot_be_blank")
	GenericChallengeAwayTeamIDCannotBeBlank               = errors.New("generic_challenge_away_team_id_cannot_be_blank")
	GenericChallengeHomeTeamScoreCannotBeBlank            = errors.New("generic_challenge_home_team_score_cannot_be_blank")
	GenericChallengeAwayTeamScoreCannotBeBlank            = errors.New("generic_challenge_away_team_score_cannot_be_blank")
	GenericChallengePuzzleIDCannotBeBlank                 = errors.New("generic_challenge_puzzle_id_cannot_be_blank")
	GenericChallengeActionTypeCannotBeBlank               = errors.New("generic_challenge_action_type_cannot_be_blank")
	GenericChallengeMissionIDCannotBeBlank                = errors.New("generic_challenge_mission_id_cannot_be_blank")
	GenericChallengeCheckpointIDCannotBeBlank             = errors.New("generic_challenge_checkpoint_id_cannot_be_blank")
	GenericChallengeLevelIDCannotBeBlank                  = errors.New("generic_challenge_level_id_cannot_be_blank")
	GenericChallengeItemIDCannotBeBlank                   = errors.New("generic_challenge_item_id_cannot_be_blank")
	GenericChallengeMapNameCannotBeBlank                  = errors.New("generic_challenge_map_name_cannot_be_blank")
	GenericChallengeMapSectionNameCannotBeBlank           = errors.New("generic_challenge_map_section_name_cannot_be_blank")
	GenericChallengeSomeMetadataKeysAreNotAllowed         = errors.New("generic_challenge_some_metadata_keys_are_not_allowed")
	GenericChallengeSomeMetadataHasInvalidTypes           = errors.New("generic_challenge_some_metadata_has_invalid_types")
)
