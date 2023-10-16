package common

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WrapFunc(f interface{}, args ...interface{}) gin.HandlerFunc {
	fn := reflect.ValueOf(f)
	if fn.Type().NumIn() != len(args) {
		panic(fmt.Sprintf("函数输入参数无效：%v", fn.Type()))
	}

	outNum := fn.Type().NumOut()
	if outNum == 0 {
		panic(fmt.Sprintf("函数输出参数无效：%v，至少有一个，但是得到了：%d", fn.Type(), outNum))
	}
	inputs := make([]reflect.Value, len(args))
	for k, in := range args {
		inputs[k] = reflect.ValueOf(in)
	}
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Warnf("panic：%v", err)
				ResponseFailed(c, http.StatusInternalServerError, fmt.Errorf("%v", err))
			}
		}()

		outputs := fn.Call(inputs)
		if len(outputs) > 1 {
			err, ok := outputs[len(outputs)-1].Interface().(error)
			if ok && err != nil {
				ResponseFailed(c, http.StatusInternalServerError, err)
				return
			}
		}
		c.JSON(http.StatusOK, outputs[0].Interface())
	}
}
