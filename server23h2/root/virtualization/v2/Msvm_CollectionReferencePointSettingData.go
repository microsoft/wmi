// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CollectionReferencePointSettingData struct
type Msvm_CollectionReferencePointSettingData struct {
	*CIM_SettingData

	//
	ConsistencyLevel uint8
}

func NewMsvm_CollectionReferencePointSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReferencePointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_CollectionReferencePointSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionReferencePointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionReferencePointSettingData) SetPropertyConsistencyLevel(value uint8) (err error) {
	return instance.SetProperty("ConsistencyLevel", (value))
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionReferencePointSettingData) GetPropertyConsistencyLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
