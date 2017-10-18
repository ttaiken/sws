package main
import (
	"log"
	"net/http"
	"strconv"
	"sws/conf"
)

func main(){
	cfg := conf.ParseConfig()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./webroot"))))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil))


}

