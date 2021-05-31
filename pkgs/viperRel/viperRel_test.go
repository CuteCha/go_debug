package viperRel

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	vjson := viper.New()
	vjson.SetConfigType("json")
	var jsonExample = []byte(`{"exp_id": "base11", "exp_name": "live_first_alg_reRank_test11", "exp_desc": "first sort by recall type, second sort by score, priority up order", "flow": 1000, "param": {"sort": 1,"priority": 15}}`)

	err := viper.ReadConfig(bytes.NewBuffer(jsonExample))
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	t.Log(viper.GetString("exp_id"))
}

type ExpConfig struct {
	ExpId   string `json:"exp_id"`
	ExpName string `json:"exp_name"`
	ExpDesc string `json:"exp_desc"`
	Flow    string
	Param   map[string]interface{}
}

func Test02(t *testing.T) {
	vjson := viper.New()
	vjson.AddConfigPath("/Users/checao/workflow/go_debug/conf/json")
	vjson.SetConfigName("config4")
	vjson.SetConfigType("json")

	if err := vjson.ReadInConfig(); err != nil {
		log.Fatalf("Test02 read config failed: %v\n", err)
		return
	}
	var expConfig ExpConfig
	err := vjson.Unmarshal(&expConfig)
	if err != nil {
		return
	}
	t.Log(expConfig)
	t.Log(expConfig.ExpId)
	t.Log(expConfig.ExpName)
	t.Log(expConfig.ExpDesc)
	t.Log(expConfig.Flow)
	t.Log(expConfig.Param)
	t.Log(strings.Repeat("-", 72))
	t.Log(vjson.GetString("exp_id"))
	t.Log(vjson.GetString("exp_name"))
	t.Log(vjson.GetString("exp_desc"))
	t.Log(vjson.GetString("flow"))
	t.Log(vjson.GetStringMap("param"))

}

type SmtpStruct struct {
	Enable   bool
	Addr     string
	Username string
	Password string
	To       []string
}

type ExpConfig2 struct {
	Port  string
	Mysql map[string]interface{}
	Redis []string
	Smtp  SmtpStruct
}

func Test03(t *testing.T) {
	v := viper.New()
	v.SetConfigType("json") // 设置配置文件的类型

	// 配置文件内容
	var jsonExample = []byte(`
{
  "port": 10666,
  "mysql": {
    "url": "(127.0.0.1:3306)/biezhi",
    "username": "root",
    "password": "123456"
  },
  "redis": ["127.0.0.1:6377", "127.0.0.1:6378", "127.0.0.1:6379"],
  "smtp": {
    "enable": true,
    "addr": "mail_addr",
    "username": "mail_user",
    "password": "mail_password",
    "to": ["xxx@gmail.com", "xxx@163.com"]
  }
}
`)
	//创建io.Reader
	err := v.ReadConfig(bytes.NewBuffer(jsonExample))
	if err != nil {
		return
	}

	t.Log("获取配置文件的port", v.GetInt("port"))
	t.Log("获取配置文件的mysql.url", v.GetString(`mysql.url`))
	t.Log("获取配置文件的mysql.username", v.GetString(`mysql.username`))
	t.Log("获取配置文件的mysql.password", v.GetString(`mysql.password`))
	t.Log("获取配置文件的redis", v.GetStringSlice("redis"))
	t.Log("获取配置文件的smtp", v.GetStringMap("smtp"))

	var expConfig ExpConfig2
	err2 := v.Unmarshal(&expConfig)
	if err2 != nil {
		return
	}

	t.Log(expConfig)
	t.Log()
	if str, err := json.MarshalToString(expConfig); err == nil {
		t.Log(str)
	} else {
		t.Log(err.Error())
	}

}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json2.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func Test04(t *testing.T) {
	//b := []byte(`{"hello": "123"}`)
	jsonExample := []byte(`{"port": 10666,"mysql": {"url": "(127.0.0.1:3306)/biezhi","username": "root","password": "123456"},"redis": ["127.0.0.1:6377", "127.0.0.1:6378", "127.0.0.1:6379"],"smtp": {"enable": true,"addr": "mail_addr","username": "mail_user","password": "mail_password","to": ["xxx@gmail.com", "xxx@163.com"]}}`)
	b, _ := prettyPrint(jsonExample)
	fmt.Printf("%s", b)
}
