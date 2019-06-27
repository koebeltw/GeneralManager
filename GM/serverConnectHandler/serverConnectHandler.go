package serverConnectHandler

import (
	"github.com/koebeltw/Common/session"
	"log"
)

func NewServerConnectHandler(BaseServer session.BaseServer) *ServerConnectHandler {
	return &ServerConnectHandler{
		Server: BaseServer,
	}
}

type ServerConnectHandler struct {
	Server session.BaseServer
	Client session.BaseClient
}

// OnUserConnect 客戶端連接事件
func (ser *ServerConnectHandler) OnUserConnect(s session.Session) {
	log.Printf("LS [%d] %s connect GS %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserDisconnect 客戶端斷開連接事件
func (ser *ServerConnectHandler) OnUserDisconnect(s session.Session) {
	log.Printf("LS [%d] %s disconnect GS %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnServerInit 伺服器端初始事件
func (ser *ServerConnectHandler) OnServerInit(s session.BaseServer) {
	log.Println("OnServerInit")
}

// OnServerDestroy 伺服器端結束事件
func (ser *ServerConnectHandler) OnServerDestroy(s session.BaseServer) {
	log.Println("OnServerDestroy")
}
