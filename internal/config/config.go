package config

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"strings"

	model "github.com/lin-snow/ech0/internal/model/common"
	"github.com/spf13/viper"
)

// Config 全局配置变量
var Config AppConfig

// JWT_SECRET 用于JWT签名的密钥
var JWT_SECRET []byte

// AppConfig 应用程序配置结构体
type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	Database struct {
		Type string `yaml:"type"`
		Path string `yaml:"path"`
		// Pragma string `yaml:"pragma"`
	} `yaml:"database"`
	Auth struct {
		Jwt struct {
			Expires  int    `yaml:"expires"`
			Issuer   string `yaml:"issuer"`
			Audience string `yaml:"audience"`
		} `yaml:"jwt"`
	} `yaml:"auth"`
	Upload struct {
		ImageMaxSize int      `yaml:"imagemaxsize"`
		AudioMaxSize int      `yaml:"audiomaxsize"`
		AllowedTypes []string `yaml:"allowedtypes"`
		ImagePath    string   `yaml:"imagepath"`
		AudioPath    string   `yaml:"audiopath"`
	} `yaml:"upload"`
	Setting struct {
		SiteTitle     string `yaml:"sitetitle"`
		Servername    string `yaml:"servername"`
		Serverurl     string `yaml:"serverurl"`
		AllowRegister bool   `yaml:"allowregister"`
		Icpnumber     string `yaml:"icpnumber"`
		MetingAPI     string `yaml:"metingapi"`
		CustomCSS     string `yaml:"customcss"`
		CustomJS      string `yaml:"customjs"`
	}
}

func LoadAppConfig() {
	if os.Getenv("RUN_ON") == "kubernetes" {
		loadConfigFromEnvContent()
	} else {
		loadConfigFromFile()
	}

	JWT_SECRET = GetJWTSecret()
}

func loadConfigFromFile() {
	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}
}

func loadConfigFromEnvContent() {
	configContent := os.Getenv("CONFIG_FILE_CONTENT")
	if configContent == "" {
		panic(model.READ_CONFIG_PANIC + ": CONFIG_FILE_CONTENT is empty")
	}

	viper.SetConfigType("yaml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}
}


// GetJWTSecret 加载JWT密钥
func GetJWTSecret() []byte {
	// 从环境变量中获取JWT密钥
	secret := os.Getenv("JWT_SECRET")
	if secret == "" { // 如果没有设置环境变量，则使用UUID生成默认密钥
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal("failed to generate random JWT secret:", err)
		}
		secret = hex.EncodeToString(b)
	}

	return []byte(secret)
}
