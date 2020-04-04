package cfg

import "awesomeProject/cookie"

func InitializeConfigurations() {
	ReadJwtProperties("./properties/jwt-configuration.json", &MainJwtConfig)
	ReadDBProperties("./properties/db-configuration.json", &MainDBConfig)
	ReadCookieProperties("./properties/cookie-configuration.json", &MainCookieConfig)
	cookie.InitializeSecureCookie(MainCookieConfig.HashKey, MainCookieConfig.BlockKey, MainCookieConfig.Path, MainCookieConfig.Domain)
}
