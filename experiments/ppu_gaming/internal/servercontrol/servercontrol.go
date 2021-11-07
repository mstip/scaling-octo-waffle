package servercontrol

import (
	"context"
	"fmt"
	"log"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

type ServerInfo struct {
	ID   int
	Name string
	IP   string
}

var client *hcloud.Client

func init() {
	client = hcloud.NewClient(hcloud.WithToken("hgWYYM0SNZRsI7fZQmZf3qYLI5QFPeuTVGEnfNpgX8sNOcfi01VvpQFRcLjdQVoR"))
}

func PrintTypes() {
	types, _ := client.ServerType.All(context.Background())
	for _, v := range types {
		fmt.Println(*v)
	}
}

func PrintImages() {
	images, _ := client.Image.All(context.Background())
	for _, v := range images {
		fmt.Println(*v)
	}
}

func PrintSSHKeys() {
	keys, _ := client.SSHKey.All(context.Background())
	for _, v := range keys {
		fmt.Println(*v)
	}
}

func PrintLocations() {
	locs, _ := client.Location.All(context.Background())
	for _, v := range locs {
		fmt.Println(*v)
	}
}

func PrintDatacenter() {
	dcs, _ := client.Datacenter.All(context.Background())
	for _, v := range dcs {
		fmt.Println(*v)
	}
}

func GetAllServers() ([]ServerInfo, error) {
	servers, err := client.Server.All(context.Background())
	if err != nil {
		log.Print(err)
		return nil, err
	}

	infos := []ServerInfo{}

	for _, v := range servers {
		infos = append(infos, ServerInfo{ID: v.ID, Name: v.Name, IP: v.PublicNet.IPv4.IP.String()})
	}
	return infos, nil
}

func CreateServer(name string) (*ServerInfo, error) {
	dcFsn, _, _ := client.Datacenter.GetByID(context.Background(), 4)
	opts := hcloud.ServerCreateOpts{}
	opts.Name = name
	opts.ServerType = &hcloud.ServerType{ID: 1}
	opts.Image = &hcloud.Image{ID: 5924233} // Debian 10
	opts.SSHKeys = []*hcloud.SSHKey{}
	opts.SSHKeys = append(opts.SSHKeys, &hcloud.SSHKey{ID: 1686255}) // workspace
	opts.SSHKeys = append(opts.SSHKeys, &hcloud.SSHKey{ID: 1657531}) // dev
	opts.Datacenter = dcFsn

	result, _, err := client.Server.Create(context.Background(), opts)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	info := ServerInfo{ID: result.Server.ID, Name: result.Server.Name, IP: result.Server.PublicNet.IPv4.IP.String()}

	return &info, nil
}

func DeleteServer(ID int) error {
	server := hcloud.Server{ID: ID}
	_, err := client.Server.Delete(context.Background(), &server)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
