package main

import (
	"github.com/jbweber/packer-builder-qemuipxe/qemuipxe"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(qemuipxe.Builder))
	server.Serve()
}
