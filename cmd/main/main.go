package main
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/anas-infobell/book/pkg/routes"
)
func main(){
r := mux.NewRouter()
routes.registerBookStoreRouters(r)
http.Handle("/",r)
log.Fatal(http.listenAndServe("localhost:9010",r))
}