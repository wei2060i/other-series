package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//指定加密密钥
var jwtSecret = []byte("secret")
var iss = "gin-blog"

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type claims struct {
	Username uint
	Password string
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func GenerateToken(username uint, password string, expireTime time.Time) (string, error) {
	claims := claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: iss,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func parseToken(token string) (*claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func CheckToken(token string) (uint, string, error) {
	parseToken, err := parseToken(token)
	if err != nil {
		return 0, "", errors.New("token 无法解析:" + err.Error())
	}
	if parseToken.Issuer != iss {
		return 0, "", errors.New("token iss 错误")
	}
	unix := time.Time{}.Unix()
	expiresAt := parseToken.ExpiresAt
	if expiresAt != unix && time.Unix(expiresAt, 0).After(time.Now()) {
		return parseToken.Username, parseToken.Password, errors.New("token 超时")
	}
	return parseToken.Username, parseToken.Password, nil
}