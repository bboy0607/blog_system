package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	ListenAddr   string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdelConns int
	MaxOpenConns int
}

type EmailSettingS struct {
	Host     string
	Port     int
	Username string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type RedisSettingS struct {
	Host string
	Port string
}

// 讀取設定檔區段，並存入指定結構體的方法
func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
