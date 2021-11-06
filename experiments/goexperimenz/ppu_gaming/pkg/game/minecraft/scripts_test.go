package minecraft

import (
	"fmt"
	testifyAssert "github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupServerScript(t *testing.T) {
	assert := testifyAssert.New(t)
	assert.Contains(
		SetupServerScript("test-jvm-ram-1337"),
		fmt.Sprintf("java -Xmx%sM -Xms%sM -jar server.jar nogui", "test-jvm-ram-1337", "test-jvm-ram-1337"),
	)
}
