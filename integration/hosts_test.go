package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestHost_Mute(t *testing.T) {
	muteAction := getTestMuteAction()
	hostname := "test.host"
	muteResp, err := client.MuteHost(hostname, muteAction)
	if err != nil {
		t.Fatalf("Failed to mute host: %s, err: %s", hostname, err)
	}

	defer func() {
		unmuteResp, err := client.UnmuteHost(hostname)
		if err != nil {
			t.Fatalf("Failed to cleanup mute on host: %s, err: %s", hostname, err)
		}
		assert.Equal(t, "Unmuted", unmuteResp.Action)
		assert.Equal(t, hostname, unmuteResp.Hostname)
	}()

	assert.Equal(t, "Muted", muteResp.Action)
	assert.Equal(t, hostname, muteResp.Hostname)
	assert.Equal(t, "Muting this host for a test!", muteResp.Message)

}

func getTestMuteAction() *datadog.HostActionMute {
	return &datadog.HostActionMute{
		Message:  datadog.String("Muting this host for a test!"),
		EndTime:  datadog.String(fmt.Sprint(time.Now().Unix() + 100)),
		Override: datadog.Bool(false),
	}
}
