package common

import "github.com/micro/go-micro/v2/config"

/**
consul配置中心
{
	"host":"127.0.0.1",
  "user":"root",
  "pwd":"",
  "database":"twj_test",
  "port":3306
}
*/
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// 获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfg := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfg)
	return mysqlConfg
}
