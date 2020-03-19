// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_Job struct
type SDDC_Job struct {
	*cim.WmiInstance

	//
	BytesProcessed uint64

	//
	BytesTotal uint64

	//
	Description string

	//
	ElapsedTime string

	//
	ErrorCode uint16

	//
	ErrorDescription string

	//
	Id string

	//
	PercentComplete uint32

	//
	StartTime string

	//
	State uint16
}

func NewSDDC_JobEx1(instance *cim.WmiInstance) (newInstance *SDDC_Job, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Job{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_JobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Job, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Job{
		WmiInstance: tmp,
	}
	return
}

// SetBytesProcessed sets the value of BytesProcessed for the instance
func (instance *SDDC_Job) SetPropertyBytesProcessed(value uint64) (err error) {
	return instance.SetProperty("BytesProcessed", value)
}

// GetBytesProcessed gets the value of BytesProcessed for the instance
func (instance *SDDC_Job) GetPropertyBytesProcessed() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesProcessed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTotal sets the value of BytesTotal for the instance
func (instance *SDDC_Job) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", value)
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *SDDC_Job) GetPropertyBytesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *SDDC_Job) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *SDDC_Job) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *SDDC_Job) SetPropertyElapsedTime(value string) (err error) {
	return instance.SetProperty("ElapsedTime", value)
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *SDDC_Job) GetPropertyElapsedTime() (value string, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *SDDC_Job) SetPropertyErrorCode(value uint16) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *SDDC_Job) GetPropertyErrorCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *SDDC_Job) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *SDDC_Job) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_Job) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Job) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *SDDC_Job) SetPropertyPercentComplete(value uint32) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *SDDC_Job) GetPropertyPercentComplete() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *SDDC_Job) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *SDDC_Job) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *SDDC_Job) SetPropertyState(value uint16) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *SDDC_Job) GetPropertyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
