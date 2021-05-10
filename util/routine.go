// util 包， 该包包含了项目共用的一些常量，封装了项目中一些共用函数。
// 创建人： dfklegend
// 创建时间： 20210510
package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetRoutineID  
//		获取当前routine的id
// 		用于测试
// 参数：
//      无
// 返回值：
//      int routine的id
func GetRoutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	//str := string(buf[:n])
	//fmt.Println("str: %v", str)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
