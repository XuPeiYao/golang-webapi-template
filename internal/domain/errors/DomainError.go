package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type DomainError struct {
	StatusCode   int
	ErrorCode    string
	Message      string
	InnerError   error
	StackMessage string
}

func NewDomainError(err error, statusCode int, errorCode string, message string) *DomainError {
	if err == nil {
		err = fmt.Errorf("")
	}

	return &DomainError{
		InnerError:   err,
		StatusCode:   statusCode,
		ErrorCode:    errorCode,
		Message:      message,
		StackMessage: fmt.Sprintf("%+v", errors.Wrap(err, "")),
	}
}

func (this *DomainError) Error() string {
	if this.InnerError == nil {
		return fmt.Sprintf("[%s] %s", this.ErrorCode, this.Message)
	}

	return fmt.Sprintf("[%s] %s; %s", this.ErrorCode, this.Message, this.InnerError.Error())
}

func (this *DomainError) WithMessage(message string) *DomainError {
	return NewDomainError(this.InnerError, this.StatusCode, this.ErrorCode, message)
}

func (this *DomainError) WithStatusCode(statusCode int) *DomainError {
	return NewDomainError(this.InnerError, statusCode, this.ErrorCode, this.Message)
}

func (this *DomainError) WithErrorCode(errorCode string) *DomainError {
	return NewDomainError(this.InnerError, this.StatusCode, errorCode, this.Message)
}

func (this *DomainError) Wrap(
	statusCode int,
	message string,
) *DomainError {
	return &DomainError{
		StatusCode: statusCode,
		Message:    message,
		InnerError: this,
	}
}

func (this *DomainError) Unwrap() error {
	return this.InnerError
}

func (this *DomainError) Serialize() map[string]any {
	result := make(map[string]any)
	result["code"] = this.ErrorCode
	result["message"] = this.Message
	result["stack"] = this.StackMessage

	if this.InnerError != nil {
		if innerError, ok := this.InnerError.(interface {
			Serialize() map[string]any
		}); ok {
			result["innerError"] = innerError.Serialize()
		}
	}

	return result
}

func (this *DomainError) New(err error) *DomainError {
	return NewDomainError(err, this.StatusCode, this.ErrorCode, this.Message)
}

var (
	ServerError     = NewDomainError(nil, 500, "S-001", "Service Error")
	HelloWorldError = NewDomainError(nil, 400, "S-002", "HelloWorldService Error")
)
