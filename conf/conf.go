package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AdminUser        string
	AdminPwd         string
	DBUrl            string
	DBUser           string
	DBPwd            string
	ServerPort       string
	DBName           string
	ChangePwdToken   string
	QiNiu_ACCESS_KEY string
	QiNiu_SECRET_KEY string
}

var ConfigContext Config

func ReadConf() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	ConfigContext = Config{}
	err := decoder.Decode(&ConfigContext)
	if err != nil {
		fmt.Println("读取配置文件 conf.json 错误错误❌")
	}
}
func init() {
	ReadConf()
	fmt.Println("read conf ")
}
