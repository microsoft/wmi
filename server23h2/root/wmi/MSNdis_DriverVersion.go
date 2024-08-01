// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_DriverVersion struct
type MSNdis_DriverVersion struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdisDriverVersion uint16
}

func NewMSNdis_DriverVersionEx1(instance *cim.WmiInstance) (newInstance *MSNdis_DriverVersion, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_DriverVersion{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_DriverVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_DriverVersion, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_DriverVersion{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_DriverVersion) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_DriverVersion) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_DriverVersion) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_DriverVersion) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetNdisDriverVersion sets the value of NdisDriverVersion for the instance
func (instance *MSNdis_DriverVersion) SetPropertyNdisDriverVersion(value uint16) (err error) {
	return instance.SetProperty("NdisDriverVersion", (value))
}

// GetNdisDriverVersion gets the value of NdisDriverVersion for the instance
func (instance *MSNdis_DriverVersion) GetPropertyNdisDriverVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("NdisDriverVersion")
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
