package app

import (
    "github.com/gofiber/fiber/v2"
    "github.com/sirupsen/logrus"
    "go-webapi-template/pkg/common"
    "strconv"
    "time"
)

const (
    RFC3339Million = "2006-01-02T15:04:05.999Z07:00"
)

type LogStruct struct {
    RequestId       string
    Time            string
    Duration        int64
    StatusCode      string
    Method          string
    Path            string
    URL             string
    Host            string
    RequestBody     string
    RequestHeaders  map[string]string
    ResponseBody    string
    ResponseHeaders map[string]string
    Error           string
}

func NewLogStruct() *LogStruct {
    return &LogStruct{
        RequestId:       "",
        Time:            time.Now().UTC().Format(RFC3339Million),
        Duration:        0,
        StatusCode:      "",
        Method:          "",
        Path:            "",
        URL:             "",
        Host:            "",
        RequestBody:     "",
        RequestHeaders:  make(map[string]string),
        ResponseBody:    "",
        ResponseHeaders: make(map[string]string),
        Error:           "",
    }
}

func (this *LogStruct) ToLogrusFields() logrus.Fields {
    return logrus.Fields{
        "request_id":       this.RequestId,
        "time":             this.Time,
        "duration":         this.Duration,
        "status_code":      this.StatusCode,
        "method":           this.Method,
        "path":             this.Path,
        "url":              this.URL,
        "host":             this.Host,
        "request_body":     this.RequestBody,
        "request_headers":  this.RequestHeaders,
        "response_body":    this.ResponseBody,
        "response_headers": this.ResponseHeaders,
        "error":            this.Error,
    }
}

func ProcessRequestLoggingFuncProvider() common.CtxRequestCallbackFunc {
    return func(ctx *fiber.Ctx, data map[string]any) error {
        logS := NewLogStruct()
        logS.RequestId = string(ctx.Response().Header.Peek("X-Request-ID"))
        logS.Time = time.Now().UTC().Format(RFC3339Million)
        logS.Method = ctx.Method()
        logS.Path = ctx.Path()
        logS.URL = ctx.Request().URI().String()
        logS.Host = ctx.Hostname()
        logS.RequestBody = string(ctx.Request().Body())

        headerKeys := ctx.Request().Header.PeekKeys()
        for _, key := range headerKeys {
            logS.RequestHeaders[string(key)] = string(ctx.Request().Header.Peek(string(key)))
        }

        data["log"] = logS
        return nil
    }
}

func ProcessResponseLoggingFuncProvider(loggerProvider *LoggerProvider) common.CtxResponseCallbackFunc {
    return func(ctx *fiber.Ctx, data map[string]any) error {
        logS := data["log"].(*LogStruct)
        logS.StatusCode = strconv.Itoa(ctx.Response().StatusCode())
        logS.ResponseBody = string(ctx.Response().Body())
        logS.Duration = time.Since(ctx.Context().Time()).Milliseconds()

        headerKeys := ctx.Response().Header.PeekKeys()
        for _, key := range headerKeys {
            logS.ResponseHeaders[string(key)] = string(ctx.Response().Header.Peek(string(key)))
        }

        loggerProvider.GetLoggerEntry().WithFields(logS.ToLogrusFields()).Info("RequestResponseLogging")

        return nil
    }
}
