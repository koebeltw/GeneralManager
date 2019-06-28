package connectHandler

import (
	"github.com/koebeltw/Common/tcp"
	"log"
)

type ConnectHandler struct {}

// OnUserConnect 客戶端連接事件
func (ser *ConnectHandler) OnUserConnect(s tcp.Session) {
	log.Printf("[%d] %s connect %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserDisconnect 客戶端斷開連接事件
func (ser *ConnectHandler) OnUserDisconnect(s tcp.Session) {
	log.Printf("[%d] %s disconnect %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnServerInit 伺服器端初始事件
func (ser *ConnectHandler) OnServerInit(s tcp.BaseServer) {
	log.Println("OnServerInit")
}

// OnServerDestroy 伺服器端結束事件
func (ser *ConnectHandler) OnServerDestroy(s tcp.BaseServer) {
	log.Println("OnServerDestroy")
}
