package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server EchoServer) Handle(method, params string) *Response {
	// return "ECHO" + method
	body := "ECHO:" + method
	return &Response{Body: body}
}

func (server EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	ipcServer := NewIpcServer(EchoServer{})

	client1 := NewIpcClient(ipcServer)
	client2 := NewIpcClient(ipcServer)
	defer client1.Close()
	defer client2.Close()

	resp1, _ := client1.Call("From Client1", "")
	resp2, _ := client1.Call("From Client2", "")

	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
	}

	// client1.Close()
	// client2.Close()
}
