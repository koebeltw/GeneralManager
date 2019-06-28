package connectHandler

import (
	"github.com/koebeltw/Common/tcp"
	"log"
)

type ConnectHandler struct{}

// OnUserConnect 客户端連接事件
func (c *ConnectHandler) OnUserConnect(s tcp.Session) {
	log.Printf("[%d] %s connect %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserDisconnect 客户端斷開連接事件
func (c *ConnectHandler) OnUserDisconnect(s tcp.Session) {
	log.Printf("[%d] %s disconnect %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}

// OnUserReConnect blabla
func (c *ConnectHandler) OnUserReConnect(s tcp.Session) {
	log.Printf("[%d] %s reconnect %s", s.GetID(), s.LocalAddr(), s.RemoteAddr())
}
