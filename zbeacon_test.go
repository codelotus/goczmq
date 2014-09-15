package goczmq

import (
	"testing"
)

func TestZbeacon(t *testing.T) {
	// Create a Zbeacon
	speaker := NewZbeacon()

	err := speaker.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	_, err = speaker.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	listener := NewZbeacon()
	err = listener.Verbose()
	if err != nil {
		t.Errorf("VERBOSE error: %s", err)
	}

	_, err = listener.Configure(9999)
	if err != nil {
		t.Errorf("CONFIGURE error: %s", err)
	}

	err = listener.Subscribe("HI")
	if err != nil {
		t.Errorf("SUBSCRIBE error: %s", err)
	}

	speaker.Publish("HI", 100)
	reply := listener.Recv(500)
	t.Logf("%v", reply)
	listener.Destroy()
	speaker.Destroy()
}
