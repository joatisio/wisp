package router

import "github.com/gin-gonic/gin"

type Pagination struct {
	Total       uint64 `json:"total"`
	PerPage     uint64 `json:"per_page"`
	CurrentPage uint64 `json:"current_page"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type NormalResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Errors map[string]interface{}
}

func GinJsonError(c *gin.Context, errResponse interface{}, statusCode int) {
	c.AbortWithStatusJSON(statusCode, errResponse)
}

func GinJsonResponse(c *gin.Context, resp interface{}, statusCode int) {
	c.JSON(statusCode, resp)
}
