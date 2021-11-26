package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSetting struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type UploadFile struct {
	UploadSavePath     string
	UploadServerUrl    string
	UploadImageMaxSize int
	UploadImageAllExts []string
}

type JaegerSetting struct {
	Host string
	Name string
}

type RedisSetting struct {
	MaxIdleNum     int    `yaml:"maxIdleNum"`
	MaxActive      int    `yaml:"maxActive"`
	MaxIdleTimeout int    `yaml:"maxIdleTimeout"`
	ConnectTimeout int    `yaml:"connectTimeout"`
	ReadTimeout    int    `yaml:"readTimeout"`
	Host           string `yaml:"host"`
	Password       string `yaml:"password"`
	Database       int    `yaml:"database"`
}
