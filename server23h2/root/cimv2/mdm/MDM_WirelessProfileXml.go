// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_WirelessProfileXml struct
type MDM_WirelessProfileXml struct {
	*cim.WmiInstance

	//
	Name string

	//
	ProfileXml string
}

func NewMDM_WirelessProfileXmlEx1(instance *cim.WmiInstance) (newInstance *MDM_WirelessProfileXml, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WirelessProfileXml{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WirelessProfileXmlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WirelessProfileXml, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WirelessProfileXml{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MDM_WirelessProfileXml) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MDM_WirelessProfileXml) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetProfileXml sets the value of ProfileXml for the instance
func (instance *MDM_WirelessProfileXml) SetPropertyProfileXml(value string) (err error) {
	return instance.SetProperty("ProfileXml", (value))
}

// GetProfileXml gets the value of ProfileXml for the instance
func (instance *MDM_WirelessProfileXml) GetPropertyProfileXml() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileXml")
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
