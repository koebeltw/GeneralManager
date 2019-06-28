package create

import (
	"github.com/koebeltw/Common/tcp"
	"github.com/koebeltw/Common/type"
	"github.com/koebeltw/Common/util"
	"github.com/koebeltw/LineSlotCreator/ls/loadConfig"
	"github.com/koebeltw/GeneralManager/gm/server/connectHandler"
	"github.com/koebeltw/GeneralManager/gm/server/eventHandler"
	"github.com/koebeltw/GeneralManager/gm/sockets"
	"log"
)

func init() {
	//ok := flag.Bool("ok", false, "is ok")
	//id := flag.Int("id", 0, "id")
	//port := flag.String("port", ":8080", "http listen port")
	//var name string
	//flag.StringVar(&name, "name", "123", "name")
	//
	//flag.Parse() // 這一句至關重要！！必須先解析才能拿到參數值
	//
	//fmt.Println("ok:", *ok)
	//fmt.Println("id:", *id)
	//fmt.Println("port:", *port)
	//fmt.Println("name:", name)

	sockets.Server = tcp.CreateBaseServer()

	coder := tcp.NewMsgHead()
	serverEventHandler := tcp.GetEventHandler(eventHandler.EventHandler{}, nil)
	serverConnectHandler := &connectHandler.ConnectHandler{}
	sockets.Server.SetEventHandler(serverEventHandler)
	sockets.Server.SetCoder(coder)
	sockets.Server.SetUserHandler(serverConnectHandler)
	sockets.Server.SetServerHandler(serverConnectHandler)
	sockets.Server.SetSessionMgr(tcp.NewSessionsMgr(1))
	sockets.Server.SetServerAddr(Type.Addr{
		ID:   0,
		Name: "GM",
		IP:   util.GetIPs()[0],
		Port: int(util.ConvInt32(loadConfig.Config.LSPort)),
	})

	sockets.Server.SetClientsAddr([]Type.Addr{{
		ID:   0,
		Name: "GM",
		IP:   util.GetIPs()[0],
		Port: int(util.ConvInt32(loadConfig.Config.LSPort)),
	}})

	go sockets.Server.Start()

	log.Println("Socket start at " + util.GetIPs()[0] + ":" +  loadConfig.Config.LSPort)
}
