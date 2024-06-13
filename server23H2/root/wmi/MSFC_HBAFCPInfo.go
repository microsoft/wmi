// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_HBAFCPInfo struct
type MSFC_HBAFCPInfo struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMSFC_HBAFCPInfoEx1(instance *cim.WmiInstance) (newInstance *MSFC_HBAFCPInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAFCPInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_HBAFCPInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_HBAFCPInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAFCPInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_HBAFCPInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_HBAFCPInfo) GetPropertyActive() (value bool, err error) {
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
func (instance *MSFC_HBAFCPInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_HBAFCPInfo) GetPropertyInstanceName() (value string, err error) {
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

// <param name="HbaPortWWN" type="uint8 []"></param>
// <param name="InEntryCount" type="uint32 "></param>

// <param name="Entry" type="HBAFCPScsiEntry []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutEntryCount" type="uint32 "></param>
// <param name="TotalEntryCount" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) GetFcpTargetMapping( /* IN */ HbaPortWWN []uint8,
	/* IN */ InEntryCount uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalEntryCount uint32,
	/* OUT */ OutEntryCount uint32,
	/* OUT */ Entry []HBAFCPScsiEntry) (err error) {
	_, err = instance.InvokeMethod("GetFcpTargetMapping", HbaPortWWN, InEntryCount)
	if err != nil {
		return
	}
	return

}

//

// <param name="InEntryCount" type="uint32 "></param>

// <param name="Entry" type="HBAFCPBindingEntry []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutEntryCount" type="uint32 "></param>
// <param name="TotalEntryCount" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) GetFcpPersistentBinding( /* IN */ InEntryCount uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalEntryCount uint32,
	/* OUT */ OutEntryCount uint32,
	/* OUT */ Entry []HBAFCPBindingEntry) (err error) {
	_, err = instance.InvokeMethod("GetFcpPersistentBinding", InEntryCount)
	if err != nil {
		return
	}
	return

}

//

// <param name="PortWWN" type="uint8 []"></param>

// <param name="BindType" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) GetBindingCapability( /* IN */ PortWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ BindType uint32) (err error) {
	_, err = instance.InvokeMethod("GetBindingCapability", PortWWN)
	if err != nil {
		return
	}
	return

}

//

// <param name="PortWWN" type="uint8 []"></param>

// <param name="BindType" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) GetBindingSupport( /* IN */ PortWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ BindType uint32) (err error) {
	_, err = instance.InvokeMethod("GetBindingSupport", PortWWN)
	if err != nil {
		return
	}
	return

}

//

// <param name="BindType" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) SetBindingSupport( /* IN */ PortWWN []uint8,
	/* IN */ BindType uint32,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SetBindingSupport", PortWWN, BindType)
	if err != nil {
		return
	}
	return

}

//

// <param name="InEntryCount" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="Bindings" type="HBAFCPBindingEntry2 []"></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="OutEntryCount" type="uint32 "></param>
// <param name="TotalEntryCount" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) GetPersistentBinding2( /* IN */ PortWWN []uint8,
	/* IN */ InEntryCount uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalEntryCount uint32,
	/* OUT */ OutEntryCount uint32,
	/* OUT */ Bindings []HBAFCPBindingEntry2) (err error) {
	_, err = instance.InvokeMethod("GetPersistentBinding2", PortWWN, InEntryCount)
	if err != nil {
		return
	}
	return

}

//

// <param name="Binding" type="HBAFCPBindingEntry2 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) SetPersistentEntry( /* IN */ PortWWN []uint8,
	/* IN */ Binding HBAFCPBindingEntry2,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SetPersistentEntry", PortWWN, Binding)
	if err != nil {
		return
	}
	return

}

//

// <param name="Binding" type="HBAFCPBindingEntry2 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAFCPInfo) RemovePersistentEntry( /* IN */ PortWWN []uint8,
	/* IN */ Binding HBAFCPBindingEntry2,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("RemovePersistentEntry", PortWWN, Binding)
	if err != nil {
		return
	}
	return

}
