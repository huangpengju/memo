// conf 包是处理配置数据
package conf

import (
	"fmt"
	"memo/model"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string // 应用程序模式
	HttpPort   string // 服务器端口
	Db         string // 数据库
	DbHost     string // 数据库主机
	DbPort     string // 数据库端口
	DbUser     string // 数据库用户名
	DbPassWord string // 数据库密码
	DbName     string // 数据库库名
)

func Init() {
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Printf("读取配置文件config.ini出错: %v\n", err)
		fmt.Println("请检查路径")
		os.Exit(1)
	}
	// 读取 操作分区中的键值
	// 获取 service 服务器配置
	LoadServer(cfg)
	// redis

	// 获取 mysql 数据库配置
	LoadMysql(cfg)

	// 使用变量值 DbUser(root) DbPassWord(root) DbHost(127.0.0.1) DbPort(3306) DbName(memo)
	// 拼接出 Mysql 连接路径 DSN（Data Source Name 表示数据库连接来源,用于定义如何连接数据库）
	// func Join(a []string, sep string) string 将一系列字符串连接为一个字符串，之间用sep来分隔。

	// 用户名[:密码]@][协议(数据库服务器地址)]]/数据库名称?参数列表
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	//  var DSN = "root:root@tcp(localhost:3306)/memo?charset=utf8&parseTime=True&loc=Local"
	DSN := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	// 连接数据库
	model.Database(Db, DSN)

}

// LoadServer 获取配置文件中操作分区为 service 中的键值（获取服务器配置）
func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("service").Key("AppMode").String()
	HttpPort = cfg.Section("service").Key("HttpPort").String()
}

// func LoadRedis(cfg *ini.File){

// }

// LoadMysql 获取配置文件中操作分区为 mysql 中的键值（获取数据库配置）
func LoadMysql(cfg *ini.File) {
	Db = cfg.Section("mysql").Key("Db").String()
	DbHost = cfg.Section("mysql").Key("DbHost").String()
	DbPort = cfg.Section("mysql").Key("DbPort").String()
	DbUser = cfg.Section("mysql").Key("DbUser").String()
	DbPassWord = cfg.Section("mysql").Key("DbPassWord").String()
	DbName = cfg.Section("mysql").Key("DbName").String()
}
