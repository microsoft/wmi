// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Process_V2_TypeGroup2 struct
type Process_V2_TypeGroup2 struct {
	*Process_V2

	//
	HandleCount uint32

	//
	PageFaultCount uint32

	//
	PagefileUsage interface{}

	//
	PeakPagefileUsage interface{}

	//
	PeakVirtualSize interface{}

	//
	PeakWorkingSetSize interface{}

	//
	PrivatePageCount interface{}

	//
	ProcessId uint32

	//
	QuotaNonPagedPoolUsage interface{}

	//
	QuotaPagedPoolUsage interface{}

	//
	QuotaPeakNonPagedPoolUsage interface{}

	//
	QuotaPeakPagedPoolUsage interface{}

	//
	Reserved uint32

	//
	VirtualSize interface{}

	//
	WorkingSetSize interface{}
}

func NewProcess_V2_TypeGroup2Ex1(instance *cim.WmiInstance) (newInstance *Process_V2_TypeGroup2, err error) {
	tmp, err := NewProcess_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup2{
		Process_V2: tmp,
	}
	return
}

func NewProcess_V2_TypeGroup2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Process_V2_TypeGroup2, err error) {
	tmp, err := NewProcess_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Process_V2_TypeGroup2{
		Process_V2: tmp,
	}
	return
}

// SetHandleCount sets the value of HandleCount for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyHandleCount(value uint32) (err error) {
	return instance.SetProperty("HandleCount", (value))
}

// GetHandleCount gets the value of HandleCount for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyHandleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandleCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPageFaultCount sets the value of PageFaultCount for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPageFaultCount(value uint32) (err error) {
	return instance.SetProperty("PageFaultCount", (value))
}

// GetPageFaultCount gets the value of PageFaultCount for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPageFaultCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageFaultCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPagefileUsage sets the value of PagefileUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPagefileUsage(value interface{}) (err error) {
	return instance.SetProperty("PagefileUsage", (value))
}

// GetPagefileUsage gets the value of PagefileUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPagefileUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PagefileUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPeakPagefileUsage sets the value of PeakPagefileUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPeakPagefileUsage(value interface{}) (err error) {
	return instance.SetProperty("PeakPagefileUsage", (value))
}

// GetPeakPagefileUsage gets the value of PeakPagefileUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPeakPagefileUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PeakPagefileUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPeakVirtualSize sets the value of PeakVirtualSize for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPeakVirtualSize(value interface{}) (err error) {
	return instance.SetProperty("PeakVirtualSize", (value))
}

// GetPeakVirtualSize gets the value of PeakVirtualSize for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPeakVirtualSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PeakVirtualSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPeakWorkingSetSize sets the value of PeakWorkingSetSize for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPeakWorkingSetSize(value interface{}) (err error) {
	return instance.SetProperty("PeakWorkingSetSize", (value))
}

// GetPeakWorkingSetSize gets the value of PeakWorkingSetSize for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPeakWorkingSetSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PeakWorkingSetSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetPrivatePageCount sets the value of PrivatePageCount for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyPrivatePageCount(value interface{}) (err error) {
	return instance.SetProperty("PrivatePageCount", (value))
}

// GetPrivatePageCount gets the value of PrivatePageCount for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyPrivatePageCount() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PrivatePageCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetQuotaNonPagedPoolUsage sets the value of QuotaNonPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyQuotaNonPagedPoolUsage(value interface{}) (err error) {
	return instance.SetProperty("QuotaNonPagedPoolUsage", (value))
}

// GetQuotaNonPagedPoolUsage gets the value of QuotaNonPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyQuotaNonPagedPoolUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("QuotaNonPagedPoolUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetQuotaPagedPoolUsage sets the value of QuotaPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyQuotaPagedPoolUsage(value interface{}) (err error) {
	return instance.SetProperty("QuotaPagedPoolUsage", (value))
}

// GetQuotaPagedPoolUsage gets the value of QuotaPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyQuotaPagedPoolUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("QuotaPagedPoolUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetQuotaPeakNonPagedPoolUsage sets the value of QuotaPeakNonPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyQuotaPeakNonPagedPoolUsage(value interface{}) (err error) {
	return instance.SetProperty("QuotaPeakNonPagedPoolUsage", (value))
}

// GetQuotaPeakNonPagedPoolUsage gets the value of QuotaPeakNonPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyQuotaPeakNonPagedPoolUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("QuotaPeakNonPagedPoolUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetQuotaPeakPagedPoolUsage sets the value of QuotaPeakPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyQuotaPeakPagedPoolUsage(value interface{}) (err error) {
	return instance.SetProperty("QuotaPeakPagedPoolUsage", (value))
}

// GetQuotaPeakPagedPoolUsage gets the value of QuotaPeakPagedPoolUsage for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyQuotaPeakPagedPoolUsage() (value interface{}, err error) {
	retValue, err := instance.GetProperty("QuotaPeakPagedPoolUsage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetVirtualSize sets the value of VirtualSize for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyVirtualSize(value interface{}) (err error) {
	return instance.SetProperty("VirtualSize", (value))
}

// GetVirtualSize gets the value of VirtualSize for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyVirtualSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("VirtualSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetWorkingSetSize sets the value of WorkingSetSize for the instance
func (instance *Process_V2_TypeGroup2) SetPropertyWorkingSetSize(value interface{}) (err error) {
	return instance.SetProperty("WorkingSetSize", (value))
}

// GetWorkingSetSize gets the value of WorkingSetSize for the instance
func (instance *Process_V2_TypeGroup2) GetPropertyWorkingSetSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("WorkingSetSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
