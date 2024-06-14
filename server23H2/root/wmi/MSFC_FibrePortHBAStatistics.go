// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSFC_FibrePortHBAStatistics struct
type MSFC_FibrePortHBAStatistics struct {
	*cim.WmiInstance

	//
	Active bool

	//
	HBAStatus uint32

	//
	InstanceName string

	//
	Statistics MSFC_HBAPortStatistics

	//
	UniquePortId uint64
}

func NewMSFC_FibrePortHBAStatisticsEx1(instance *cim.WmiInstance) (newInstance *MSFC_FibrePortHBAStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortHBAStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_FibrePortHBAStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_FibrePortHBAStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_FibrePortHBAStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_FibrePortHBAStatistics) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_FibrePortHBAStatistics) GetPropertyActive() (value bool, err error) {
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

// SetHBAStatus sets the value of HBAStatus for the instance
func (instance *MSFC_FibrePortHBAStatistics) SetPropertyHBAStatus(value uint32) (err error) {
	return instance.SetProperty("HBAStatus", (value))
}

// GetHBAStatus gets the value of HBAStatus for the instance
func (instance *MSFC_FibrePortHBAStatistics) GetPropertyHBAStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("HBAStatus")
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
func (instance *MSFC_FibrePortHBAStatistics) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_FibrePortHBAStatistics) GetPropertyInstanceName() (value string, err error) {
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

// SetStatistics sets the value of Statistics for the instance
func (instance *MSFC_FibrePortHBAStatistics) SetPropertyStatistics(value MSFC_HBAPortStatistics) (err error) {
	return instance.SetProperty("Statistics", (value))
}

// GetStatistics gets the value of Statistics for the instance
func (instance *MSFC_FibrePortHBAStatistics) GetPropertyStatistics() (value MSFC_HBAPortStatistics, err error) {
	retValue, err := instance.GetProperty("Statistics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFC_HBAPortStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFC_HBAPortStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFC_HBAPortStatistics(valuetmp)

	return
}

// SetUniquePortId sets the value of UniquePortId for the instance
func (instance *MSFC_FibrePortHBAStatistics) SetPropertyUniquePortId(value uint64) (err error) {
	return instance.SetProperty("UniquePortId", (value))
}

// GetUniquePortId gets the value of UniquePortId for the instance
func (instance *MSFC_FibrePortHBAStatistics) GetPropertyUniquePortId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniquePortId")
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
