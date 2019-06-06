package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiSession struct {
}

// Connect
func (c WmiSession) Connect(serverName, wmiNamespace, userName, password, domain string) (bool, error) {
	panic("not implemented")
}

// ConnectEx
func (c WmiSession) ConnectEx(serverName, wmiNamespace string, credentials wmi.Credentials) (bool, error) {
	panic("not implemented")
}

// Dispose
func (c WmiSession) Dispose() {
	panic("not implemented")
}
