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

// MSFC_FibrePortNPIVMethodsEx struct
type MSFC_FibrePortNPIVMethodsEx struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMSFC_FibrePortNPIVMethodsExEx1(instance *cim.WmiInstance) (newInstance *MSFC_FibrePortNPIVMethodsEx, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVMethodsEx{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FibrePortNPIVMethodsExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FibrePortNPIVMethodsEx, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortNPIVMethodsEx{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVMethodsEx) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FibrePortNPIVMethodsEx) GetPropertyActive() (value bool, err error) {
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
func (instance *MSFC_FibrePortNPIVMethodsEx) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FibrePortNPIVMethodsEx) GetPropertyInstanceName() (value string, err error) {
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

// <param name="CHAP" type="MSFC_DH_Chap_Parameters "></param>

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) SetChapForPhysicalPort( /* OUT */ Status FibrePortNPIVMethodsEx_Status,
	/* OPTIONAL IN */ CHAP MSFC_DH_Chap_Parameters) (err error) {
	_, err = instance.InvokeMethod("SetChapForPhysicalPort", CHAP)
	if err != nil {
		return
	}
	return

}

//

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) RemoveChapForPhysicalPort( /* OUT */ Status FibrePortNPIVMethodsEx_Status) (err error) {
	_, err = instance.InvokeMethod("RemoveChapForPhysicalPort")
	if err != nil {
		return
	}
	return

}

//

// <param name="Tag" type="uint8 []"></param>
// <param name="VirtualName" type="uint16 []"></param>
// <param name="WWNN" type="uint8 []"></param>
// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) CreateVirtualPortEx( /* OUT */ Status FibrePortNPIVMethodsEx_Status,
	/* OPTIONAL IN */ WWPN []uint8,
	/* OPTIONAL IN */ WWNN []uint8,
	/* OPTIONAL IN */ Tag []uint8,
	/* OPTIONAL IN */ VirtualName []uint16) (err error) {
	_, err = instance.InvokeMethod("CreateVirtualPortEx", WWPN, WWNN, Tag, VirtualName)
	if err != nil {
		return
	}
	return

}

//

// <param name="CHAP" type="MSFC_DH_Chap_Parameters "></param>
// <param name="Tag" type="uint8 []"></param>
// <param name="VirtualName" type="uint16 []"></param>
// <param name="WWNN" type="uint8 []"></param>
// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) CreateVirtualPortExUsingDHCHAP( /* OUT */ Status FibrePortNPIVMethodsEx_Status,
	/* OPTIONAL IN */ WWPN []uint8,
	/* OPTIONAL IN */ WWNN []uint8,
	/* OPTIONAL IN */ Tag []uint8,
	/* OPTIONAL IN */ VirtualName []uint16,
	/* OPTIONAL IN */ CHAP MSFC_DH_Chap_Parameters) (err error) {
	_, err = instance.InvokeMethod("CreateVirtualPortExUsingDHCHAP", WWPN, WWNN, Tag, VirtualName, CHAP)
	if err != nil {
		return
	}
	return

}

//

// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) RemoveVirtualPortEx( /* OUT */ Status FibrePortNPIVMethodsEx_Status,
	/* OPTIONAL IN */ WWPN []uint8) (err error) {
	_, err = instance.InvokeMethod("RemoveVirtualPortEx", WWPN)
	if err != nil {
		return
	}
	return

}

//

// <param name="WWPN" type="uint8 []"></param>

// <param name="Status" type="FibrePortNPIVMethodsEx_Status "></param>
func (instance *MSFC_FibrePortNPIVMethodsEx) RescanVirtualPort( /* OUT */ Status FibrePortNPIVMethodsEx_Status,
	/* OPTIONAL IN */ WWPN []uint8) (err error) {
	_, err = instance.InvokeMethod("RescanVirtualPort", WWPN)
	if err != nil {
		return
	}
	return

}
