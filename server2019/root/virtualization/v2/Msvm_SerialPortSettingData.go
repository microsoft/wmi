// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SerialPortSettingData struct
type Msvm_SerialPortSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	DebuggerMode bool
}

func NewMsvm_SerialPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SerialPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_SerialPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SerialPortSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SerialPortSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDebuggerMode sets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) SetPropertyDebuggerMode(value bool) (err error) {
	return instance.SetProperty("DebuggerMode", value)
}

// GetDebuggerMode gets the value of DebuggerMode for the instance
func (instance *Msvm_SerialPortSettingData) GetPropertyDebuggerMode() (value bool, err error) {
	retValue, err := instance.GetProperty("DebuggerMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_SerialPortSettingData) GetRelatedVirtualSystemSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualSystemSettingData")
}

func (instance *Msvm_SerialPortSettingData) GetRelatedResourceAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourceAllocationSettingData")
}

func (instance *Msvm_SerialPortSettingData) GetRelatedSerialPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_SerialPort")
}

func (instance *Msvm_SerialPortSettingData) GetRelatedSerialPortSettingData() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_SerialPortSettingData")
}
