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

// MSiSCSI_NICPerformance struct
type MSiSCSI_NICPerformance struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	BytesReceived uint32

	//
	BytesTransmitted uint32

	//
	InstanceName string

	//
	PDUReceived uint32

	//
	PDUTransmitted uint32
}

func NewMSiSCSI_NICPerformanceEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_NICPerformance, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_NICPerformance{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_NICPerformanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_NICPerformance, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_NICPerformance{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyActive() (value bool, err error) {
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

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyBytesReceived(value uint32) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyBytesReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
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

// SetBytesTransmitted sets the value of BytesTransmitted for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyBytesTransmitted(value uint32) (err error) {
	return instance.SetProperty("BytesTransmitted", (value))
}

// GetBytesTransmitted gets the value of BytesTransmitted for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyBytesTransmitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesTransmitted")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyInstanceName() (value string, err error) {
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

// SetPDUReceived sets the value of PDUReceived for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyPDUReceived(value uint32) (err error) {
	return instance.SetProperty("PDUReceived", (value))
}

// GetPDUReceived gets the value of PDUReceived for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyPDUReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("PDUReceived")
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

// SetPDUTransmitted sets the value of PDUTransmitted for the instance
func (instance *MSiSCSI_NICPerformance) SetPropertyPDUTransmitted(value uint32) (err error) {
	return instance.SetProperty("PDUTransmitted", (value))
}

// GetPDUTransmitted gets the value of PDUTransmitted for the instance
func (instance *MSiSCSI_NICPerformance) GetPropertyPDUTransmitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("PDUTransmitted")
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
