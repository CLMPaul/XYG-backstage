package middleware

import (
	"errors"
	"strings"
	"time"
	"xueyigou_demo/global"
	"xueyigou_demo/pkg/e"
	"xueyigou_demo/serializer"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	Id        int64  `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Privilege int    `json:"privilege"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	// MyJWTKey 自定义的token秘钥
	MyJWTKey = []byte("SuperKey")
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
	//换票区间
	bufferTime = int64(60 * 24 * time.Minute)

	RefreshKey = []byte("xueyigourefreshtoken")
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) (string, string) {
	var err error
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	claims.Issuer = "douyin"

	//生成刷新token
	refresh_claim := &UserClaims{
		Id: claims.Id,
		//Name: claims.Name,
		Phone: claims.Phone,
	}
	refresh_claim.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).Unix()
	refresh_claim.Issuer = "xueyigou"

	// 生成token，并签名生成JWT
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MyJWTKey)

	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claim).SignedString(RefreshKey)
	if err != nil {
		panic(err)
	}
	return token, refresh_token
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*UserClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyJWTKey, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok {
			return claims, nil
		}
	}
	return nil, err
}

// ParseToken 解析Token
func ParseRefreshToken(tokenString string) (*UserClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return RefreshKey, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok {
			return claims, nil
		}
	}
	return nil, err
}

// Refresh 刷新token
func Refresh(token string) (interface{}, error) {
	//token := strings.Replace(c.GetHeader("Authorization"), "Bearer", "", -1)
	code := 200
	if token == "" {
		code = 404
	}
	claims, err := ParseRefreshToken(token)
	if err != nil {
		code = e.ErrorAuthCheckTokenFail
	} else if time.Now().Unix()+global.RTokenExpires.Unix() > claims.ExpiresAt {
		code = e.ErrorAuthCheckRTokenTimeout
	}
	if code != 200 {
		return gin.H{
			"result_status": code,
			"result_msg":    e.GetMsg(code),
		}, errors.New("err")
	}
	a_token, _ := GenerateToken(claims)
	Token := serializer.Token{
		AccessToken:  a_token,
		RefreshToken: token,
	}
	return gin.H{
		"result_status": 0,
		"result_msg":    "refresh",
		"Token":         Token,
	}, nil
}

// JWT token验证中间件
// param level 0:普通 1:管理员 2：超管
func JWT(level int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer", "", -1)
		if token == "" {
			code = 401
		} else {
			claims, err := ParseToken(token)
			if claims.Privilege < level {
				code = e.ErrorPrivilege
			} else {
				c.Set("userClaim", claims)
				if err != nil {
					code = e.ErrorAuthCheckTokenFail
				} else if time.Now().Unix()+global.Expires.Unix() > claims.ExpiresAt { //测试环境弄久一点
					code = e.ErrorAuthCheckTokenTimeout
				}
			}

		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"result_status": code,
				"result_msg":    e.GetMsg(code),
				"data":          data,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
