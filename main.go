package main

import (
	"GiliVideo/router"
)

func main(){
	r:=router.NewRouter()
	r.Run(":3000")
}

