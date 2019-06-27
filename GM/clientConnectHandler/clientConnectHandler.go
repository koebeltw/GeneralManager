package clientConnectHandler

import (
	"github.com/koebeltw/Common/session"
	"log"
)

type ClientConnectHandler struct{
	Client session.BaseClient
}

func NewClientConnectHandler(BaseClient session.BaseClient) *ClientConnectHandler {
	return &ClientConnectHandler{
		Client: BaseClient,
	}
}

// OnUserConnect 客户端連接事件
func (c *ClientConnectHandler) OnUserConnect(s session.Session) {
	log.Printf("GS [%d] %s connect DB %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserDisconnect 客户端斷開連接事件
func (c *ClientConnectHandler) OnUserDisconnect(s session.Session) {
	log.Printf("GS [%d] %s disconnect DB %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserReConnect blabla
func (c *ClientConnectHandler) OnUserReConnect(s session.Session) {
	log.Printf("GS [%d] %s reconnect DB %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}
