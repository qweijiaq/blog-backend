package core

import (
	"backend/config"
	"backend/global"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml configuration error: %v", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("unmarshal yaml configuration error: %v", err)
	}
	log.Println("yaml configuration initialized")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}

	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
