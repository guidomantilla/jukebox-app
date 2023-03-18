package config

const (

	// Hardcoded Values

	APPLICATION_NAME         = "oracle-api-orchestrator"
	CMD_PROPERTY_SOURCE_NAME = "CMD"
	OS_PROPERTY_SOURCE_NAME  = "OS"
	LOG_LEVEL                = "LOG_LEVEL"
	SENTRY_DSN               = "SENTRY_DSN"
	SENTRY_ENVIRONMENT       = "SENTRY_ENVIRONMENT"
	SENTRY_RELEASE           = "SENTRY_RELEASE"

	// Mandatory Values

	DATASOURCE_DRIVER   = "DATASOURCE_DRIVER"
	DATASOURCE_USERNAME = "DATASOURCE_USERNAME"
	DATASOURCE_PASSWORD = "DATASOURCE_PASSWORD"
	DATASOURCE_SERVER   = "DATASOURCE_SERVER"
	DATASOURCE_SERVICE  = "DATASOURCE_SERVICE"
	DATASOURCE_URL      = "DATASOURCE_URL"
	ORACLE_API_URL      = "ORACLE_API_URL"
	ORACLE_API_USERNAME = "ORACLE_API_USERNAME"
	ORACLE_API_PASSWORD = "ORACLE_API_PASSWORD"

	// Optional Values

	HOST_PORT                 = "HOST_PORT"
	SIGNING_KEY               = "SIGNING_KEY"
	BCRYPT_COST               = "BCRYPT_COST"
	PASSWORD_MIN_SPECIAL_CHAR = "PASSWORD_MIN_SPECIAL_CHAR"
	PASSWORD_MIN_NUMBER       = "PASSWORD_MIN_NUMBER"
	PASSWORD_MIN_UPPER_CASE   = "PASSWORD_MIN_UPPER_CASE"
	PASSWORD_LENGTH           = "PASSWORD_LENGTH"
)

var (
	ENV_VAR_UNDEFINED_MESSAGES_MAP = map[string]string{
		DATASOURCE_DRIVER:   "server starting up - error setting up DB connection: empty driver name",
		DATASOURCE_USERNAME: "server starting up - error setting up DB connection: empty username",
		DATASOURCE_PASSWORD: "server starting up - error setting up DB connection: empty password",
		DATASOURCE_SERVER:   "server starting up - error setting up DB connection: empty server",
		DATASOURCE_SERVICE:  "server starting up - error setting up DB connection: empty service",
		DATASOURCE_URL:      "server starting up - error setting up DB connection: empty url",
	}

	ENV_VAR_DEFAULT_VALUES_MAP = map[string]string{
		HOST_PORT:                 ":8080",
		SIGNING_KEY:               "some_random_secure_long_key",
		BCRYPT_COST:               "15",
		PASSWORD_MIN_SPECIAL_CHAR: "0",
		PASSWORD_MIN_NUMBER:       "1",
		PASSWORD_MIN_UPPER_CASE:   "1",
		PASSWORD_LENGTH:           "8",
	}
)
