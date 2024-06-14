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

// RoutingDomainConfiguration struct
type RoutingDomainConfiguration struct {
	*cim.WmiInstance

	//
	RoutingDomain string

	//
	RoutingDomainID string

	//
	RoutingDomainStatus string

	//
	RoutingStatus uint32

	//
	VpnS2SStatus uint32

	//
	VpnStatus uint32
}

func NewRoutingDomainConfigurationEx1(instance *cim.WmiInstance) (newInstance *RoutingDomainConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewRoutingDomainConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RoutingDomainConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *RoutingDomainConfiguration) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *RoutingDomainConfiguration) GetPropertyRoutingDomain() (value string, err error) {
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

// SetRoutingDomainID sets the value of RoutingDomainID for the instance
func (instance *RoutingDomainConfiguration) SetPropertyRoutingDomainID(value string) (err error) {
	return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *RoutingDomainConfiguration) GetPropertyRoutingDomainID() (value string, err error) {
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

// SetRoutingDomainStatus sets the value of RoutingDomainStatus for the instance
func (instance *RoutingDomainConfiguration) SetPropertyRoutingDomainStatus(value string) (err error) {
	return instance.SetProperty("RoutingDomainStatus", (value))
}

// GetRoutingDomainStatus gets the value of RoutingDomainStatus for the instance
func (instance *RoutingDomainConfiguration) GetPropertyRoutingDomainStatus() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainStatus")
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

// SetRoutingStatus sets the value of RoutingStatus for the instance
func (instance *RoutingDomainConfiguration) SetPropertyRoutingStatus(value uint32) (err error) {
	return instance.SetProperty("RoutingStatus", (value))
}

// GetRoutingStatus gets the value of RoutingStatus for the instance
func (instance *RoutingDomainConfiguration) GetPropertyRoutingStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("RoutingStatus")
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

// SetVpnS2SStatus sets the value of VpnS2SStatus for the instance
func (instance *RoutingDomainConfiguration) SetPropertyVpnS2SStatus(value uint32) (err error) {
	return instance.SetProperty("VpnS2SStatus", (value))
}

// GetVpnS2SStatus gets the value of VpnS2SStatus for the instance
func (instance *RoutingDomainConfiguration) GetPropertyVpnS2SStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("VpnS2SStatus")
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

// SetVpnStatus sets the value of VpnStatus for the instance
func (instance *RoutingDomainConfiguration) SetPropertyVpnStatus(value uint32) (err error) {
	return instance.SetProperty("VpnStatus", (value))
}

// GetVpnStatus gets the value of VpnStatus for the instance
func (instance *RoutingDomainConfiguration) GetPropertyVpnStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("VpnStatus")
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
