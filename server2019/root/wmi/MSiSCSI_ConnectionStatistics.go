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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_ConnectionStatistics struct
type MSiSCSI_ConnectionStatistics struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	BytesReceived uint64

	//
	BytesSent uint64

	//
	CID uint16

	//
	InstanceName string

	// Name of the iSCSI Target
	iSCSIName string

	//
	PDUCommandsSent uint64

	//
	PDUResponsesReceived uint64

	//
	UniqueAdapterId uint64

	//
	USID uint64
}

func NewMSiSCSI_ConnectionStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_ConnectionStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_ConnectionStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_ConnectionStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_ConnectionStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_ConnectionStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", (value))
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCID sets the value of CID for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyCID(value uint16) (err error) {
	return instance.SetProperty("CID", (value))
}

// GetCID gets the value of CID for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyCID() (value uint16, err error) {
	retValue, err := instance.GetProperty("CID")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyInstanceName() (value string, err error) {
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

// SetiSCSIName sets the value of iSCSIName for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyiSCSIName(value string) (err error) {
	return instance.SetProperty("iSCSIName", (value))
}

// GetiSCSIName gets the value of iSCSIName for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyiSCSIName() (value string, err error) {
	retValue, err := instance.GetProperty("iSCSIName")
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

// SetPDUCommandsSent sets the value of PDUCommandsSent for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyPDUCommandsSent(value uint64) (err error) {
	return instance.SetProperty("PDUCommandsSent", (value))
}

// GetPDUCommandsSent gets the value of PDUCommandsSent for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyPDUCommandsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("PDUCommandsSent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPDUResponsesReceived sets the value of PDUResponsesReceived for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyPDUResponsesReceived(value uint64) (err error) {
	return instance.SetProperty("PDUResponsesReceived", (value))
}

// GetPDUResponsesReceived gets the value of PDUResponsesReceived for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyPDUResponsesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("PDUResponsesReceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyUniqueAdapterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueAdapterId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUSID sets the value of USID for the instance
func (instance *MSiSCSI_ConnectionStatistics) SetPropertyUSID(value uint64) (err error) {
	return instance.SetProperty("USID", (value))
}

// GetUSID gets the value of USID for the instance
func (instance *MSiSCSI_ConnectionStatistics) GetPropertyUSID() (value uint64, err error) {
	retValue, err := instance.GetProperty("USID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
