// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAddressFilter struct
type MSFT_NetAddressFilter struct {
	CIM_FilterEntryBase

	//
	LocalAddress []string

	//
	RemoteAddress []string
}

// SetLocalAddress sets the value of LocalAddress for the instance
func (instance *MSFT_NetAddressFilter) SetPropertyLocalAddress(value []string) (err error) {
	return instance.SetProperty("LocalAddress", value)
}

// GetLocalAddress gets the value of LocalAddress for the instance
func (instance *MSFT_NetAddressFilter) GetPropertyLocalAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteAddress sets the value of RemoteAddress for the instance
func (instance *MSFT_NetAddressFilter) SetPropertyRemoteAddress(value []string) (err error) {
	return instance.SetProperty("RemoteAddress", value)
}

// GetRemoteAddress gets the value of RemoteAddress for the instance
func (instance *MSFT_NetAddressFilter) GetPropertyRemoteAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoteAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="RemoteAddress" type="string "></param>

// <param name="IsolationType" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAddressFilter) QueryIsolationType( /* IN */ InterfaceIndex uint32,
	/* IN */ RemoteAddress string,
	/* OUT */ IsolationType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("QueryIsolationType", InterfaceIndex, RemoteAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
