// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_FibrePortNPIVMethods struct
type MSFC_FibrePortNPIVMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMSFC_FibrePortNPIVMethodsEx1(instance *cim.WmiInstance) (newInstance *MSFC_FibrePortNPIVMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVMethods{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FibrePortNPIVMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FibrePortNPIVMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVMethods) GetPropertyActive() (value bool, err error) {
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
func (instance *MSFC_FibrePortNPIVMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FibrePortNPIVMethods) GetPropertyInstanceName() (value string, err error) {
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

//

// <param name="Tag" type="uint8 []"></param>
// <param name="VirtualName" type="uint16 []"></param>
// <param name="WWNN" type="uint8 []"></param>
// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethods_Status "></param>
func (instance *MSFC_FibrePortNPIVMethods) CreateVirtualPort( /* OUT */ Status FibrePortNPIVMethods_Status,
	/* OPTIONAL IN */ WWPN []uint8,
	/* OPTIONAL IN */ WWNN []uint8,
	/* OPTIONAL IN */ Tag []uint8,
	/* OPTIONAL IN */ VirtualName []uint16) (err error) {
	_, err = instance.InvokeMethod("CreateVirtualPort", WWPN, WWNN, Tag, VirtualName)
	if err != nil {
		return
	}
	return

}

//

// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethods_Status "></param>
func (instance *MSFC_FibrePortNPIVMethods) RemoveVirtualPort( /* OUT */ Status FibrePortNPIVMethods_Status,
	/* OPTIONAL IN */ WWPN []uint8) (err error) {
	_, err = instance.InvokeMethod("RemoveVirtualPort", WWPN)
	if err != nil {
		return
	}
	return

}
