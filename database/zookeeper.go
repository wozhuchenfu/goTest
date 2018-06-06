package database

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"encoding/asn1"
	"encoding/json"
)

type ServiceNode struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int `json:"port"`
}

type SdClient struct {
	zkServers []string
	zkRoot string
	conn *zk.Conn
}

func NewClient(zkServers []string,zkRoot string,timeout int ) (*SdClient,error) {
	client := new(SdClient)
	client.zkServers = zkServers
	client.zkRoot = zkRoot
	conn,_,err := zk.Connect(zkServers,time.Duration(timeout)*time.Second)
	if err!=nil {
		fmt.Println(err)
		return nil,err
	}
	client.conn = conn

	if err := client.ensureRoot();err != nil {
		client.conn.Close()
		return nil,err
	}
	return client,nil

}

func (s *SdClient) close() {
	s.conn.Close()
}

func (s *SdClient) ensureRoot() error {
	exit,_,err := s.conn.Exists(s.zkRoot)
	if err != nil {
		return err
	}
	if !exit {
		_,err := s.conn.Create(s.zkRoot,[]byte(""),0,zk.WorldACL(zk.PermAll))
		if err != nil && err!=zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

func (s *SdClient) ensureName(name string) error{
	path := s.zkRoot+"/"+name
	exist,_,err := s.conn.Exists(path)
	if err != nil {
		return err
	}
	if !exist {
		_,err := s.conn.Create(path,[]byte(""),0,zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}
	}
	return nil

}
//服务注册方法
func (s *SdClient) Register(node *ServiceNode) error{
	if err := s.ensureName(node.Name); err!=nil {
		return err
	}
	path := s.zkRoot+"/"+node.Name+"/n"
	data,err := asn1.Marshal(node)
	if err!=nil {
		return err
	}
	_,err = s.conn.CreateProtectedEphemeralSequential(path,data,zk.WorldACL(zk.PermAll))
	if err!=nil {
		return err
	}
	return nil
}

//消费者获取服务列表方法
func (s *SdClient) GetNodes(name string) ([]*ServiceNode,error){

	path := s.zkRoot+"/"+name;
	childs,_,err := s.conn.Children(path)
	if err!=nil {
		if err==zk.ErrNoNode {
			return []*ServiceNode{},nil
		}
		return nil,err
	}
	nodes := []*ServiceNode{}
	for _,child := range childs  {
		fullPath := path + "/" + child
		data,_, err := s.conn.Get(fullPath)
		if err!=nil {
			if err == zk.ErrNoNode {
				continue
			}
			return nil,err
		}
		node := new(ServiceNode)
		err= json.Unmarshal(data,node)
		if err!=nil {
			return nil,err
		}
		nodes = append(nodes,node)
	}
	return nodes,nil
}



















