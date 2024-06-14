// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Gateway
//////////////////////////////////////////////
package gateway

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// GatewayRoutingDomainConfiguration struct
type GatewayRoutingDomainConfiguration struct {
	*cim.WmiInstance

	//
	InterfaceName string

	//
	RoutingDomain string

	//
	RoutingDomainID string

	//
	RoutingDomainStatus string

	//
	VpnS2SStatus uint32

	//
	VpnStatus uint32
}

func NewGatewayRoutingDomainConfigurationEx1(instance *cim.WmiInstance) (newInstance *GatewayRoutingDomainConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &GatewayRoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewGatewayRoutingDomainConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *GatewayRoutingDomainConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &GatewayRoutingDomainConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetInterfaceName sets the value of InterfaceName for the instance
func (instance *GatewayRoutingDomainConfiguration) SetPropertyInterfaceName(value string) (err error) {
	return instance.SetProperty("InterfaceName", (value))
}

// GetInterfaceName gets the value of InterfaceName for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyInterfaceName() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceName")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *GatewayRoutingDomainConfiguration) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyRoutingDomain() (value string, err error) {
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
func (instance *GatewayRoutingDomainConfiguration) SetPropertyRoutingDomainID(value string) (err error) {
	return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyRoutingDomainID() (value string, err error) {
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
func (instance *GatewayRoutingDomainConfiguration) SetPropertyRoutingDomainStatus(value string) (err error) {
	return instance.SetProperty("RoutingDomainStatus", (value))
}

// GetRoutingDomainStatus gets the value of RoutingDomainStatus for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyRoutingDomainStatus() (value string, err error) {
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

// SetVpnS2SStatus sets the value of VpnS2SStatus for the instance
func (instance *GatewayRoutingDomainConfiguration) SetPropertyVpnS2SStatus(value uint32) (err error) {
	return instance.SetProperty("VpnS2SStatus", (value))
}

// GetVpnS2SStatus gets the value of VpnS2SStatus for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyVpnS2SStatus() (value uint32, err error) {
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
func (instance *GatewayRoutingDomainConfiguration) SetPropertyVpnStatus(value uint32) (err error) {
	return instance.SetProperty("VpnStatus", (value))
}

// GetVpnStatus gets the value of VpnStatus for the instance
func (instance *GatewayRoutingDomainConfiguration) GetPropertyVpnStatus() (value uint32, err error) {
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
