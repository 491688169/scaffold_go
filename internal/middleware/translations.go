/*
 * @Author: Kim
 * @Date: 2021-05-12 14:00:40
 * @LastEditTime: 2021-05-12 14:14:44
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/middleware/translations.go
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)

		if ok {
			switch locale {
			case "zh":
				zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
