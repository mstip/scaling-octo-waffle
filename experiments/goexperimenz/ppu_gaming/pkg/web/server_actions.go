package web

import (
	"github.com/ungefaehrlich/ppu_gaming/pkg/game/minecraft"
	"log"
	"strconv"
	"strings"
	"time"
)

func (s *server) listSavesByCustomerID(customerID string) ([]string, error) {
	out, err := s.storeServer.Exec(s.storeServerIP, minecraft.ListSavesByCustomerIDScript(customerID))
	if err != nil {
		return nil, err
	}

	var saves []string

	for _, v := range strings.Split(out, "\n") {
		if len(v) == 0 {
			continue
		}
		splitted := strings.Split(v, "/")
		saves = append(saves, splitted[len(splitted)-1])
	}

	return saves, nil
}

func (s *server) createServerAction(serverType int) error {
	customerId := "1"
	game := "minecraftvanilla"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	log.Println("create server action " + game + " for customer " + customerId)

	newServer, err := s.sc.CreateServer(customerId+"-"+game+"-"+timestamp, serverType)
	if err != nil {
		log.Println(err)
		return err
	}

	jvmRam := minecraft.GetJVMRam(serverType)

	out, err := s.gameServer.Exec(newServer.IP, minecraft.SetupServerScript(jvmRam))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Print("finished create server")
	log.Print(out)
	return nil
}

func (s *server) saveAndDeleteServerAction(serverID string) error {
	log.Print("save and delete server" + serverID)

	ID, err := strconv.Atoi(serverID)
	if err != nil {
		log.Println(err)
		return err
	}

	server, err := s.sc.GetServerById(ID)
	if err != nil {
		log.Println(err)
		return err
	}
	customerID := "1"
	game := "minecraftvanilla"
	log.Println(minecraft.SaveScript(customerID, game))
	out, err := s.gameServer.Exec(server.IP, minecraft.SaveScript(customerID, game))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Print("save + copy success", out)

	if err := s.sc.DeleteServer(ID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *server) deleteServerAction(serverID string) {
	ID, err := strconv.Atoi(serverID)
	if err != nil {
		log.Println(err)
		return
	}
	if err := s.sc.DeleteServer(ID); err != nil {
		log.Println(err)
		return
	}
}

func (s *server) loadSaveInNewServerAction(serverType int, saveFileName string) error {
	customerId := "1"
	game := "minecraftvanilla"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	log.Println("loadSaveInNewServerAction " + game + " for customer " + customerId + " save game " + saveFileName)

	newServer, err := s.sc.CreateServer(customerId+"-"+game+"-"+timestamp, serverType)
	if err != nil {
		log.Println(err)
		return err
	}

	jvmRam := minecraft.GetJVMRam(serverType)

	log.Print(minecraft.SetupServerAndLoadSaveScript(saveFileName, jvmRam))
	out, err := s.gameServer.Exec(newServer.IP, minecraft.SetupServerAndLoadSaveScript(saveFileName, jvmRam))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Print("loadSaveInNewServerAction success")
	log.Print(out)
	return nil
}

func (s *server) deleteSaveAction(saveFileName string) error {
	_, err := s.storeServer.Exec(s.storeServerIP, minecraft.DeleteSaveScript(saveFileName))
	if err != nil {
		return err
	}

	return nil
}
