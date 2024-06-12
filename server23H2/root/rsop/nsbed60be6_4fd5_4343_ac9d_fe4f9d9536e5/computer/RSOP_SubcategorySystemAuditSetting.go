// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSBED60BE6_4FD5_4343_AC9D_FE4F9D9536E5.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_SubcategorySystemAuditSetting struct
type RSOP_SubcategorySystemAuditSetting struct {
	*RSOP_PolicySetting

	//
	SettingValue uint32

	//
	SubcategoryGuid string

	//
	SubcategoryName string
}

func NewRSOP_SubcategorySystemAuditSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_SubcategorySystemAuditSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SubcategorySystemAuditSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_SubcategorySystemAuditSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SubcategorySystemAuditSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SubcategorySystemAuditSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySettingValue(value uint32) (err error) {
	return instance.SetProperty("SettingValue", (value))
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySettingValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingValue")
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

// SetSubcategoryGuid sets the value of SubcategoryGuid for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySubcategoryGuid(value string) (err error) {
	return instance.SetProperty("SubcategoryGuid", (value))
}

// GetSubcategoryGuid gets the value of SubcategoryGuid for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySubcategoryGuid() (value string, err error) {
	retValue, err := instance.GetProperty("SubcategoryGuid")
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

// SetSubcategoryName sets the value of SubcategoryName for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySubcategoryName(value string) (err error) {
	return instance.SetProperty("SubcategoryName", (value))
}

// GetSubcategoryName gets the value of SubcategoryName for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySubcategoryName() (value string, err error) {
	retValue, err := instance.GetProperty("SubcategoryName")
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
