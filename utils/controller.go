package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)


func GetInt64(ctx *gin.Context, col string) int64 {
	r := ""
	if ctx.Request.Method == "GET" {
		r = ctx.Query(col)
	} else if ctx.Request.Method == "POST" {
		r = ctx.PostForm(col)
	}
	i, _ := strconv.ParseInt(r, 10, 64)
	return i
}

func GetFloat64(ctx *gin.Context, col string) float64 {
	r := ""
	if ctx.Request.Method == "GET" {
		r = ctx.Query(col)
	} else if ctx.Request.Method == "POST" {
		r = ctx.PostForm(col)
	}
	i, _ := strconv.ParseFloat(r, 64)
	return i
}

func GetBool(ctx *gin.Context, col string) bool {
	r := ""
	if ctx.Request.Method == "GET" {
		r = ctx.Query(col)
	} else if ctx.Request.Method == "POST" {
		r = ctx.PostForm(col)
	}
	i, _ := strconv.ParseBool(r)
	return i
}

func GetInt(ctx *gin.Context, col string) int {
	return int(GetInt64(ctx, col))
}

func GetLimit(ctx *gin.Context) int {
	limit := GetInt(ctx, "limit")

	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return limit
}


func GetPageSize(ctx *gin.Context) int {
	limit := GetInt(ctx, "pageSize")

	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return limit
}

func GetPage(ctx *gin.Context) int {
	page := GetInt(ctx, "page")

	if page == 0 {
		page = 1
	}

	return page
}

