package handler

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type Payload struct {
	Aud         []string
	Azp         string
	Iat         int
	Iss         string
	Permissions []string
	Scope       string
	Sub         string
}

func Auth(c *gin.Context, permissions []string) {
	rawToken := c.GetHeader("Authorization")
	data, err := decode(rawToken)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(403, gin.H{"message": "Unauthorized"})
		return
	}
	for _, s := range permissions {
		if slices.Contains(data.Permissions, s) {
			c.Next()
			return
		}
	}
	c.AbortWithStatusJSON(403, gin.H{"message": "Unauthorized"})
}

func decode(token string) (Payload, error) {
	var data Payload
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return data, &JwtError{Message: "Invalid Token"}
	}
	bytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return data, err
	}
	err = json.Unmarshal([]byte(bytes), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

type JwtError struct {
	Message string
}

func (e *JwtError) Error() string {
	return e.Message
}
