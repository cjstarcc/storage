package main

import (
	"log"
	"net/http"
	"storage/conf"
	"storage/objects"
	"storage/util/heartBeat"
	"strconv"
)

func main()  {
	conf.InitConfig()
	go heartBeat.StartHeartBeat()
	go heartBeat.StartLocate()
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(conf.GetConfig().Env.Port),nil))
}
