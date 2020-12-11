package app

import (
	"time"

	"github.com/go-programming-tour-book/blog-service/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-programming-tour-book/blog-service/global"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

/*
GenerateToken方法的主要功能是生成JWT Token，其流程是根据客户端传入的AppKey和AppSecret，
以及在项目配置中设置的签发者（Issuer）和过期时间（ExpiresAt），根据指定的算法生成签名后的Token

· jwt.NewWithClaims：根据Claims结构体创建Token实例。
	它一共包含两个形参，第一个形参是 SigningMethod，其包含 SigningMethodHS256、SigningMethodHS384 和SigningMethodHS512三种crypto.Hash加密算法的方案。
	第二个形参是Claims，主要用于传递用户预定义的一些权利要求，以便后续的加密、校验等行为。

· tokenClaims.SignedString：生成签名字符串，根据传入的 Secret，进行签名并返回标准的Token。
 */
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	// 第一块是嵌入的AppKey和AppSecret，用于我们自定义的认证信息
	// 第二块是jwt.StandardClaims结构体，它是在jwt-go库中预定义的, 对应Payload的相关字段，这些字段都是非强制性的，但官方建议使用预定义权利要求，能够提供一组有用的、可相互操作的约定
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

/*
ParseToken方法的主要的功能是解析和校验Token，其流程是解析传入的Token，然后根据Claims的相关属性要求进行校验

· ParseWithClaims：用于解析鉴权的声明，方法内部是具体的解码和校验的过程，最终返回*Token。

· Valid：验证基于时间的声明，如过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before）。需要注意的是，即便在令牌中没有任何声明，也仍然被认为是有效的。
 */
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
