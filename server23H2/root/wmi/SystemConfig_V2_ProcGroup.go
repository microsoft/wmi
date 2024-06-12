// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_ProcGroup struct
type SystemConfig_V2_ProcGroup struct {
	*SystemConfig_V2

	//
	Affinity []uint32

	//
	GroupCount uint32
}

func NewSystemConfig_V2_ProcGroupEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_ProcGroup, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_ProcGroup{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_ProcGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_ProcGroup, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_ProcGroup{
		SystemConfig_V2: tmp,
	}
	return
}

// SetAffinity sets the value of Affinity for the instance
func (instance *SystemConfig_V2_ProcGroup) SetPropertyAffinity(value []uint32) (err error) {
	return instance.SetProperty("Affinity", (value))
}

// GetAffinity gets the value of Affinity for the instance
func (instance *SystemConfig_V2_ProcGroup) GetPropertyAffinity() (value []uint32, err error) {
	retValue, err := instance.GetProperty("Affinity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetGroupCount sets the value of GroupCount for the instance
func (instance *SystemConfig_V2_ProcGroup) SetPropertyGroupCount(value uint32) (err error) {
	return instance.SetProperty("GroupCount", (value))
}

// GetGroupCount gets the value of GroupCount for the instance
func (instance *SystemConfig_V2_ProcGroup) GetPropertyGroupCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupCount")
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
