package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var secretKey = []byte(viper.GetString("jwt.secret")) // 通过环境变量配置‌

// CustomClaims 自定义声明结构体
type CustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(id uint, name string) (string, error) {
	claims := CustomClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效期 24 小时
			Issuer:    "web_go",                                           // 签发者标识 例如：应用名或者域名
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err // 包含 Token 过期/格式错误等情况
}

// AuthMiddleware JWT 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/api/v1/user/login" || c.Request.URL.String() == "/api/v1/user/register" {
			c.Next()
			return
		}
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "未提供认证令牌"})
			return
		}

		claims, err := ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "无效令牌: " + err.Error()})
			return
		}

		// 将用户信息存入上下文
		c.Set("id", claims.Id)
		c.Set("name", claims.Name)
		c.Next()
	}
}
