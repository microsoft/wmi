// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DNSProtocolEndpoint struct
type CIM_DNSProtocolEndpoint struct {
	*CIM_ProtocolEndpoint

	// 650
	DHCPOptionsToUse []DNSProtocolEndpoint_DHCPOptionsToUse

	// 649
	Hostname string
}

func NewCIM_DNSProtocolEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_DNSProtocolEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DNSProtocolEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_DNSProtocolEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DNSProtocolEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DNSProtocolEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetDHCPOptionsToUse sets the value of DHCPOptionsToUse for the instance
func (instance *CIM_DNSProtocolEndpoint) SetPropertyDHCPOptionsToUse(value []DNSProtocolEndpoint_DHCPOptionsToUse) (err error) {
	return instance.SetProperty("DHCPOptionsToUse", (value))
}

// GetDHCPOptionsToUse gets the value of DHCPOptionsToUse for the instance
func (instance *CIM_DNSProtocolEndpoint) GetPropertyDHCPOptionsToUse() (value []DNSProtocolEndpoint_DHCPOptionsToUse, err error) {
	retValue, err := instance.GetProperty("DHCPOptionsToUse")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DNSProtocolEndpoint_DHCPOptionsToUse(valuetmp))
	}

	return
}

// SetHostname sets the value of Hostname for the instance
func (instance *CIM_DNSProtocolEndpoint) SetPropertyHostname(value string) (err error) {
	return instance.SetProperty("Hostname", (value))
}

// GetHostname gets the value of Hostname for the instance
func (instance *CIM_DNSProtocolEndpoint) GetPropertyHostname() (value string, err error) {
	retValue, err := instance.GetProperty("Hostname")
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
