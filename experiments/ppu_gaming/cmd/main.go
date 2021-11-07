package main

import (
	"fmt"
	"log"

	"github.com/ungefaehrlich/ppu_gaming/internal/rexec"
	"github.com/ungefaehrlich/ppu_gaming/internal/servercontrol"
	"github.com/ungefaehrlich/ppu_gaming/internal/web"
)

func webber() {
	web.RunAndServe()
}

func control() {
	newServer, err := servercontrol.CreateServer("w00p")
	if err != nil {
		return
	}
	fmt.Println(newServer.ID, newServer.Name)

	servercontrol.DeleteServer(newServer.ID)
	if err != nil {
		return
	}
	servers, err := servercontrol.GetAllServers()
	if err != nil {
		return
	}
	for _, v := range servers {
		fmt.Println(v.ID, v.Name, v.IP)
	}
}

func main() {
	newServer, err := servercontrol.CreateServer("autoMine")
	if err != nil {
		return
	}

	setupServer := `
		mkdir /opt/minecraft
		echo "eula=true" > /opt/minecraft/eula.txt
		wget https://launcher.mojang.com/v1/objects/a412fd69db1f81db3f511c1463fd304675244077/server.jar
		mv server.jar /opt/minecraft/
		apt-get update 
		apt-get install screen default-jre -y
		screen -dm bash -c "cd /opt/minecraft; java -Xmx1024M -Xms1024M -jar server.jar nogui"
		`

	//out, err := rexec.Rexec("116.202.107.212", setupServer)
	out, err := rexec.Rexec(newServer.IP, setupServer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}
