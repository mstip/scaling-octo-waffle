package servercontrol

import (
	"context"
	"fmt"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"log"
)

type ServerInfo struct {
	ID             int
	Name           string
	IP             string
	ServerTypeID   int
	ServerTypeName string
}

const (
	CX11ServerType1CPU2GBRam20GBDisk   = 1
	CX21ServerType2CPU4GBRam40GBDisk   = 3
	CX31ServerType2CPU8GBRam80GBDisk   = 5
	CX41ServerType4CPU16GBRam160GBDisk = 7
	CX51ServerType8CPU32GBRam240GBDisk = 9
)

type Servercontroler interface {
	GetAllServers() ([]ServerInfo, error)
	CreateServer(name string, serverType int) (*ServerInfo, error)
	DeleteServer(ID int) error
	GetServerById(ID int) (ServerInfo, error)
}

type Servercontrol struct {
	client *hcloud.Client
}

func NewServercontrol(token string) *Servercontrol {
	s := &Servercontrol{}
	s.client = hcloud.NewClient(hcloud.WithToken(token))
	return s
}

func (s Servercontrol) PrintTypes() {
	types, _ := s.client.ServerType.All(context.Background())
	for _, v := range types {
		fmt.Println(*v)
	}
}

func (s Servercontrol) PrintImages() {
	images, _ := s.client.Image.All(context.Background())
	for _, v := range images {
		fmt.Println(*v)
	}
}

func (s Servercontrol) PrintSSHKeys() {
	keys, _ := s.client.SSHKey.All(context.Background())
	for _, v := range keys {
		fmt.Println(*v)
	}
}

func (s Servercontrol) PrintLocations() {
	locs, _ := s.client.Location.All(context.Background())
	for _, v := range locs {
		fmt.Println(*v)
	}
}

func (s Servercontrol) PrintDatacenter() {
	dcs, _ := s.client.Datacenter.All(context.Background())
	for _, v := range dcs {
		fmt.Println(*v)
	}
}

func (s Servercontrol) GetAllServers() ([]ServerInfo, error) {
	servers, err := s.client.Server.All(context.Background())
	if err != nil {
		log.Print(err)
		return nil, err
	}

	infos := []ServerInfo{}

	for _, v := range servers {
		infos = append(infos, ServerInfo{
			ID:             v.ID,
			Name:           v.Name,
			IP:             v.PublicNet.IPv4.IP.String(),
			ServerTypeID:   v.ServerType.ID,
			ServerTypeName: v.ServerType.Name,
		})
	}
	return infos, nil
}

func (s Servercontrol) CreateServer(name string, serverType int) (*ServerInfo, error) {
	dcFsn, _, _ := s.client.Datacenter.GetByID(context.Background(), 4)
	opts := hcloud.ServerCreateOpts{}
	opts.Name = name
	opts.ServerType = &hcloud.ServerType{ID: serverType}
	opts.Image = &hcloud.Image{ID: 5924233} // Debian 10
	opts.SSHKeys = []*hcloud.SSHKey{}
	opts.SSHKeys = append(opts.SSHKeys, &hcloud.SSHKey{ID: 1686255}) // workspace
	opts.SSHKeys = append(opts.SSHKeys, &hcloud.SSHKey{ID: 1754781}) // marc@marcbox
	opts.Datacenter = dcFsn

	result, _, err := s.client.Server.Create(context.Background(), opts)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	info := ServerInfo{
		ID:             result.Server.ID,
		Name:           result.Server.Name,
		IP:             result.Server.PublicNet.IPv4.IP.String(),
		ServerTypeID:   result.Server.ServerType.ID,
		ServerTypeName: result.Server.ServerType.Name,
	}

	return &info, nil
}

func (s Servercontrol) DeleteServer(ID int) error {
	server := hcloud.Server{ID: ID}
	_, err := s.client.Server.Delete(context.Background(), &server)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s Servercontrol) GetServerById(ID int) (ServerInfo, error) {
	server, _, err := s.client.Server.GetByID(context.Background(), ID)
	if err != nil {
		log.Print(err)
		return ServerInfo{}, err
	}

	return ServerInfo{
		ID:             server.ID,
		Name:           server.Name,
		IP:             server.PublicNet.IPv4.IP.String(),
		ServerTypeID:   server.ServerType.ID,
		ServerTypeName: server.ServerType.Name,
	}, nil
}
