package controllers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"net/http"
	"github.com/guowenshuai/apiproject/models"
)

// WebSocket controller
type WebSocketController struct {
	beego.Controller
}

// http upGrade to websocket connecting
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 不检查请求源
		return true
	},
}


var (
	cMsgIn      = make(chan models.Event, 10)
	cRegister   = make(chan *connHolder, 10)
	cUnRegister = make(chan string, 10)

	connPools = list.New()
)

type connHolder struct {
	uid  string
	Conn *websocket.Conn
}

func init() {
	go daemonService()
}

// Get method handles GET request for webSocket router
func (this *WebSocketController) Get() {
	//beego.Info("websocket index page")
	//this.TplName = "webSocket.html"
}


func (this *WebSocketController) Upgrade() {
	uid := this.GetString(":uid")
	beego.Info("new websocket user: ", uid)
	if len(uid) == 0 {
		this.Redirect("/", 302)
		return
	}

	conn, err := upGrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		logs.Error("websocket upgrade error: ", err.Error())
		this.Redirect("/", 302)
		return
	}
	cRegister <- &connHolder{uid, conn}
	defer func() {
		cUnRegister <- uid
	}()

	for {
		msgType, p, err := conn.ReadMessage()
		logs.Info("msgType: ", msgType)
		if err != nil {
			logs.Error("webSocket uid %s readmessag error: %s", err.Error())
			return
		}
		cMsgIn <- newMessage(uid, string(p))
	}

}

// broadCastWebSocket method broadCast message to all client
func broadCastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("json marshal error: ", err.Error())
		return
	}
	for e := connPools.Front(); e != nil; e = e.Next() {
		conn := e.Value.(*connHolder).Conn
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			cUnRegister <- e.Value.(*connHolder).uid
		}
	}
}

// newMessage return a models.Event message
func newMessage(uid string, msg string) models.Event {
	return models.Event{models.EVENT_MESSAGE, uid,
		int(time.Now().Unix()), msg}
}


func daemonService() {
	for {
		select {
		// 注册客户端保持
		case connH := <-cRegister:
			if !isUidExist(connH.uid) {
				connPools.PushBack(connH)
				cMsgIn <- models.Event{models.EVENT_JOIN, connH.uid, int(time.Now().Unix()), ""}
				beego.Info("New User: ", connH.uid, " join")
			} else {
				beego.Info("Old User: ", connH.uid)
			}
			// 清除客户端保持
		case uid := <-cUnRegister:
			for e := connPools.Front(); e != nil; e = e.Next() {
				if e.Value.(*connHolder).uid == uid {
					connPools.Remove(e)
					if conn := e.Value.(*connHolder).Conn; conn != nil {
						conn.Close()
						beego.Error("User ", uid, " close")
					}
					cMsgIn <- models.Event{models.EVENT_LEAVE, uid, int(time.Now().Unix()), ""}
					break
				}
			}
		case event := <-cMsgIn:
			broadCastWebSocket(event)
			models.NewEvent(event)
		}
	}
}

func isUidExist(uid string) bool {
	for e := connPools.Front(); e != nil; e = e.Next() {
		if e.Value.(*connHolder).uid == uid {
			return true
		}
	}
	return false
}

