// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_PolicyAction struct
type CIM_PolicyAction struct {
	*CIM_Policy

	//
	CreationClassName string

	//
	DoActionLogging bool

	//
	PolicyActionName string

	//
	PolicyRuleCreationClassName string

	//
	PolicyRuleName string

	//
	SystemCreationClassName string

	//
	SystemName string
}

func NewCIM_PolicyActionEx1(instance *cim.WmiInstance) (newInstance *CIM_PolicyAction, err error) {
	tmp, err := NewCIM_PolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyAction{
		CIM_Policy: tmp,
	}
	return
}

func NewCIM_PolicyActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PolicyAction, err error) {
	tmp, err := NewCIM_PolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicyAction{
		CIM_Policy: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetDoActionLogging sets the value of DoActionLogging for the instance
func (instance *CIM_PolicyAction) SetPropertyDoActionLogging(value bool) (err error) {
	return instance.SetProperty("DoActionLogging", (value))
}

// GetDoActionLogging gets the value of DoActionLogging for the instance
func (instance *CIM_PolicyAction) GetPropertyDoActionLogging() (value bool, err error) {
	retValue, err := instance.GetProperty("DoActionLogging")
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

// SetPolicyActionName sets the value of PolicyActionName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyActionName(value string) (err error) {
	return instance.SetProperty("PolicyActionName", (value))
}

// GetPolicyActionName gets the value of PolicyActionName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyActionName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyActionName")
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

// SetPolicyRuleCreationClassName sets the value of PolicyRuleCreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyRuleCreationClassName(value string) (err error) {
	return instance.SetProperty("PolicyRuleCreationClassName", (value))
}

// GetPolicyRuleCreationClassName gets the value of PolicyRuleCreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyRuleCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyRuleCreationClassName")
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

// SetPolicyRuleName sets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyRuleName(value string) (err error) {
	return instance.SetProperty("PolicyRuleName", (value))
}

// GetPolicyRuleName gets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyRuleName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyRuleName")
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

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_PolicyAction) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_PolicyAction) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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
