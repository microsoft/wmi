// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnIPAddressRange struct
type VpnIPAddressRange struct {
	*cim.WmiInstance

	//
	EndIPAddress string

	//
	StartIPAddress string
}

func NewVpnIPAddressRangeEx1(instance *cim.WmiInstance) (newInstance *VpnIPAddressRange, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnIPAddressRange{
		WmiInstance: tmp,
	}
	return
}

func NewVpnIPAddressRangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnIPAddressRange, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnIPAddressRange{
		WmiInstance: tmp,
	}
	return
}

// SetEndIPAddress sets the value of EndIPAddress for the instance
func (instance *VpnIPAddressRange) SetPropertyEndIPAddress(value string) (err error) {
	return instance.SetProperty("EndIPAddress", (value))
}

// GetEndIPAddress gets the value of EndIPAddress for the instance
func (instance *VpnIPAddressRange) GetPropertyEndIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("EndIPAddress")
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

// SetStartIPAddress sets the value of StartIPAddress for the instance
func (instance *VpnIPAddressRange) SetPropertyStartIPAddress(value string) (err error) {
	return instance.SetProperty("StartIPAddress", (value))
}

// GetStartIPAddress gets the value of StartIPAddress for the instance
func (instance *VpnIPAddressRange) GetPropertyStartIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("StartIPAddress")
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
