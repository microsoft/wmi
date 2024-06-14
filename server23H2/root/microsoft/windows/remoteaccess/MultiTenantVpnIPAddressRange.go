// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MultiTenantVpnIPAddressRange struct
type MultiTenantVpnIPAddressRange struct {
	*VpnIPAddressRange

	//
	RoutingDomain string
}

func NewMultiTenantVpnIPAddressRangeEx1(instance *cim.WmiInstance) (newInstance *MultiTenantVpnIPAddressRange, err error) {
	tmp, err := NewVpnIPAddressRangeEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MultiTenantVpnIPAddressRange{
		VpnIPAddressRange: tmp,
	}
	return
}

func NewMultiTenantVpnIPAddressRangeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MultiTenantVpnIPAddressRange, err error) {
	tmp, err := NewVpnIPAddressRangeEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MultiTenantVpnIPAddressRange{
		VpnIPAddressRange: tmp,
	}
	return
}

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *MultiTenantVpnIPAddressRange) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *MultiTenantVpnIPAddressRange) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
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
