// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbMultichannelConnection struct
type MSFT_SmbMultichannelConnection struct {
	*cim.WmiInstance

	//
	ClientInterfaceFriendlyName string

	//
	ClientInterfaceIndex uint32

	//
	ClientIpAddress string

	//
	ClientLinkSpeed uint64

	//
	ClientRdmaCapable bool

	//
	ClientRSSCapable bool

	//
	CurrentChannels uint32

	//
	Failed bool

	//
	FailureCount uint32

	//
	MaxChannels uint32

	//
	Selected bool

	//
	ServerInterfaceIndex uint32

	//
	ServerIpAddress string

	//
	ServerLinkSpeed uint64

	//
	ServerName string

	//
	ServerRdmaCapable bool

	//
	ServerRSSCapable bool

	//
	SmbInstance SmbMultichannelConnection_SmbInstance
}

func NewMSFT_SmbMultichannelConnectionEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbMultichannelConnection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMultichannelConnection{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbMultichannelConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbMultichannelConnection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMultichannelConnection{
		WmiInstance: tmp,
	}
	return
}

// SetClientInterfaceFriendlyName sets the value of ClientInterfaceFriendlyName for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientInterfaceFriendlyName(value string) (err error) {
	return instance.SetProperty("ClientInterfaceFriendlyName", value)
}

// GetClientInterfaceFriendlyName gets the value of ClientInterfaceFriendlyName for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientInterfaceFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientInterfaceFriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientInterfaceIndex sets the value of ClientInterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("ClientInterfaceIndex", value)
}

// GetClientInterfaceIndex gets the value of ClientInterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientInterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientIpAddress sets the value of ClientIpAddress for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientIpAddress(value string) (err error) {
	return instance.SetProperty("ClientIpAddress", value)
}

// GetClientIpAddress gets the value of ClientIpAddress for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ClientIpAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientLinkSpeed sets the value of ClientLinkSpeed for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("ClientLinkSpeed", value)
}

// GetClientLinkSpeed gets the value of ClientLinkSpeed for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("ClientLinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientRdmaCapable sets the value of ClientRdmaCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientRdmaCapable(value bool) (err error) {
	return instance.SetProperty("ClientRdmaCapable", value)
}

// GetClientRdmaCapable gets the value of ClientRdmaCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientRdmaCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("ClientRdmaCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClientRSSCapable sets the value of ClientRSSCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyClientRSSCapable(value bool) (err error) {
	return instance.SetProperty("ClientRSSCapable", value)
}

// GetClientRSSCapable gets the value of ClientRSSCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyClientRSSCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("ClientRSSCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentChannels sets the value of CurrentChannels for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyCurrentChannels(value uint32) (err error) {
	return instance.SetProperty("CurrentChannels", value)
}

// GetCurrentChannels gets the value of CurrentChannels for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyCurrentChannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentChannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailed sets the value of Failed for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyFailed(value bool) (err error) {
	return instance.SetProperty("Failed", value)
}

// GetFailed gets the value of Failed for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyFailed() (value bool, err error) {
	retValue, err := instance.GetProperty("Failed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailureCount sets the value of FailureCount for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyFailureCount(value uint32) (err error) {
	return instance.SetProperty("FailureCount", value)
}

// GetFailureCount gets the value of FailureCount for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailureCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxChannels sets the value of MaxChannels for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyMaxChannels(value uint32) (err error) {
	return instance.SetProperty("MaxChannels", value)
}

// GetMaxChannels gets the value of MaxChannels for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyMaxChannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxChannels")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelected sets the value of Selected for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertySelected(value bool) (err error) {
	return instance.SetProperty("Selected", value)
}

// GetSelected gets the value of Selected for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertySelected() (value bool, err error) {
	retValue, err := instance.GetProperty("Selected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerInterfaceIndex sets the value of ServerInterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("ServerInterfaceIndex", value)
}

// GetServerInterfaceIndex gets the value of ServerInterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServerInterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerIpAddress sets the value of ServerIpAddress for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerIpAddress(value string) (err error) {
	return instance.SetProperty("ServerIpAddress", value)
}

// GetServerIpAddress gets the value of ServerIpAddress for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerIpAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerLinkSpeed sets the value of ServerLinkSpeed for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("ServerLinkSpeed", value)
}

// GetServerLinkSpeed gets the value of ServerLinkSpeed for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("ServerLinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerRdmaCapable sets the value of ServerRdmaCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerRdmaCapable(value bool) (err error) {
	return instance.SetProperty("ServerRdmaCapable", value)
}

// GetServerRdmaCapable gets the value of ServerRdmaCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerRdmaCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerRdmaCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerRSSCapable sets the value of ServerRSSCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertyServerRSSCapable(value bool) (err error) {
	return instance.SetProperty("ServerRSSCapable", value)
}

// GetServerRSSCapable gets the value of ServerRSSCapable for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertyServerRSSCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerRSSCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmbInstance sets the value of SmbInstance for the instance
func (instance *MSFT_SmbMultichannelConnection) SetPropertySmbInstance(value SmbMultichannelConnection_SmbInstance) (err error) {
	return instance.SetProperty("SmbInstance", value)
}

// GetSmbInstance gets the value of SmbInstance for the instance
func (instance *MSFT_SmbMultichannelConnection) GetPropertySmbInstance() (value SmbMultichannelConnection_SmbInstance, err error) {
	retValue, err := instance.GetProperty("SmbInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbMultichannelConnection_SmbInstance)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMultichannelConnection) Refresh( /* IN */ ServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh", ServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
