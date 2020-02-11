package schema

import (
	"errors"
	"fmt"
	"log"
)

/* 全局接口错误码定义 99开头的为系统级错误 */

// 错误码：24***
var (
	// 通用错误码
	SUCCESS         = errors.New("00000:ok")
	FAIL            = errors.New("99999:%s")
	MONGO_ERROR     = errors.New("99100:mongo error -> %s")
	TCC_VALUE_ERROR = errors.New("99201:无效tcc值 %s")

	// 业务错误码
	E24000 = errors.New("24000:points %s")
)

func Panic(userDefinedErr error, args ...interface{}) {
	if userDefinedErr == nil {
		Panic(FAIL, "Panic is nil")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	log.Printf("Panic Exception： %s", errMsg)
	panic(errors.New(errMsg))
}
