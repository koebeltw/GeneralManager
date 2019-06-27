package clientEventHandler

import (
	"github.com/koebeltw/Common/session"
)

// dbMgrEventHandler blabla
type clientEventHandler struct {
	Server session.BaseServer
	Client session.BaseClient
}

func NewClientEventHandler(BaseServer session.BaseServer, BaseClient session.BaseClient) clientEventHandler {
	return clientEventHandler{
		Server: BaseServer,
		Client: BaseClient,
	}
}
