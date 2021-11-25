package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"time"
)

func (v *ValidatorX) Register() {
	_ = v.Validate.RegisterValidation("urlFormat", CheckUrl)
	_ = v.Validate.RegisterValidation("passwordFormat", Password)
	_ = v.Validate.RegisterValidation("phoneFormat", Phone)
	_ = v.Validate.RegisterValidation("ltNowTime", LtNowTime)

	// 添加 RegisterTranslation
	v.RegisterTranslation()
}

/*
urlFormat 		校验url
passwordFormat  校验password
phoneFormat		校验手机号
ltNowTime		校验时间小于当前时间
*/
func (v *ValidatorX) RegisterTranslation() {
	_ = v.Validate.RegisterTranslation("urlFormat", v.trans, func(ut ut.Translator) error {
		return ut.Add("urlFormat", "{0} 不符合url或uri的格式", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("urlFormat", fe.Field())
		return t
	})

	_ = v.Validate.RegisterTranslation("passwordFormat", v.trans, func(ut ut.Translator) error {
		return ut.Add("passwordFormat", "{0} 密码必须为字母，数字，特殊字符中至少两种的组合", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwordFormat", fe.Field())
		return t
	})

	_ = v.Validate.RegisterTranslation("phoneFormat", v.trans, func(ut ut.Translator) error {
		return ut.Add("phoneFormat", "{0} 数字组成，1开头并且长度必须为11", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phoneFormat", fe.Field())
		return t
	})

	_ = v.Validate.RegisterTranslation("ltNowTime", v.trans, func(ut ut.Translator) error {
		return ut.Add("ltNowTime", "{0} 参数错误", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ltNowTime", fe.Field())
		return t
	})
}

func CheckUrl(f validator.FieldLevel) bool {
	val := f.Field().String()
	urlPartten := "^(http|https|ftp)\\://([a-zA-Z0-9\\.\\-]+(\\:[a-zA-Z0-9\\.&amp;%\\$\\-]+)*@)*((25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])|localhost|([a-zA-Z0-9\\-]+\\.)*[a-zA-Z0-9\\-]+\\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(\\:[0-9]+)*(/($|[a-zA-Z0-9\\.\\,\\?\\'\\\\\\+&amp;%\\$#\\=~_\\-]+))*$"
	if ok, _ := regexp.MatchString(urlPartten, val); ok {
		return true
	}
	return false
}

func Password(f validator.FieldLevel) bool {
	val := f.Field().String()
	if ok, _ := regexp.MatchString(`[^\x00-\xff]`, val); ok {
		return false
	}
	if ok, _ := regexp.MatchString(`\s`, val); ok {
		return false
	}

	num := 0

	if ok, _ := regexp.MatchString(`[0-9]`, val); ok { //数字命中
		num++
	}
	if ok, _ := regexp.MatchString(`[a-zA-Z]`, val); ok { //字母命中
		num++
	}

	if ok, _ := regexp.MatchString(`[^A-z0-9]`, val); ok { //特殊符号命中
		num++
	}

	if num < 2 { //命中少于两次 返回错误
		return false
	}

	return true
}

func Phone(f validator.FieldLevel) bool {
	val := f.Field().String()
	if ok, _ := regexp.MatchString(`^1[0-9]{10}$`, val); !ok {
		return false
	}
	return true
}

func LtNowTime(f validator.FieldLevel) bool {
	t := f.Field().Int()
	if int64(t) >= time.Now().Unix() {
		return false
	}
	return true
}