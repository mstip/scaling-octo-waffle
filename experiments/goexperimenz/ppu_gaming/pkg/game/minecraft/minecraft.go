package minecraft

import (
	"github.com/ungefaehrlich/ppu_gaming/pkg/servercontrol"
)

func GetJVMRam(serverType int) string {
	serverTypeJvmMap := map[int]string{
		servercontrol.CX11ServerType1CPU2GBRam20GBDisk:   "1536",
		servercontrol.CX21ServerType2CPU4GBRam40GBDisk:   "3584",
		servercontrol.CX31ServerType2CPU8GBRam80GBDisk:   "7168",
		servercontrol.CX41ServerType4CPU16GBRam160GBDisk: "15360",
		servercontrol.CX51ServerType8CPU32GBRam240GBDisk: "30720",
	}

	return serverTypeJvmMap[serverType]
}
