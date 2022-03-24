package utils

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Env struct {
	Db *sql.DB
}

func NewEnv() *Env {
	return &Env{}
}

func (e Env) SetEnv(c *gin.Context) {
	c.Set("globalVariables", e)
	c.Next()
}

func GetEnv(c *gin.Context) Env {
	a, _ := c.Get("globalVariables")
	env := a.(Env)
	return env
}
