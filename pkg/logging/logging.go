package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

type LogConfig struct {
	LogDir     string `json:"log_dir"`
	LogFile    string `json:"log_file"`
	LogLevel   string `json:"log_level"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
	Format     string `json:"format"` //"json" or "text"
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// func InitLogger(path string, fileLog string) error {
func InitLogger(configPath string) error {
	//configPath := filepath.Join(path, fileLog)
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	var config LogConfig
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return err
	}
	if config.LogDir != "" {
		if _, err := os.Stat(config.LogDir); os.IsNotExist(err) {
			err := os.MkdirAll(config.LogDir, os.ModePerm)
			if err != nil {
				return err
			}
		}

	}
	logFilePath := filepath.Join(config.LogDir, config.LogFile)

	logWriter := &lumberjack.Logger{
		Filename:   logFilePath, //config.LogFile,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	formatter, err := getFormatter(config.Format)

	if err != nil {
		return err
	}
	log.SetFormatter(formatter)
	log.SetOutput(io.MultiWriter(os.Stdout, logWriter))

	return nil

}
func getFormatter(format string) (log.Formatter, error) {
	switch format {
	case "json":
		return &log.JSONFormatter{
			FieldMap: log.FieldMap{
				log.FieldKeyTime:  "time",
				log.FieldKeyLevel: "level",
				log.FieldKeyMsg:   "msg",
			},
		}, nil
	case "text":
		return &log.TextFormatter{
			FullTimestamp: true,
		}, nil
	default:
		return nil, fmt.Errorf("invalid log format: %s", format)
	}
}

func runApplication(logger *log.Logger) {
	logger.Info("Loggin started")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	logger.Info("Application stopped")
}
func logFields(handler string) log.Fields {
	return log.Fields{
		"handler": handler,
	}
}

func LogError(handler string, err error) {
	log.WithFields(logFields(handler)).Error(err)
}

func LogInfo(message string, fields log.Fields) {
	log.WithFields(fields).Info(message)
}

func LogFatal(message string, fields log.Fields) {
	log.WithFields(fields).Fatal(message)
}
