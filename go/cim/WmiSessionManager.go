package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiSessionManager struct {
}

func (c WmiSessionManager) GetSession(serverName, wmiNamespace, userName, password, domain string) (wmi.Session, error) {
	panic("not implemented")
}
