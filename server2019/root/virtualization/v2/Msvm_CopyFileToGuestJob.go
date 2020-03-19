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

// Msvm_CopyFileToGuestJob struct
type Msvm_CopyFileToGuestJob struct {
	*CIM_ConcreteJob

	//
	Cancellable bool

	//
	CopyFileToGuestSettingData []string

	//
	ErrorSummaryDescription string

	//
	VirtualSystemName string
}

func NewMsvm_CopyFileToGuestJobEx1(instance *cim.WmiInstance) (newInstance *Msvm_CopyFileToGuestJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CopyFileToGuestJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMsvm_CopyFileToGuestJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CopyFileToGuestJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CopyFileToGuestJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_CopyFileToGuestJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", value)
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_CopyFileToGuestJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCopyFileToGuestSettingData sets the value of CopyFileToGuestSettingData for the instance
func (instance *Msvm_CopyFileToGuestJob) SetPropertyCopyFileToGuestSettingData(value []string) (err error) {
	return instance.SetProperty("CopyFileToGuestSettingData", value)
}

// GetCopyFileToGuestSettingData gets the value of CopyFileToGuestSettingData for the instance
func (instance *Msvm_CopyFileToGuestJob) GetPropertyCopyFileToGuestSettingData() (value []string, err error) {
	retValue, err := instance.GetProperty("CopyFileToGuestSettingData")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_CopyFileToGuestJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", value)
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_CopyFileToGuestJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemName sets the value of VirtualSystemName for the instance
func (instance *Msvm_CopyFileToGuestJob) SetPropertyVirtualSystemName(value string) (err error) {
	return instance.SetProperty("VirtualSystemName", value)
}

// GetVirtualSystemName gets the value of VirtualSystemName for the instance
func (instance *Msvm_CopyFileToGuestJob) GetPropertyVirtualSystemName() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_CopyFileToGuestJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
