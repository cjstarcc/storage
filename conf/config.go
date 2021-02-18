package conf

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Config struct {
	Env struct {
		Port int    `yaml:"port"`
		Dir  string `yaml:"dir"`
		Rabbitmq string `yaml:"rabbitmq"`
	} `yaml:"env"`
}

var conf Config

func InitConfig() error {
	yamlFile,err := ioutil.ReadFile("depoly/config.yaml")
	if err!=nil{
		err = errors.Wrap(err,"read config file failed.")
		return err
	}
	err = yaml.Unmarshal(yamlFile,&conf)
	if err!=nil{
		err = errors.Wrap(err, "Unmarshal config file failed.")
		return err
	}
	log.Print(conf)
	return nil
}

func GetConfig() Config{
	return conf
}