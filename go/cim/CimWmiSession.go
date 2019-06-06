package cim;

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type CimWmiSession struct {
}


// Name
func (c CimWmiSession)  Connect(serverName, wmiNamespace, userName, password, domain string) (bool, error) {
	panic("not implemented")
}

// Value
func (c CimWmiSession) Connect(serverName, wmiNamespace string, credentials WmiCredentials) (bool, error) {
	panic("not implemented")
}

// Dispose
func (c CimWmiSession) Dispose() {
	panic("not implemented")
}