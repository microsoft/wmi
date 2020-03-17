// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_SubcategorySystemAuditSetting struct
type RSOP_SubcategorySystemAuditSetting struct {
	RSOP_PolicySetting

	//
	SettingValue uint32

	//
	SubcategoryGuid string

	//
	SubcategoryName string
}

// SetSettingValue sets the value of SettingValue for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySettingValue(value uint32) (err error) {
	return instance.SetProperty("SettingValue", value)
}

// GetSettingValue gets the value of SettingValue for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySettingValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("SettingValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubcategoryGuid sets the value of SubcategoryGuid for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySubcategoryGuid(value string) (err error) {
	return instance.SetProperty("SubcategoryGuid", value)
}

// GetSubcategoryGuid gets the value of SubcategoryGuid for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySubcategoryGuid() (value string, err error) {
	retValue, err := instance.GetProperty("SubcategoryGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubcategoryName sets the value of SubcategoryName for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) SetPropertySubcategoryName(value string) (err error) {
	return instance.SetProperty("SubcategoryName", value)
}

// GetSubcategoryName gets the value of SubcategoryName for the instance
func (instance *RSOP_SubcategorySystemAuditSetting) GetPropertySubcategoryName() (value string, err error) {
	retValue, err := instance.GetProperty("SubcategoryName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
