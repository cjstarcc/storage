package heartBeat

import (
	"os"
	"storage/conf"
	"storage/rabbitmq"
	"strconv"
	"time"
)

func StartHeartBeat()  {
 q:=rabbitmq.New(conf.GetConfig().Env.Rabbitmq)
 defer q.Close()
	for  {
		q.Publish("apiServer",conf.GetConfig().Env.Port)
		time.Sleep(time.Second*5)
	}
}

func StartLocate()  {
	q:=rabbitmq.New(conf.GetConfig().Env.Rabbitmq)
	defer q.Close()
	q.Bind("dataServers")
	c:=q.Consume()
	for msg :=range c{
		object,err := strconv.Unquote(string(msg.Body))
		if err!=nil{
			panic(err)
		}
		if Locate(conf.GetConfig().Env.Dir+"/objects"+object){
			q.Send(msg.ReplyTo,conf.GetConfig().Env.Port)
		}
	}
}

func Locate(name string) bool{
	_,err := os.Stat(name)
	return !os.IsNotExist(err)
}