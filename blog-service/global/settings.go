package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)

func SetupSetting() error {
	st, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = st.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}
	err = st.ReadSection("JWT", &JWTSetting)
	if err != nil {
		return err
	}
	JWTSetting.Expire *= time.Second
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	AppSetting.DefaultContextTimeout *= time.Second
	return nil
}

func SetupLogger() error {
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
