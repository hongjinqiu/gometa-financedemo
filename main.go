package main

import (
	"flag"
	"fmt"
	"github.com/hongjinqiu/gometa/app"
	"github.com/hongjinqiu/gometa/console"
	"github.com/hongjinqiu/gometa/log"
	"github.com/hongjinqiu/gometa/route"
	"github.com/hongjinqiu/gometa/session"
	"html"
	"net/http"
	"reflect"
	"strings"
)

var addr = flag.String("addr", ":8888", "address")

var constrollersDict []reflect.Type = []reflect.Type{}

func init() {
	constrollersDict = append(constrollersDict, reflect.TypeOf(console.Console{}))
	constrollersDict = append(constrollersDict, reflect.TypeOf(app.App{}))

	route.FilterLi = append(route.FilterLi, SessionAdapter)
	route.FilterLi = append(route.FilterLi, UTF8Filter)
}

func SessionAdapter(w http.ResponseWriter, r *http.Request, li []route.HttpHandleFilterFunc) {
	userAgent := strings.Join(r.Header["User-Agent"], ",")
	if strings.Index(userAgent, "Firefox") > -1 {
		session.SetToSession(w, r, "userId", "10")
	} else {
		session.SetToSession(w, r, "userId", "20")
	}
	session.SetToSession(w, r, "adminUserId", "1")
	li[0](w, r, li[1:])
}

func UTF8Filter(w http.ResponseWriter, r *http.Request, li []route.HttpHandleFilterFunc) {
	w.Header()["Content-Type"] = []string{"text/html; charset=utf-8"}
	li[0](w, r, li[1:])
}

func registerRoute() {
	route.RegisteStaticFilePath()
	route.RegisteReflectController(constrollersDict)

	http.HandleFunc("/bar/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	if true {
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(constrollersDict[0].Name())
		for i := 0; i < constrollersDict[0].NumMethod(); i++ {
			fmt.Println(constrollersDict[0].Method(i).Name)
		}
		//		inst := reflect.New(constrollersDict[0]).Elem().Interface()
		//		reflectValue := reflect.ValueOf(inst)
		//		fmt.Print(reflectValue)
		//		for i := 0; i < reflectValue.NumMethod(); i++ {
		//			fmt.Print(reflectValue.Method(i).Name)
		//		}
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "Hello World")
		} else {
			pathLi := strings.Split(r.URL.Path, "/")
			fmt.Println("gogogog")
			fmt.Println(pathLi)
			tmpPathLi := []string{}
			for _, item := range pathLi {
				if item != "" {
					tmpPathLi = append(tmpPathLi, item)
				}
			}
			if len(tmpPathLi) > 1 {
				fmt.Println("^^^^^^^*******))))))))")
				fmt.Println(tmpPathLi)
			}
			fmt.Fprintf(w, "Hello 2/, %q", html.EscapeString(r.URL.Path))
		}

	})
}

//devsrvr -addr=:8001 -targetaddr=localhost:8888 D:\goworkspace\src\github.com\gometa-financedemo D:\goworkspace\src\github.com\gometa
//devsrvr -addr=:8001 -targetaddr=localhost:8888 /home/hongjinqiu/goworkspace/src/github.com/hongjinqiu/gometa-financedemo /home/hongjinqiu/goworkspace/src/github.com/hongjinqiu/gometa
func main() {
	flag.Parse()
	log.Warn("listening on:", *addr)

	registerRoute()

	log.Error(http.ListenAndServe(*addr, nil))
}
