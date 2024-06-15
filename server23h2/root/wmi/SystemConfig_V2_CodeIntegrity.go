// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_CodeIntegrity struct
type SystemConfig_V2_CodeIntegrity struct {
	*SystemConfig_V2

	//
	CodeIntegrityInfo uint32
}

func NewSystemConfig_V2_CodeIntegrityEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_CodeIntegrity, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_CodeIntegrity{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_CodeIntegrityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_CodeIntegrity, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_CodeIntegrity{
		SystemConfig_V2: tmp,
	}
	return
}

// SetCodeIntegrityInfo sets the value of CodeIntegrityInfo for the instance
func (instance *SystemConfig_V2_CodeIntegrity) SetPropertyCodeIntegrityInfo(value uint32) (err error) {
	return instance.SetProperty("CodeIntegrityInfo", (value))
}

// GetCodeIntegrityInfo gets the value of CodeIntegrityInfo for the instance
func (instance *SystemConfig_V2_CodeIntegrity) GetPropertyCodeIntegrityInfo() (value uint32, err error) {
	retValue, err := instance.GetProperty("CodeIntegrityInfo")
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
