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

// CIM_ProtocolEndpoint struct
type CIM_ProtocolEndpoint struct {
	*CIM_ServiceAccessPoint

	//
	NameFormat string

	//
	OtherTypeDescription string

	//
	ProtocolIFType uint16

	//
	ProtocolType uint16
}

func NewCIM_ProtocolEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_ProtocolEndpoint, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolEndpoint{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

func NewCIM_ProtocolEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProtocolEndpoint, err error) {
	tmp, err := NewCIM_ServiceAccessPointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolEndpoint{
		CIM_ServiceAccessPoint: tmp,
	}
	return
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyNameFormat(value string) (err error) {
	return instance.SetProperty("NameFormat", (value))
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("NameFormat")
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

// SetOtherTypeDescription sets the value of OtherTypeDescription for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyOtherTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherTypeDescription", (value))
}

// GetOtherTypeDescription gets the value of OtherTypeDescription for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyOtherTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTypeDescription")
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

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyProtocolIFType(value uint16) (err error) {
	return instance.SetProperty("ProtocolIFType", (value))
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyProtocolIFType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
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

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *CIM_ProtocolEndpoint) SetPropertyProtocolType(value uint16) (err error) {
	return instance.SetProperty("ProtocolType", (value))
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *CIM_ProtocolEndpoint) GetPropertyProtocolType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProtocolType")
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
