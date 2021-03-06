package utils

import (
	//"fmt"
	"math/rand"
	"errors"
)

type Licence struct {
	letterRunes []rune
	length      int
	attempts    int
	quantity    int

	licenseMap   map[string]string
	licenseSlice []string
}
func (l *Licence) Generate(quantity int64) ([]string, error) {
	l.quantity = int(quantity)
	l.licenseMap = make(map[string]string, quantity)
	l.licenseSlice = make([]string, 0, quantity)

	for i := 0; i < int(quantity); i++ { //根据指定数量生成
		if licenseStr, err := l.getLicense(1); err != nil { //每生成一个code都是第1次尝试
			return l.licenseSlice, err
		} else {
			l.licenseMap[licenseStr] = licenseStr
			l.licenseSlice = append(l.licenseSlice, licenseStr)
		}
	}
	return l.licenseSlice, nil

}

func (l *Licence) getLicense(attempts int) (string, error) {
	bytes := make([]rune, l.length)
	for i := 0; i < l.length; i++ {
		bytes[i] = l.letterRunes[rand.Intn(len(l.letterRunes))]
	}
	licenseStr := string(bytes)
	if _, ok := l.licenseMap[licenseStr]; ok { //code存在 重新生成
		if attempts <= l.attempts { //尝试次数大于指定的尝试次数之后返回错误
			attempts = attempts + 1
			return l.getLicense(attempts)
		} else {
			return licenseStr, errors.New("尝试失败")
		}
	}
	return licenseStr, nil
}

func NewLicence(length, attempts int) *Licence {
	return &Licence{
		letterRunes: []rune("0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuioplkjhgfdsazxcvbnm"),
		length:      length,
		attempts:    attempts,
	}
}

/**
l := NewLicence(8, 5)
licenseSlice, err := l.Generate(2)
if err != nil {
	fmt.Println(err)
	return 
}
fmt.Println(licenseSlice, err)
**/
