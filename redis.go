package gin_session_persistence

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSessionStore(route *gin.Engine, size int, network string, address string, password string, secret string, sessionname string)  {


	store, err := redis.NewStore(size,network,address,password, []byte(secret))
	if err != nil{
		fmt.Println(err)
	}
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 3600, Secure: true, Path: "/"})

	route.Use(sessions.Sessions(sessionname, store))

}
