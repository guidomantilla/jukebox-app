package config

type Config struct {
	Host                    *string `env:"HOST,default=localhost"`
	Port                    *string `env:"PORT,default=8080"`
	TokenSignatureKey       *string `env:"TOKEN_SIGNATURE_KEY,default=SecretYouShouldHide"`
	PasswordMinSpecialChars *string `env:"PASSWORD_MIN_SPECIAL_CHAR,default=0"`
	PasswordMinNumber       *string `env:"PASSWORD_MIN_NUMBER,default=1"`
	PasswordMinUpperCase    *string `env:"PASSWORD_MIN_UPPER_CASE,default=1"`
	PasswordLength          *string `env:"PASSWORD_LENGTH,default=8"`
	ParamHolder             *string `env:"PARAM_HOLDER,default=named"`
	DatasourceDriver        *string `env:"DATASOURCE_DRIVER,required"`
	DatasourceUsername      *string `env:"DATASOURCE_USERNAME,required"`
	DatasourcePassword      *string `env:"DATASOURCE_PASSWORD,required"`
	DatasourceServer        *string `env:"DATASOURCE_SERVER,required"`
	DatasourceService       *string `env:"DATASOURCE_SERVICE,required"`
	DatasourceUrl           *string `env:"DATASOURCE_URL,required"`
}
