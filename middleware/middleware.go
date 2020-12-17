package middleware

import (
	"flag"
	"fmt"
	"go-clean/controllers"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"go-clean/config"

	socketio "github.com/bimagusta/go-socket.io"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Socket struct {
	Server *socketio.Server
	Err    error
}

var sock = Socket{}
var server = sock.Server
var err = sock.Err

func Middleware() {
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	dir := config.GetDir()
	//result := strings.Replace(dir, "project", "log", -1)
	// Logging to a file.

	// gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	router := gin.Default()
	router.Use(gin.Recovery())
	// f, _ := os.Create(result + "/eventRequest.go")
	// gin.DefaultWriter = io.MultiWriter(f)

	// g, _ := os.Create(result + "/eventError.go")
	// gin.DefaultErrorWriter = io.MultiWritdirer(g)

	// log.SetOutput(gin.DefaultWriter)

	router.Static("/public", dir+"/vendor/assets/")
	router.Static("/css", "../")

	if port == "" {
		port = "8000"
	}

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type,Access-Control-Allow-Origin",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	api := router.Group("/api")
	// trans.Use(gin.BasicAuth(gin.Accounts{"code": "cod3"}))
	{
		Machine := new(controllers.ControlMachine)

		//get
		api.GET("/ceksn", Machine.GetDeviceCmd)
		// api.GET("/cdata", Machine.Configuration)
		// create
		api.POST("/postdata/:sn", Machine.PostData)
		// update
		// delete
		//lain lain
	}
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"messsage": "ADMS ON"})
	})
	router.NoRoute(func(c *gin.Context) {
		// c.String(http.StatusNotFound, "Page Not Found")
		req := c.Request
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", b)
	})
	allrouter := router.Routes()
	for _, route := range allrouter {
		generate(route.Path, route.Method, route.Handler)
	}
	router.Run(":9001")
}

//MENCARI
func getvalue(data, start, end string) (value string) {
	pecah1 := strings.Split(data, start)
	pecah2 := strings.Split(pecah1[1], end)
	return pecah2[0]
}

func generate(url, method, handler string) {
	if strings.Contains(handler, "controllers") {
		findcontroll := getvalue(handler, "controllers.(*", ").")
		findfunc := getvalue(handler, ").", "-fm")
		fmt.Println("===================================")
		fmt.Println("URL : ", url)
		fmt.Println("METHOD : ", method)
		fmt.Println("HANDLER : ", handler)
		fmt.Println("CONTROLLER : ", findcontroll)
		fmt.Println("FUNCTION : ", findfunc)
		dir := config.GetDir()
		data, _ := ioutil.ReadFile(dir + "/vendor/controllers/" + findcontroll + ".go")
		// fmt.Println(err)
		parsingtahap1(findfunc, string(data), method, findcontroll)
	}
}

//LIST FUNGSI
func parsingtahap1(prefix, data, method, findcontroll string) {
	listfunc := strings.Split(string(data), "func")
	for i, func_row := range listfunc {
		if i > 0 {
			// fmt.Println(i, "->", func_row)
			status := parsingtahap2(prefix, func_row, method, findcontroll)
			if status == true {
				break
			}
		}
	}
}

//MENCARI FUNGSI
func parsingtahap2(prefix, data, method, findcontroll string) bool {
	if strings.Contains(data, prefix) {
		// fmt.Println(data) //FUNGSI DI TEMUKAN
		query := parsingquery(data)
		param := parsingparam(data)
		fmt.Println("QUERY :", query)
		fmt.Println("PARAM :", param)
		if method != "GET" {
			body := parsingbody(data, findcontroll)
			fmt.Println("BODY :", body)
		}
		return true
	} else {
		return false
	}
}

//LIST QUERY
func parsingquery(data string) (query []string) {
	listquery := strings.Split(string(data), "Query(\"")
	for i, query_row := range listquery {
		if i > 0 {
			getvalue := strings.Split(query_row, "\")")
			// fmt.Println(i, "[QUERY]->", getvalue[0])
			query = append(query, getvalue[0])
		}
	}
	return
}

//LIST PARAM
func parsingparam(data string) (param []string) {
	listparam := strings.Split(string(data), "Param(\"")
	for i, param_row := range listparam {
		if i > 0 {
			getvalue := strings.Split(param_row, "\")")
			// fmt.Println(i, "[PARAM]->", getvalue[0])
			param = append(param, getvalue[0])
		}
	}
	return
}

//MENCARI BODY IF PUT,POST
func parsingbody(data, namecontrol string) (body string) {
	listbody := strings.Split(data, "var data request.")
	getvalue := strings.Split(listbody[1], "\n")
	namestructbody := getvalue[0]
	dir := config.GetDir()
	parsingstruct(dir+"/vendor/request/"+namecontrol+".go", namestructbody)
	return
}

//FIND STRUCT
func parsingstruct(path, namestructbody string) {
	data, _ := ioutil.ReadFile(path)
	listvartext := strings.Split(string(data), "type "+namestructbody+" struct {")
	getvalue := strings.Split(listvartext[1], "}")
	// fmt.Println(getvalue[0])
	listall := strings.Split(getvalue[0], "\n")
	for _, a := range listall {
		fmt.Println(a)
	}
}
