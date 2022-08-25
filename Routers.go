package Routers

import (
	Components "enhance-be/Compoonents"
	"enhance-be/Compoonents/GoLog"
	"enhance-be/Compoonents/Host"
	"enhance-be/Compoonents/Middle"
	"enhance-be/Compoonents/Transform"
	"enhance-be/Compoonents/WebSocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/guoliben/testpublish"
	"net/http"
	"time"
)

var connectionPool = make(map[string]WebSocket.Connection)
var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func GoRouter(host string) {
	GoLog.Info(nil, "Start Router")
	testpublish.SayHi("aaafdafdafad")

	cfg := Components.InitConfigFromToml()
	artifactPath, _ := cfg.GetValue("APP_ARTIFACT", "root_path")

	router := gin.Default()

	router.LoadHTMLGlob("public/views/index.html")
	//router.LoadHTMLGlob("index.html")
	//get请求返回显示页面 index.html
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Use(Middle.Cors(), Middle.SetRequestId())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	/*
	 * 分组：静态资源访问
	 */
	//router.Static("/static", "./public")
	//router.StaticFS("/static", http.Dir("public"))
	router.Static("/favicon.ico", "./public/favicon.ico")
	router.StaticFS("/images", http.Dir("./public/images/upload"))
	router.StaticFS("/download", http.Dir(artifactPath))

	//staticer := router.Group("/public")
	//{
	//	//staticer.Static("/static", "public")
	//	staticer.StaticFS("/images", http.Dir("./public/images/"))
	//}
	//router.Static("/assets", "./assets")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	/*
	 * 分组：测试模块
	 */
	test := router.Group("/test") // Middle.AuthChecker())
	{
		test.POST("/info", func(c *gin.Context) {
			//c.JSON(http.StatusOK, Admin.Test(c))
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "ok", "result": ""})
		})
	}

	//-----开始 websocket 配置------//
	router.GET("/send", func(c *gin.Context) {
		ModifyClientInfo(c.Query("ip"))
		c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "ok", "result": ""})
	})
	router.GET("/sendclient", func(c *gin.Context) {
		sendByClientId(c.Query("ip"))
		c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "ok", "result": ""})
	})
	router.GET("/ws", HandleWebSocket)
	//-----结束 websocket 配置------//


	//-----开始 https/http 配置------//
	isHttps, _ := cfg.GetValue("SERVER", "https")
	port, _ := cfg.GetValue("SERVER", "port")
	certPath, _ := cfg.GetValue("SERVER", "cert_path")
	keyPath, _ := cfg.GetValue("SERVER", "key_path")
	portInt := Transform.StringToInt(port)

	if isHttps == "true" {
		router.Use(Middle.TlsHandler(portInt))
		err := router.RunTLS(":"+port, certPath, keyPath)
		GoLog.Info(nil, "Route run! ", fmt.Sprintf("%s", err))
	} else {
		err := router.Run(":" + port)
		GoLog.Info(nil, "Route run! ", fmt.Sprintf("%s", err))
	}
	//-----结束 https/http 配置------//

}

/**
 * 处理限定客户端发送需求  Start .........................
 * 后期修改为消息队列模式 保证不丢包
 */
func sendByClientId(clientIp string) {

	//fmt.Println("发送给执行客户端。。。。。", connectionPool[clientIp].Conn.WriteMessage())
	err := connectionPool[clientIp].Conn.WriteMessage(100, []byte("您被选中作为幸运儿 发送数据"))
	if err != nil {
		fmt.Println(err)
	}
}
func ReverseClientInfo(clientIp string) {
	fmt.Println(connectionPool[clientIp].AllowSend)

	con := &WebSocket.Connection{
		AllowSend: false,
		ClientIp:  connectionPool[clientIp].ClientIp,
		Conn:      connectionPool[clientIp].Conn,
		InCh:      connectionPool[clientIp].InCh,
		OutCh:     connectionPool[clientIp].OutCh,
		ExitCh:    connectionPool[clientIp].ExitCh,
	}
	connectionPool[clientIp] = *con
}

func ModifyClientInfo(clientIp string) {
	fmt.Println(connectionPool[clientIp].AllowSend)

	con := &WebSocket.Connection{
		AllowSend: true,
		ClientIp:  connectionPool[clientIp].ClientIp,
		Conn:      connectionPool[clientIp].Conn,
		InCh:      connectionPool[clientIp].InCh,
		OutCh:     connectionPool[clientIp].OutCh,
		ExitCh:    connectionPool[clientIp].ExitCh,
	}
	connectionPool[clientIp] = *con
}

func checkClient(conn *WebSocket.Connection) bool {

	fmt.Println("connection Pool")
	//fmt.Println(connectionPool)

	if connectionPool[conn.ClientIp].AllowSend {
		return true
	}

	//查询符合条件客户端连接
	//if conn.AllowSend  {
	//	return true
	//}
	return false
}

/**
 * 处理限定客户端发送需求  End .........................
 */

func HandleWebSocket(ctx *gin.Context) {
	conn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	wsConn, _ := WebSocket.NewConnection(conn)

	//fmt.Println("---->",wsConn.Conn)
	wsConn.Start()

	hostIp := Host.GetHostIp()

	//发送客户端心跳
	go func() {
		fmt.Println("==============发送客户端消息=============")

		for {
			time.Sleep(time.Second * 5)
			err := wsConn.WriteMessage([]byte("{'host_ip':'" + hostIp + "'}"))
			if err != nil {
				fmt.Println(err)
			}
			if checkClient(wsConn) { // 对符合条件连接客户端执行 连接操作
				ReverseClientInfo(wsConn.ClientIp)
				err = wsConn.WriteMessage([]byte("您被选中作为幸运儿 发送数据"))
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}()

	for {
		fmt.Println("==============读取客户端消息=============")
		msg, _ := wsConn.ReadMessage()
		wsConn.ClientIp = string(msg)
		fmt.Println("------从客户端读取消息", string(msg))
		connectionPool[string(msg)] = *wsConn
		//err := wsConn.WriteMessage(msg)
		//fmt.Println("------把消息返给客户端", string(msg))

		//if err != nil {
		//	fmt.Println(err)
		//}
	}
}
