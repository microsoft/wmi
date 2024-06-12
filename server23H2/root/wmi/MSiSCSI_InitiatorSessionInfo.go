// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MSiSCSI_InitiatorSessionInfo struct
type MSiSCSI_InitiatorSessionInfo struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Number of elements in SessionList array
	SessionCount uint32

	// Variable length array of sessions.  SessionCount specifies the number of elements in the array. NOTE: this is a variable length array.
	SessionsList []ISCSI_SessionStaticInfo

	// Id that is globally unique to each instance of each adapter. Using the address of the Adapter Extension is a good idea.
	UniqueAdapterId uint64
}

func NewMSiSCSI_InitiatorSessionInfoEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_InitiatorSessionInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorSessionInfo{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_InitiatorSessionInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_InitiatorSessionInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_InitiatorSessionInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_InitiatorSessionInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) GetPropertyInstanceName() (value string, err error) {
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

// SetSessionCount sets the value of SessionCount for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) SetPropertySessionCount(value uint32) (err error) {
	return instance.SetProperty("SessionCount", (value))
}

// GetSessionCount gets the value of SessionCount for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) GetPropertySessionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionCount")
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

// SetSessionsList sets the value of SessionsList for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) SetPropertySessionsList(value []ISCSI_SessionStaticInfo) (err error) {
	return instance.SetProperty("SessionsList", (value))
}

// GetSessionsList gets the value of SessionsList for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) GetPropertySessionsList() (value []ISCSI_SessionStaticInfo, err error) {
	retValue, err := instance.GetProperty("SessionsList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_SessionStaticInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_SessionStaticInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_SessionStaticInfo(valuetmp))
	}

	return
}

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_InitiatorSessionInfo) GetPropertyUniqueAdapterId() (value uint64, err error) {
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
