package tools

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type MyClaims struct {
	ID int64 `json:"username"`
	jwt.StandardClaims
}

var Secret = []byte("secret") //密码自行设定
const TokenExpireDuration = time.Hour * 24 //设置过期时间

//生成token
func GenToken(ID int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		ID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "superxon",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}


//对密码进行一个加密和解密
func HashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

//对密码进行解析
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)

    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}