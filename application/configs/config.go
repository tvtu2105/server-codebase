package configs

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tvtu2105/go-commons/files"
	"github.com/tvtu2105/go-commons/utils"
	"sever-codebase/application/models"
)

var cfg *models.Config

func LoadConfig(path, env string) (*models.Config, error) {
	//verify path
	exist := files.CheckFolderExist(path)
	if !exist {
		return nil, errors.New("config path is not exist")
	}
	var configFile string
	if utils.IsEmpty(env) {
		env = "dev"
	}
	if utils.IsNotEmpty(path) {
		configFile = "application" + "-" + env + ".yaml"
	}
	exist = files.CheckFolderExist(path)
	if !exist {
		return nil, errors.New("config file is not exist")
	}

	//read config from path to cfg
	viper.AddConfigPath(path)
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadEnv() (string, error) {

	return "", nil
}
