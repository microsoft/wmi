// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MSiSCSIInitiator_LUNList struct
type MSiSCSIInitiator_LUNList struct {
	*cim.WmiInstance

	//
	OSLunNumber uint32

	//
	TargetLun uint64
}

func NewMSiSCSIInitiator_LUNListEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_LUNList, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_LUNList{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_LUNListEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_LUNList, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_LUNList{
		WmiInstance: tmp,
	}
	return
}

// SetOSLunNumber sets the value of OSLunNumber for the instance
func (instance *MSiSCSIInitiator_LUNList) SetPropertyOSLunNumber(value uint32) (err error) {
	return instance.SetProperty("OSLunNumber", (value))
}

// GetOSLunNumber gets the value of OSLunNumber for the instance
func (instance *MSiSCSIInitiator_LUNList) GetPropertyOSLunNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSLunNumber")
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

// SetTargetLun sets the value of TargetLun for the instance
func (instance *MSiSCSIInitiator_LUNList) SetPropertyTargetLun(value uint64) (err error) {
	return instance.SetProperty("TargetLun", (value))
}

// GetTargetLun gets the value of TargetLun for the instance
func (instance *MSiSCSIInitiator_LUNList) GetPropertyTargetLun() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetLun")
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
