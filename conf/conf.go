package conf

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

var (
	C            = Conf{}
	MongoSession *mgo.Session
)

type Conf struct {
	Server struct {
		HttpPort int `yaml:"httpPort"`
	} `server:"server"`

	AuthTokenExpireHour int `yaml:"authTokenExpireHour"`

	Mongo struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"mongo"`
}

// InitConf 读取并检查配置文件
func InitConf() (err error) {
	pwd, _ := os.Getwd()
	confFile := pwd + string(os.PathSeparator) + "conf.yaml"

	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		logrus.Errorf("read conf file err: %s, file: %s", confFile)
		os.Exit(1)
	}

	if err = yaml.Unmarshal(data, &C); err != nil {
		logrus.Errorf("parse conf file err: %s, file: %s", err.Error(), confFile)
		os.Exit(1)
	}

	if err = checkConf(&C); err != nil {
		logrus.Errorf("check conf err: %s", err.Error())
		return
	}

	return
}

func checkConf(c *Conf) error {
	if !(1 <= c.Server.HttpPort && c.Server.HttpPort <= 65535) {
		return fmt.Errorf("server.httPort无效, port=%d", c.Server.HttpPort)
	}

	if c.Mongo.Dsn == "" {
		return errors.New("mongo.dsn为空")
	}

	return nil
}

func InitDB() (err error) {
	s, err := mgo.DialWithTimeout(C.Mongo.Dsn, time.Second*3)
	if err != nil {
		logrus.Errorf("mgo.Dial err: %s", err.Error())
		return
	}

	MongoSession = s
	return
}
