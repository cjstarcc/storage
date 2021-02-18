package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"storage/conf"
	"strings"
)

func Handler(w http.ResponseWriter,r *http.Request){
	m :=r.Method
	if m == http.MethodPut{
		Put(w,r)
		return
	}
	if m == http.MethodGet {
		Get(w,r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Put(w http.ResponseWriter,r *http.Request)  {
	f,err := os.Create(conf.GetConfig().Env.Dir+"/objects/"+strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f,r.Body)
}

func Get(w http.ResponseWriter,r *http.Request)  {
	f,err:=os.Open(conf.GetConfig().Env.Dir+"/objects/"+strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w,f)
}