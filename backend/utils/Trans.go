package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

func GetTrans() ut.Translator {

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")
	return trans
	
}
