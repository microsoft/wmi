// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_ScriptPolicySetting struct
type RSOP_ScriptPolicySetting struct {
	*RSOP_PolicySetting

	//
	psScriptOrder uint32

	//
	scriptList []RSOP_ScriptCmd

	//
	scriptOrder uint32

	//
	scriptType uint32
}

func NewRSOP_ScriptPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_ScriptPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_ScriptPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_ScriptPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ScriptPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ScriptPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetpsScriptOrder sets the value of psScriptOrder for the instance
func (instance *RSOP_ScriptPolicySetting) SetPropertypsScriptOrder(value uint32) (err error) {
	return instance.SetProperty("psScriptOrder", value)
}

// GetpsScriptOrder gets the value of psScriptOrder for the instance
func (instance *RSOP_ScriptPolicySetting) GetPropertypsScriptOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("psScriptOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetscriptList sets the value of scriptList for the instance
func (instance *RSOP_ScriptPolicySetting) SetPropertyscriptList(value []RSOP_ScriptCmd) (err error) {
	return instance.SetProperty("scriptList", value)
}

// GetscriptList gets the value of scriptList for the instance
func (instance *RSOP_ScriptPolicySetting) GetPropertyscriptList() (value []RSOP_ScriptCmd, err error) {
	retValue, err := instance.GetProperty("scriptList")
	if err != nil {
		return
	}
	value, ok := retValue.([]RSOP_ScriptCmd)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetscriptOrder sets the value of scriptOrder for the instance
func (instance *RSOP_ScriptPolicySetting) SetPropertyscriptOrder(value uint32) (err error) {
	return instance.SetProperty("scriptOrder", value)
}

// GetscriptOrder gets the value of scriptOrder for the instance
func (instance *RSOP_ScriptPolicySetting) GetPropertyscriptOrder() (value uint32, err error) {
	retValue, err := instance.GetProperty("scriptOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetscriptType sets the value of scriptType for the instance
func (instance *RSOP_ScriptPolicySetting) SetPropertyscriptType(value uint32) (err error) {
	return instance.SetProperty("scriptType", value)
}

// GetscriptType gets the value of scriptType for the instance
func (instance *RSOP_ScriptPolicySetting) GetPropertyscriptType() (value uint32, err error) {
	retValue, err := instance.GetProperty("scriptType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
