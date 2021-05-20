package viperRel

import (
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type UserInfo struct {
	UserName string
	Address  string
	Sex      byte
	Company  Company
}

type Company struct {
	Name       string
	EmployeeId int
	Department []interface{}
}

func ViperYamlTest11() {
	v := viper.New()
	v.AddConfigPath("./conf/yaml")
	v.SetConfigName("config1")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf("userName:%s sex:%s company.name:%s \n",
		v.Get("userName"), v.Get("sex"), v.Get("company.name"))

	var userInfo UserInfo
	if err := v.Unmarshal(&userInfo); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println(userInfo)

}

func ViperCommandLineTest12() {
	pflag.String("ip", "127.0.0.1", "Server running address")
	pflag.Int64("port", 8080, "Server running port")
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return
	}
	fmt.Printf("ip :%s , port:%s\n",
		viper.GetString("ip"), viper.GetString("port"))

}

type ViperJsonConfig struct {
	Redis string
	MySQL ViperMySQLConfig
}

type ViperMySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

func ViperJsonTest13() {
	vjson := viper.New()
	vjson.AddConfigPath("./conf/json")
	vjson.SetConfigName("config2")
	vjson.SetConfigType("json")

	if err := vjson.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	var viperJsonConfig ViperJsonConfig
	err := vjson.Unmarshal(&viperJsonConfig)
	if err != nil {
		return
	}
	fmt.Println("read config3.json")
	fmt.Println("config: ", viperJsonConfig, "redis: ", viperJsonConfig.Redis)
	fmt.Println(viperJsonConfig)

	if viperJson, err := json.MarshalToString(viperJsonConfig); err == nil {
		fmt.Println(viperJson)
	} else {
		fmt.Println(err.Error())
	}

}
