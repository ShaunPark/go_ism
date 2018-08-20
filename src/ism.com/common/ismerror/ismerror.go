package ismerror

import "fmt"

type IsmError struct {
	Code    int
	Message string
}

func (err *IsmError) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}
