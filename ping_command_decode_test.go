package pulsar_proto_go

import (
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPingCommandDecode(t *testing.T) {
	cmd := &BaseCommand{}
	bytes := testHex2Bytes(t, "000000050812920100")
	err := proto.Unmarshal(bytes[4:], cmd)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, BaseCommand_PING, *cmd.Type)
}
