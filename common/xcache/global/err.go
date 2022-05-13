package global

import (
	"fmt"
)

var (
	ErrTablerNotImplement  = fmt.Errorf("Tabler接口未实现")
	ErrIGetIDNotImplement  = fmt.Errorf("IGetID接口未实现")
	ErrInputListNotPtr     = fmt.Errorf("传入的list字段并非指针")
	ErrInputListNotSlice   = fmt.Errorf("传入的list字段并非切片")
	ErrInputModelNotPtr    = fmt.Errorf("传入的model字段并非指针")
	ErrInputModelNotStruct = fmt.Errorf("传入的model字段并非结构体")
	ErrInputMpNotMap       = fmt.Errorf("参数mp的类型不是Map")
)
