package auth

import (
	"net/http"
	"strings"
	"uapply/utils/jwt"
	"uapply/utils/response"

	"github.com/gin-gonic/gin"
)

const (
	UserIDKey = "user_id"
	OrgaIDKey = "orga_id"
	DepaIDKey = "depa_id"
	InteIDKey = "inte_id"
)

func JWTAuthUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		mc, err := jwt.ParseTokenUser(parts[1])
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(UserIDKey, mc.UserID)
		println("success")
		// Subsequent handlers can use c.Get(UserIDKey) gets the borrower
		c.Next()
	}
}

func JWTAuthOrga() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		mc, err := jwt.ParseTokenOrga(parts[1])
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(OrgaIDKey, mc.OrgaID)
		println("success")
		// Subsequent handlers can use c.Get(UserIDKey) gets the borrower
		c.Next()
	}
}

func JWTAuthDepa() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		mc, err := jwt.ParseTokenDepa(parts[1])
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(DepaIDKey, mc.DepaID)
		c.Set(OrgaIDKey, mc.OrgaID)
		// Subsequent handlers can use c.Get(UserIDKey) gets the borrower
		c.Next()
	}
}

func JWTAuthInte() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Split by space
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// parts[1] is the acquired tokenString,
		//and we use the previously defined function to parse JWT to parse it
		mc, err := jwt.ParseTokenInte(parts[1])
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, response.CodeTokenInvalid)
			c.Abort()
			return
		}
		// Save the currently requested message information to the requested context c
		c.Set(InteIDKey, mc.InteID)
		c.Set(DepaIDKey, mc.DepaID)
		c.Set(OrgaIDKey, mc.OrgaID)
		println("success ", mc.OrgaID, " ", mc.DepaID)
		// Subsequent handlers can use c.Get(UserIDKey) gets the borrower
		c.Next()
	}
}
