// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetVirtualizationNextHopLookup struct
type MSFT_NetVirtualizationNextHopLookup struct {
	*MSFT_NetSettingData

	// 52
	DestinationCustomerAddress string

	// 53
	NextHopAddress string

	// 55
	NextHopMACAddress string

	// 51
	SourceCustomerAddress string

	// 54
	SourceMACAddress string

	// 31
	SourceVirtualSubnetID uint32
}

func NewMSFT_NetVirtualizationNextHopLookupEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationNextHopLookup, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationNextHopLookup{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationNextHopLookupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationNextHopLookup, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationNextHopLookup{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetDestinationCustomerAddress sets the value of DestinationCustomerAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertyDestinationCustomerAddress(value string) (err error) {
	return instance.SetProperty("DestinationCustomerAddress", (value))
}

// GetDestinationCustomerAddress gets the value of DestinationCustomerAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertyDestinationCustomerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationCustomerAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNextHopAddress sets the value of NextHopAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertyNextHopAddress(value string) (err error) {
	return instance.SetProperty("NextHopAddress", (value))
}

// GetNextHopAddress gets the value of NextHopAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertyNextHopAddress() (value string, err error) {
	retValue, err := instance.GetProperty("NextHopAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNextHopMACAddress sets the value of NextHopMACAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertyNextHopMACAddress(value string) (err error) {
	return instance.SetProperty("NextHopMACAddress", (value))
}

// GetNextHopMACAddress gets the value of NextHopMACAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertyNextHopMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("NextHopMACAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSourceCustomerAddress sets the value of SourceCustomerAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertySourceCustomerAddress(value string) (err error) {
	return instance.SetProperty("SourceCustomerAddress", (value))
}

// GetSourceCustomerAddress gets the value of SourceCustomerAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertySourceCustomerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("SourceCustomerAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSourceMACAddress sets the value of SourceMACAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertySourceMACAddress(value string) (err error) {
	return instance.SetProperty("SourceMACAddress", (value))
}

// GetSourceMACAddress gets the value of SourceMACAddress for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertySourceMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("SourceMACAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSourceVirtualSubnetID sets the value of SourceVirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) SetPropertySourceVirtualSubnetID(value uint32) (err error) {
	return instance.SetProperty("SourceVirtualSubnetID", (value))
}

// GetSourceVirtualSubnetID gets the value of SourceVirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationNextHopLookup) GetPropertySourceVirtualSubnetID() (value uint32, err error) {
	retValue, err := instance.GetProperty("SourceVirtualSubnetID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// 56

// <param name="DestinationCustomerAddress" type="string "></param>
// <param name="SourceCustomerAddress" type="string "></param>
// <param name="SourceVirtualSubnetID" type="uint32 "></param>

// <param name="CmdletOutput" type="MSFT_NetVirtualizationNextHopLookup []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetVirtualizationNextHopLookup) Select( /* IN */ SourceCustomerAddress string,
	/* IN */ DestinationCustomerAddress string,
	/* IN */ SourceVirtualSubnetID uint32,
	/* OUT */ CmdletOutput []MSFT_NetVirtualizationNextHopLookup) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Select", SourceCustomerAddress, DestinationCustomerAddress, SourceVirtualSubnetID)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
