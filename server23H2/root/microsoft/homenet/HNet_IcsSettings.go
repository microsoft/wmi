// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_IcsSettings struct
type HNet_IcsSettings struct {
	*cim.WmiInstance

	//
	DhcpEnabled bool

	//
	DnsEnabled bool

	//
	Id string
}

func NewHNet_IcsSettingsEx1(instance *cim.WmiInstance) (newInstance *HNet_IcsSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_IcsSettings{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_IcsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_IcsSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_IcsSettings{
		WmiInstance: tmp,
	}
	return
}

// SetDhcpEnabled sets the value of DhcpEnabled for the instance
func (instance *HNet_IcsSettings) SetPropertyDhcpEnabled(value bool) (err error) {
	return instance.SetProperty("DhcpEnabled", (value))
}

// GetDhcpEnabled gets the value of DhcpEnabled for the instance
func (instance *HNet_IcsSettings) GetPropertyDhcpEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DhcpEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDnsEnabled sets the value of DnsEnabled for the instance
func (instance *HNet_IcsSettings) SetPropertyDnsEnabled(value bool) (err error) {
	return instance.SetProperty("DnsEnabled", (value))
}

// GetDnsEnabled gets the value of DnsEnabled for the instance
func (instance *HNet_IcsSettings) GetPropertyDnsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DnsEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *HNet_IcsSettings) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *HNet_IcsSettings) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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
