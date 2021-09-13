package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
)

/*
	1)
		github.com/go-playground/locales 多种语言包, 从 CLDR 项目生成的一组多语言环境
		与 github.com/go-playground/universal-translator 配合使用
	2)
		github.com/go-playground/universal-translator 通用翻译器
	3)
		github.com/go-playground/validator/v10 validator的翻译器
	4)
		我们从请求头中的locale来判断亲请求的语种是中文还是英文,如果有其他的语种我们还需要引入其他的语言包
	5)
		最后我们通过RegisterDefaultTranslation方法, 将验证器和对应的语言在Translations中注册进来
		并将Translations设置为上下文
	6)
		在路由文件中注册中间件
*/
func Translations() gin.HandlerFunc {
	return func(context *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := context.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translation.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_translation.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_translation.RegisterDefaultTranslations(v, trans)
			}
			context.Set("trans", trans)
		}
		context.Next()
	}
}
