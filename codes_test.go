package status

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	s := Init()
	fmt.Println(s)
}

func TestInitWith(t *testing.T) {
	s := InitWith(ZH, map[string]string{
		"key": "值",
	})
	fmt.Println(s)
	fmt.Println(s.CurrentLanguage())
	s.SetLanguage(EN)
	fmt.Println(s.CurrentLanguage())
	fmt.Println(s.T("key"))
	err := s.New(Server, "test error message")
	fmt.Println(err)
	err = s.New(Cache, s.T("key"))
	fmt.Println(err)
	fmt.Println(s.Code())
	fmt.Println(s.Msg())
	err = s.Newf(Gateway, "%s--%d", "string", 123)
	fmt.Println(err)
	err = s.FromError(Redirect, err)
	fmt.Println(err)

}
