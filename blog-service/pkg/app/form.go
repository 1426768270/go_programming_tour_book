package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Message string
}

type ValidErrors []*ValidError

func(v *ValidError) Error() string{
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{})(bool, ValidErrors){
	var errs ValidErrors
	// ShouldBind 进行参数绑定和入参校验，当发生错误后，再通过上一步在中间件 Translations 设置的 Translator 来对错误消息体进行具体的翻译行为。
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		// 类型断言
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for key, value := range verrs.Translate(trans){
			errs =append(errs, &ValidError{
				Key : key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}