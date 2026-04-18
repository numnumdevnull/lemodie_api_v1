package router

import (
	"database/sql"

	"lemodie_api_v1/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(db *sql.DB) *gin.Engine {
	r := gin.Default()

	dictTypeHandler := &handler.DictionaryTypeHandler{DB: db}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.GET("/dictionary-types", dictTypeHandler.GetAll)
	}

	return r
}
