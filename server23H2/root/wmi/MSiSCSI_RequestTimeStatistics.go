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

// MSiSCSI_RequestTimeStatistics struct
type MSiSCSI_RequestTimeStatistics struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	AverageProcessingTime uint32

	//
	CID uint16

	//
	InstanceName string

	// Name of the iSCSI Target
	iSCSIName string

	//
	MaximumProcessingTime uint32

	//
	UniqueAdapterId uint64

	//
	USID uint64
}

func NewMSiSCSI_RequestTimeStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_RequestTimeStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RequestTimeStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_RequestTimeStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_RequestTimeStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RequestTimeStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyActive() (value bool, err error) {
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

// SetAverageProcessingTime sets the value of AverageProcessingTime for the instance
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyAverageProcessingTime(value uint32) (err error) {
	return instance.SetProperty("AverageProcessingTime", (value))
}

// GetAverageProcessingTime gets the value of AverageProcessingTime for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyAverageProcessingTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("AverageProcessingTime")
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

// SetCID sets the value of CID for the instance
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyCID(value uint16) (err error) {
	return instance.SetProperty("CID", (value))
}

// GetCID gets the value of CID for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyCID() (value uint16, err error) {
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
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyInstanceName() (value string, err error) {
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
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyiSCSIName(value string) (err error) {
	return instance.SetProperty("iSCSIName", (value))
}

// GetiSCSIName gets the value of iSCSIName for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyiSCSIName() (value string, err error) {
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

// SetMaximumProcessingTime sets the value of MaximumProcessingTime for the instance
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyMaximumProcessingTime(value uint32) (err error) {
	return instance.SetProperty("MaximumProcessingTime", (value))
}

// GetMaximumProcessingTime gets the value of MaximumProcessingTime for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyMaximumProcessingTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumProcessingTime")
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
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyUniqueAdapterId() (value uint64, err error) {
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
func (instance *MSiSCSI_RequestTimeStatistics) SetPropertyUSID(value uint64) (err error) {
	return instance.SetProperty("USID", (value))
}

// GetUSID gets the value of USID for the instance
func (instance *MSiSCSI_RequestTimeStatistics) GetPropertyUSID() (value uint64, err error) {
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
