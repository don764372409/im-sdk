package util

import (
	"encoding/json"
	"fmt"
)

var Debug = false

func OpenLog() {
	//开启日志
	Debug = true
}

// Out 统一的日志打印函数
func Out(str string, params ...interface{}) {
	if Debug {
		fmt.Printf(str, params)
	}
}

// Err 统一的错误日志打印函数
func Err(str string, params ...interface{}) {
	fmt.Errorf(str, params)
}
func Map2Obj(m interface{}, obj interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}
