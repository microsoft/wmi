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

// MSiSCSI_SessionStatistics struct
type MSiSCSI_SessionStatistics struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	BytesReceived uint64

	//
	BytesSent uint64

	//
	ConnectionTimeoutErrors uint64

	//
	DigestErrors uint64

	//
	FormatErrors uint64

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

func NewMSiSCSI_SessionStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_SessionStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_SessionStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_SessionStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_SessionStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_SessionStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_SessionStatistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", (value))
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyBytesReceived() (value uint64, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", (value))
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyBytesSent() (value uint64, err error) {
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

// SetConnectionTimeoutErrors sets the value of ConnectionTimeoutErrors for the instance
func (instance *MSiSCSI_SessionStatistics) SetPropertyConnectionTimeoutErrors(value uint64) (err error) {
	return instance.SetProperty("ConnectionTimeoutErrors", (value))
}

// GetConnectionTimeoutErrors gets the value of ConnectionTimeoutErrors for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyConnectionTimeoutErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConnectionTimeoutErrors")
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

// SetDigestErrors sets the value of DigestErrors for the instance
func (instance *MSiSCSI_SessionStatistics) SetPropertyDigestErrors(value uint64) (err error) {
	return instance.SetProperty("DigestErrors", (value))
}

// GetDigestErrors gets the value of DigestErrors for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyDigestErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("DigestErrors")
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

// SetFormatErrors sets the value of FormatErrors for the instance
func (instance *MSiSCSI_SessionStatistics) SetPropertyFormatErrors(value uint64) (err error) {
	return instance.SetProperty("FormatErrors", (value))
}

// GetFormatErrors gets the value of FormatErrors for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyFormatErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("FormatErrors")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_SessionStatistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyiSCSIName(value string) (err error) {
	return instance.SetProperty("iSCSIName", (value))
}

// GetiSCSIName gets the value of iSCSIName for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyiSCSIName() (value string, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyPDUCommandsSent(value uint64) (err error) {
	return instance.SetProperty("PDUCommandsSent", (value))
}

// GetPDUCommandsSent gets the value of PDUCommandsSent for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyPDUCommandsSent() (value uint64, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyPDUResponsesReceived(value uint64) (err error) {
	return instance.SetProperty("PDUResponsesReceived", (value))
}

// GetPDUResponsesReceived gets the value of PDUResponsesReceived for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyPDUResponsesReceived() (value uint64, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyUniqueAdapterId() (value uint64, err error) {
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
func (instance *MSiSCSI_SessionStatistics) SetPropertyUSID(value uint64) (err error) {
	return instance.SetProperty("USID", (value))
}

// GetUSID gets the value of USID for the instance
func (instance *MSiSCSI_SessionStatistics) GetPropertyUSID() (value uint64, err error) {
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
