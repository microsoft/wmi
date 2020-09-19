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

// MSFT_NetVirtualizationCustomerRouteSettingData struct
type MSFT_NetVirtualizationCustomerRouteSettingData struct {
	*MSFT_NetSettingData

	// 32
	CustomerID string

	// 46
	DestinationPrefix string

	// 48
	Metric uint32

	// 47
	NextHop string

	// 49
	RoutingDomainID string

	// 45
	VirtualSubnetID uint32
}

func NewMSFT_NetVirtualizationCustomerRouteSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetVirtualizationCustomerRouteSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationCustomerRouteSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetVirtualizationCustomerRouteSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetVirtualizationCustomerRouteSettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetVirtualizationCustomerRouteSettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetCustomerID sets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyCustomerID(value string) (err error) {
	return instance.SetProperty("CustomerID", (value))
}

// GetCustomerID gets the value of CustomerID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyCustomerID() (value string, err error) {
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
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyDestinationPrefix(value string) (err error) {
	return instance.SetProperty("DestinationPrefix", (value))
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyDestinationPrefix() (value string, err error) {
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

// SetMetric sets the value of Metric for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyMetric(value uint32) (err error) {
	return instance.SetProperty("Metric", (value))
}

// GetMetric gets the value of Metric for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyMetric() (value uint32, err error) {
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
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", (value))
}

// GetNextHop gets the value of NextHop for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyNextHop() (value string, err error) {
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

// SetRoutingDomainID sets the value of RoutingDomainID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyRoutingDomainID(value string) (err error) {
	return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyRoutingDomainID() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainID")
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

// SetVirtualSubnetID sets the value of VirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) SetPropertyVirtualSubnetID(value uint32) (err error) {
	return instance.SetProperty("VirtualSubnetID", (value))
}

// GetVirtualSubnetID gets the value of VirtualSubnetID for the instance
func (instance *MSFT_NetVirtualizationCustomerRouteSettingData) GetPropertyVirtualSubnetID() (value uint32, err error) {
	retValue, err := instance.GetProperty("VirtualSubnetID")
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
