package cim;

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type CimWmiSessionManager struct {

}

func (c CimWmiSessionManager)  GetSession(serverName, wmiNamespace, userName, password, domain string) (CimWmiSession, error) {
	panic("not implemented")
}
