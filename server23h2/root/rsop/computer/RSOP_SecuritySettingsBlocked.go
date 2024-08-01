// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_SecuritySettingsBlocked struct
type RSOP_SecuritySettingsBlocked struct {
	*RSOP_PolicySetting

	//
	ErrorCode uint32

	//
	Status uint32
}

func NewRSOP_SecuritySettingsBlockedEx1(instance *cim.WmiInstance) (newInstance *RSOP_SecuritySettingsBlocked, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingsBlocked{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_SecuritySettingsBlockedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SecuritySettingsBlocked, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SecuritySettingsBlocked{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *RSOP_SecuritySettingsBlocked) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", (value))
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *RSOP_SecuritySettingsBlocked) GetPropertyErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
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

// SetStatus sets the value of Status for the instance
func (instance *RSOP_SecuritySettingsBlocked) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *RSOP_SecuritySettingsBlocked) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
