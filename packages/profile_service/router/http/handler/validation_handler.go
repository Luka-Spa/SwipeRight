package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func validate(c *gin.Context, body interface{}) []model.ErrorMsg {
	if err:=c.ShouldBindJSON(&body);err!=nil{
		var ve validator.ValidationErrors
		fmt.Println(err)
		if errors.As(err, &ve) {
				out := make([]model.ErrorMsg, len(ve))
				for i, fe := range ve {
						out[i] = model.ErrorMsg{Field: fe.Field(), Message: getErrorMsg(fe)}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
				return out
		}
 }
 return nil
}

