package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Mysql struct {
	Dialect  string `yaml:"dialect"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	ShowSql  bool   `yaml:"show_sql"`
	MaxOpen  int    `yaml:"max_open"`
	MaxIdle  int    `yaml:"max_idle"`
}

type Server struct {
	HttpPort    int    `yaml:"http_port"`
	Env         string `yaml:"env"`
	EnablePprof bool   `yaml:"enable_pprof"`
	LogLevel    string `yaml:"log_level"`
}

type Config struct {
	Server `yaml:"server"`
	Mysql  `yaml:"mysql"`
}

// 初始化：根据环境变量加载对应配置
func Init() (*Config, error) {
	// 1. 获取环境
	env := os.Getenv(EnvKeyAppEnv)
	if env == "" {
		env = DefaultEnv
	}

	// 2. 拼接路径
	configPath := ConfigDir + env + ConfigFileExt

	// 3. 读取文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败 [%s]: %w", configPath, err)
	}

	// 4. 解析 YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败 [%s]: %w", configPath, err)
	}

	// 5. 生产环境强制校验
	if env == EnvProd {
		if os.Getenv(EnvKeyDBPwd) == "" {
			return nil, fmt.Errorf("生产环境错误：DB_PASSWORD 不能为空")
		}
		cfg.Mysql.Password = os.Getenv(EnvKeyDBPwd)

		// if os.Getenv(EnvKeyRedisPwd) == "" {
		// 	return nil, fmt.Errorf("生产环境错误：REDIS_PASSWORD 不能为空")
		// }
		// cfg.Redis.PWD = os.Getenv(EnvKeyRedisPwd)
	}

	return &cfg, nil
}
