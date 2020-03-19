// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetNeighbor struct
type MSFT_NetNeighbor struct {
	*CIM_RemoteServiceAccessPoint

	//
	AddressFamily uint16

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	IPAddress string

	//
	LinkLayerAddress string

	//
	State uint8

	//
	Store uint8
}

func NewMSFT_NetNeighborEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetNeighbor, err error) {
	tmp, err := NewCIM_RemoteServiceAccessPointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNeighbor{
		CIM_RemoteServiceAccessPoint: tmp,
	}
	return
}

func NewMSFT_NetNeighborEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetNeighbor, err error) {
	tmp, err := NewCIM_RemoteServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetNeighbor{
		CIM_RemoteServiceAccessPoint: tmp,
	}
	return
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_NetNeighbor) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", value)
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_NetNeighbor) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetNeighbor) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", value)
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetNeighbor) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetNeighbor) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetNeighbor) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MSFT_NetNeighbor) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", value)
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MSFT_NetNeighbor) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkLayerAddress sets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetNeighbor) SetPropertyLinkLayerAddress(value string) (err error) {
	return instance.SetProperty("LinkLayerAddress", value)
}

// GetLinkLayerAddress gets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetNeighbor) GetPropertyLinkLayerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("LinkLayerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetNeighbor) SetPropertyState(value uint8) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetNeighbor) GetPropertyState() (value uint8, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetNeighbor) SetPropertyStore(value uint8) (err error) {
	return instance.SetProperty("Store", value)
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetNeighbor) GetPropertyStore() (value uint8, err error) {
	retValue, err := instance.GetProperty("Store")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AddressFamily" type="uint16 "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="IPAddress" type="string "></param>
// <param name="LinkLayerAddress" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyStore" type="string "></param>
// <param name="State" type="uint8 "></param>

// <param name="CmdletOutput" type="MSFT_NetNeighbor []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetNeighbor) Create( /* IN */ InterfaceIndex uint32,
	/* IN */ InterfaceAlias string,
	/* IN */ IPAddress string,
	/* IN */ PolicyStore string,
	/* IN */ LinkLayerAddress string,
	/* IN */ State uint8,
	/* IN */ AddressFamily uint16,
	/* IN */ PassThru bool,
	/* OUT */ CmdletOutput []MSFT_NetNeighbor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", InterfaceIndex, InterfaceAlias, IPAddress, PolicyStore, LinkLayerAddress, State, AddressFamily, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
