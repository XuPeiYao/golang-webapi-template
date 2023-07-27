package app

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-webapi-template/pkg/common"
)

type LoggerProvider struct {
	configuration common.Configuration
}

func NewLoggerProvider(configuration common.Configuration) *LoggerProvider {
	loggerProvider := &LoggerProvider{
		configuration: configuration,
	}

	return loggerProvider
}

const ConfigurationEnvironmentName = "env"

func (this *LoggerProvider) GetLoggerEntry() *logrus.Entry {
	logger := logrus.New()
	//logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetFormatter(&MyJSONFormatter{})

	return logger.WithFields(logrus.Fields{
		"env": this.configuration.GetValueOrDefault(ConfigurationEnvironmentName, "dev"),
	})
}

type MyJSONFormatter struct {
}

func (f *MyJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Note this doesn't include Time, Level and Message which are available on
	// the Entry. Consult `godoc` on information about those fields or read the
	// source of the official loggers.
	if len(entry.Message) > 0 {
		entry.Data["message"] = entry.Message
	}

	if timeVal, hasTime := entry.Data["time"]; hasTime {
		if len(timeVal.(string)) == 0 {
			entry.Data["time"] = entry.Time
		}
	}

	if errorVal, hasError := entry.Data["error"]; hasError {
		if len(errorVal.(string)) == 0 {
			delete(entry.Data, "error")
		}
	}

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %w", err)
	}
	return append(serialized, '\n'), nil
}
