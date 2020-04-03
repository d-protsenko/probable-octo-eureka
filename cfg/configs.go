package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JwtConfig struct {
	Issuer            string `json:issuer`
	Secret            string `json:secret`
	BaseTokenDuration int64  `json:baseTokenDuration`
}

type DBConfig struct {
	Url         string `json:url`
	User        string `json:user`
	Password    string `json:password`
	Connections int64  `json:connections`
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return data
}

func ReadJwtProperties(jwtCfgPath string, cfg *JwtConfig) {
	data := readFile(jwtCfgPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func ReadDBProperties(dbCfgPath string, cfg *DBConfig) {
	data := readFile(dbCfgPath)
	err := json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
}
