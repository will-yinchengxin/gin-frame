package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, i2 := range v {
		errs = append(errs, i2.Error())
	}
	return errs
}

/*
	1)
		对入参的校验方法进行了二次封装,在BindAndValid方法中,通过shouldBind进行绑定校验的
	2)
		当发生错误后,再通过Translations中设置的Translator对错误信息进行翻译
*/
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		// 中间件注册的trans
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for i, i2 := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: i,
				Message: i2,
			})
		}
		return false, errs
	}
	return true, nil
}
