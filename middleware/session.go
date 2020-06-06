package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Session(key string) gin.HandlerFunc{
	store, _ := redis.NewStore(10, "tcp", "47.107.224.190:6379", "", []byte(key))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 60*30, Path: "/"})
	return sessions.Sessions("User-session",store)
}