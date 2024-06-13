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

// MSiSCSI_InitiatorInstanceStatistics struct
type MSiSCSI_InitiatorInstanceStatistics struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	InstanceName string

	//
	SessionConnectionTimeoutErrorCount uint32

	//
	SessionDigestErrorCount uint32

	//
	SessionFailureCount uint32

	//
	SessionFormatErrorCount uint32

	//
	UniqueAdapterId uint64
}

func NewMSiSCSI_InitiatorInstanceStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_InitiatorInstanceStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorInstanceStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_InitiatorInstanceStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_InitiatorInstanceStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorInstanceStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertyInstanceName() (value string, err error) {
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

// SetSessionConnectionTimeoutErrorCount sets the value of SessionConnectionTimeoutErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertySessionConnectionTimeoutErrorCount(value uint32) (err error) {
	return instance.SetProperty("SessionConnectionTimeoutErrorCount", (value))
}

// GetSessionConnectionTimeoutErrorCount gets the value of SessionConnectionTimeoutErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertySessionConnectionTimeoutErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionConnectionTimeoutErrorCount")
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

// SetSessionDigestErrorCount sets the value of SessionDigestErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertySessionDigestErrorCount(value uint32) (err error) {
	return instance.SetProperty("SessionDigestErrorCount", (value))
}

// GetSessionDigestErrorCount gets the value of SessionDigestErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertySessionDigestErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionDigestErrorCount")
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

// SetSessionFailureCount sets the value of SessionFailureCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertySessionFailureCount(value uint32) (err error) {
	return instance.SetProperty("SessionFailureCount", (value))
}

// GetSessionFailureCount gets the value of SessionFailureCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertySessionFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionFailureCount")
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

// SetSessionFormatErrorCount sets the value of SessionFormatErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertySessionFormatErrorCount(value uint32) (err error) {
	return instance.SetProperty("SessionFormatErrorCount", (value))
}

// GetSessionFormatErrorCount gets the value of SessionFormatErrorCount for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertySessionFormatErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionFormatErrorCount")
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorInstanceStatistics) GetPropertyUniqueAdapterId() (value uint64, err error) {
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
