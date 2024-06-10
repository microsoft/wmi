// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnS2SMultiTenantInterfaceStatistics struct
type VpnS2SMultiTenantInterfaceStatistics struct {
	*VpnS2SInterfaceStatistics

	//
	Bandwidth uint64
}

func NewVpnS2SMultiTenantInterfaceStatisticsEx1(instance *cim.WmiInstance) (newInstance *VpnS2SMultiTenantInterfaceStatistics, err error) {
	tmp, err := NewVpnS2SInterfaceStatisticsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantInterfaceStatistics{
		VpnS2SInterfaceStatistics: tmp,
	}
	return
}

func NewVpnS2SMultiTenantInterfaceStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnS2SMultiTenantInterfaceStatistics, err error) {
	tmp, err := NewVpnS2SInterfaceStatisticsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnS2SMultiTenantInterfaceStatistics{
		VpnS2SInterfaceStatistics: tmp,
	}
	return
}

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *VpnS2SMultiTenantInterfaceStatistics) SetPropertyBandwidth(value uint64) (err error) {
	return instance.SetProperty("Bandwidth", (value))
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *VpnS2SMultiTenantInterfaceStatistics) GetPropertyBandwidth() (value uint64, err error) {
	retValue, err := instance.GetProperty("Bandwidth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
