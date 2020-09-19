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

// MSFT_NetVirtualizationProviderRouteSettingData struct
type MSFT_NetVirtualizationProviderRouteSettingData struct {
	*MSFT_NetSettingData

	// 32
	CustomerID string

	// 46
	DestinationPrefix string

	// 12
	InterfaceIndex uint32

	// 48
	Metric uint32

	// 47
	NextHop string
}

func NewMSFT_NetVirtualizationProviderRouteSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationProviderRouteSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationProviderRouteSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationProviderRouteSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationProviderRouteSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationProviderRouteSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetCustomerID sets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) SetPropertyCustomerID(value string) (err error) {
	return instance.SetProperty("CustomerID", (value))
}

// GetCustomerID gets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) GetPropertyCustomerID() (value string, err error) {
	retValue, err := instance.GetProperty("CustomerID")
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

// SetDestinationPrefix sets the value of DestinationPrefix for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) SetPropertyDestinationPrefix(value string) (err error) {
	return instance.SetProperty("DestinationPrefix", (value))
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) GetPropertyDestinationPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("DestinationPrefix")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetMetric sets the value of Metric for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", (value))
}

// GetMetric gets the value of Metric for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) GetPropertyMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("Metric")
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

// SetNextHop sets the value of NextHop for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", (value))
}

// GetNextHop gets the value of NextHop for the instance
func (instance *MSFT_NetVirtualizationProviderRouteSettingData) GetPropertyNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NextHop")
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
