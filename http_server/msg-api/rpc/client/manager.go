package client

import (
	"github.com/golang/glog"
	"github.com/oikomi/FishChatServer2/http_server/msg-api/conf"
	"github.com/oikomi/FishChatServer2/protocol/rpc"
	sd "github.com/oikomi/FishChatServer2/service_discovery/etcd"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ManagerRPCCli struct {
	conn *grpc.ClientConn
}

func NewManagerRPCCli() (managerRPCCli *ManagerRPCCli, err error) {
	r := sd.NewResolver(conf.Conf.RPCClient.ManagerClient.ServiceName)
	b := grpc.RoundRobin(r)
	conn, err := grpc.Dial(conf.Conf.RPCClient.ManagerClient.EtcdAddr, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	managerRPCCli = &ManagerRPCCli{
		conn: conn,
	}
	return
}

func (managerRPCCli *ManagerRPCCli) Auth(offlineMsgsReq *rpc.MGOfflineMsgReq) (res *rpc.MGOfflineMsgRes, err error) {
	m := rpc.NewManagerServerRPCClient(managerRPCCli.conn)
	if res, err = m.GetOfflineMsgs(context.Background(), offlineMsgsReq); err != nil {
		glog.Error(err)
	}
	return
}
