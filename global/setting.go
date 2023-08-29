package global

import (
	"membership_system/pkg/logger"
	"membership_system/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
)
