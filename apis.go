/**
 * Copyright (c) yangzhao 2022/3/25
 *
 * File name: apis.go
 * Created at: 2022/3/25 4:33 PM
 * Author: yangzhao yzha5@163.com
 *
 * Description:
 */

package status

import (
	"fmt"
	"regexp"

	"errors"
)

// Init default init Status, set lang to EN
func Init() *Status {
	return &Status{
		currentLang: EN,
	}
}

// InitWith init Status with lang and messages
func InitWith(lang Lang, messages map[string]string) *Status {
	return &Status{
		currentLang: lang,
		messages:    messages,
	}
}

// SetLanguage set a language for current
func (s *Status) SetLanguage(l Lang) {
	s.currentLang = l
}

// SetMessage set message for current
func (s *Status) SetMessage(message map[string]string) {
	s.messages = message
}

// CurrentLanguage get current language
func (s *Status) CurrentLanguage() Lang {
	return s.currentLang
}

// T translate message
func (s *Status) T(key string) string {
	return s.messages[key]
}

func (s *Status) Code() Code {
	return s.code
}

func (s *Status) Msg() string {
	return s.msg
}

func (s *Status) Error() string {
	return fmt.Sprintf("code %d message: %s",
		s.code,
		s.msg,
	)
}
func (s *Status) New(code Code, text string) error {
	s.code = code
	s.msg = text
	return s
}
func (s *Status) Newf(code Code, format string, a ...interface{}) error {
	s.code = code
	s.msg = fmt.Sprintf(format, a...)
	return s
}

func (s *Status) FromError(code Code, err error) error {
	compile := regexp.MustCompile(`code \d+ message: `)
	s.code = code
	s.msg = compile.ReplaceAllString(err.Error(), "")
	return s
}

func Error(code Code,msg string) error {
	return errors.New(fmt.Sprintf("code %d message: %s",code,msg))
}
func Errorf(code Code,format string,arg ...interface{})error{
	return Error(code,fmt.Sprintf(format,arg...))
}