package cim

import (
	"errors"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/microsoft/wmicodegen/go/wmi"
)

const S_FALSE = 0x00000001

// Reference https://github.com/StackExchange/wmi
// Reference https://docs.microsoft.com/en-us/windows/desktop/WmiSdk/swbemlocator-connectserver

type WmiSessionManager struct {
	unknown  *ole.IUnknown
	wmi      *ole.IDispatch
	sessions map[string]*WmiSession
}

func (c WmiSessionManager) Init() error {
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		oleCode := err.(*ole.OleError).Code()
		if oleCode != ole.S_OK && oleCode != S_FALSE {
			return err
		}
	}

	c.unknown, err = oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		c.Dispose()
		return err
	}
	if c.unknown == nil {
		c.Dispose()
		return errors.New("CreateObject failed")
	}

	c.wmi, err = c.unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		c.Dispose()
		return err
	}

	return nil
}

// Dispose clears the WmiSessionManager
func (c WmiSessionManager) Dispose() {
	// clear the Sessions

	if c.wmi != nil {
		c.wmi.Release()
	}
	// clear ole object
	if c.unknown != nil {
		c.unknown.Release()
	}

	// clear com
	ole.CoUninitialize()
}

// GetSession
func (c WmiSessionManager) GetSession(serverName, wmiNamespace, userName, password, domain string) (wmi.Session, error) {

	return CreateSession(c.wmi, serverName, wmiNamespace, userName, password, domain)
}
