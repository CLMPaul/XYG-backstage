package e

import "fmt"

type PhonePatternError struct {
	Phone string
}

func (e PhonePatternError) Error() string {
	return fmt.Sprintf("phone pattern error:%s", e.Phone)
}

type ParmError struct {
	Parm string
}

func (e ParmError) Error() string {
	return fmt.Sprintf("参数错误:%s", e.Parm)
}
