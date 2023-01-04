package paas_common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

func GetMysqlConfig(config config.Config, path ...string) *MysqlConfig {
	config.Get(path...).Scan(&MysqlConfig{})
	return &MysqlConfig{}
}
