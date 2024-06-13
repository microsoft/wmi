// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PS_VpnConnectionRoute struct
type PS_VpnConnectionRoute struct {
	*cim.WmiInstance

	//
	AddressFamily uint16

	//
	DestinationPrefix string

	//
	RouteMetric uint32
}

func NewPS_VpnConnectionRouteEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnectionRoute, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionRoute{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnConnectionRouteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnConnectionRoute, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionRoute{
		WmiInstance: tmp,
	}
	return
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *PS_VpnConnectionRoute) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", (value))
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *PS_VpnConnectionRoute) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDestinationPrefix sets the value of DestinationPrefix for the instance
func (instance *PS_VpnConnectionRoute) SetPropertyDestinationPrefix(value string) (err error) {
	return instance.SetProperty("DestinationPrefix", (value))
}

// GetDestinationPrefix gets the value of DestinationPrefix for the instance
func (instance *PS_VpnConnectionRoute) GetPropertyDestinationPrefix() (value string, err error) {
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

// SetRouteMetric sets the value of RouteMetric for the instance
func (instance *PS_VpnConnectionRoute) SetPropertyRouteMetric(value uint32) (err error) {
	return instance.SetProperty("RouteMetric", (value))
}

// GetRouteMetric gets the value of RouteMetric for the instance
func (instance *PS_VpnConnectionRoute) GetPropertyRouteMetric() (value uint32, err error) {
	retValue, err := instance.GetProperty("RouteMetric")
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

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="ConnectionName" type="string "></param>
// <param name="DestinationPrefix" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RouteMetric" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_NetRoute "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionRoute) Add( /* IN */ ConnectionName string,
	/* IN */ DestinationPrefix string,
	/* IN */ RouteMetric uint32,
	/* IN */ PassThru bool,
	/* IN */ AllUserConnection bool,
	/* OUT */ cmdletOutput MSFT_NetRoute) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ConnectionName, DestinationPrefix, RouteMetric, PassThru, AllUserConnection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="ConnectionName" type="string "></param>
// <param name="DestinationPrefix" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetRoute "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionRoute) Remove( /* IN */ ConnectionName string,
	/* IN */ DestinationPrefix string,
	/* IN */ AllUserConnection bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput MSFT_NetRoute) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ConnectionName, DestinationPrefix, AllUserConnection, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
