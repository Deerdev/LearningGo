package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
· go-playground/locales：多语言包，从CLDR项目（Unicode通用语言环境数据存储库）生成的一组多语言环境，主要在 i18n 软件包中使用。该库需要与 universal-translator配套使用。

· go-playground/universal-translator：通用翻译器，是一个使用CLDR数据+复数规则的Go语言i18n转换器。

· go-playground/validator/v10/translations：validator的翻译器。
 */
// 国际化：对validator的语言包翻译的相关功能
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}

		c.Next()
	}
}
/*
在识别当前请求的语言类别时，我们通过GetHeader方法获取约定的header 参数locale，判别当前请求的语言类别是 en 还是 zh。
如果有其他语言环境要求，也可以继续引入其他语言类别，因为go-playground/locales支持几乎所有语言类别。

在后续的注册步骤中，我们调用 RegisterDefaultTranslations 方法，将验证器和对应语言类型的Translator注册进来，实现验证器的多语言支持。
同时将Translator存储到全局上下文中，以便在后续翻译时使用

中间件在 internal/routers/router.go 中注册
 */
