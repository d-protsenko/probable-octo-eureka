package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return data
}

type JwtConfig struct {
	Issuer               string `json:issuer`
	Secret               string `json:secret`
	RefreshTokenDuration int64  `json:refreshTokenDuration`
	AccessTokenDuration  int64  `json:accessTokenDuration`
}

var MainJwtConfig JwtConfig

func ReadJwtProperties(jwtCfgPath string, cfg *JwtConfig) {
	data := readFile(jwtCfgPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}

type DBConfig struct {
	Url         string `json:url`
	User        string `json:user`
	Password    string `json:password`
	Connections int64  `json:connections`
}

var MainDBConfig DBConfig

func ReadDBProperties(dbCfgPath string, cfg *DBConfig) {
	data := readFile(dbCfgPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}

type CookieConfig struct {
	BlockKey string `json:blockKey`
	HashKey  string `json:hashKey`
	Path     string `json:path`
	Domain   string `json:domain`
}

var MainCookieConfig CookieConfig

func ReadCookieProperties(cookieCfgPath string, cfg *CookieConfig) {
	data := readFile(cookieCfgPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}
