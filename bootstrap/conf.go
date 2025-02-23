package bootstrap

import (
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// InitConf init config
func InitConf() {
	err := os.MkdirAll("data/temp", 0700)
	if err != nil {
		log.Fatalf("create temp dir error: %s", err.Error())
	}
	log.Infof("reading config file: %s", conf.ConfigFile)
	if !utils.Exists(conf.ConfigFile) {
		log.Infof("config file not exists, creating default config file")
		_, err := utils.CreatNestedFile(conf.ConfigFile)
		if err != nil {
			log.Fatalf("failed to create config file")
		}
		conf.Conf = conf.DefaultConfig()
		if !utils.WriteToJson(conf.ConfigFile, conf.Conf) {
			log.Fatalf("failed to create default config file")
		}
		return
	}
	config, err := ioutil.ReadFile(conf.ConfigFile)
	if err != nil {
		log.Fatalf("reading config file error:%s", err.Error())
	}
	conf.Conf = new(conf.Config)
	err = utils.Json.Unmarshal(config, conf.Conf)
	if err != nil {
		log.Fatalf("load config error: %s", err.Error())
	}
	log.Debugf("config:%+v", conf.Conf)
	// update config.json struct
	confBody, err := utils.Json.MarshalIndent(conf.Conf, "", "  ")
	if err != nil {
		log.Fatalf("marshal config error:%s", err.Error())
	}
	err = ioutil.WriteFile(conf.ConfigFile, confBody, 0777)
	if err != nil {
		log.Fatalf("update config struct error: %s", err.Error())
	}
}
