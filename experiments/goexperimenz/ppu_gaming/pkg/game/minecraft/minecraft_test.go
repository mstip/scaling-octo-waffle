package minecraft

import (
	"github.com/ungefaehrlich/ppu_gaming/pkg/servercontrol"
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestGetJVMRam(t *testing.T) {
	assert := testifyAssert.New(t)
	// unknown
	assert.Equal("", GetJVMRam(1337))
	// expected
	assert.Equal("1536", GetJVMRam(servercontrol.CX11ServerType1CPU2GBRam20GBDisk))
	assert.Equal("3584", GetJVMRam(servercontrol.CX21ServerType2CPU4GBRam40GBDisk))
	assert.Equal("7168", GetJVMRam(servercontrol.CX31ServerType2CPU8GBRam80GBDisk))
	assert.Equal("15360", GetJVMRam(servercontrol.CX41ServerType4CPU16GBRam160GBDisk))
	assert.Equal("30720", GetJVMRam(servercontrol.CX51ServerType8CPU32GBRam240GBDisk))
}
