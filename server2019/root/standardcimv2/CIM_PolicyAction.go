// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_PolicyAction struct
type CIM_PolicyAction struct {
	CIM_Policy

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

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDoActionLogging sets the value of DoActionLogging for the instance
func (instance *CIM_PolicyAction) SetPropertyDoActionLogging(value bool) (err error) {
	return instance.SetProperty("DoActionLogging", value)
}

// GetDoActionLogging gets the value of DoActionLogging for the instance
func (instance *CIM_PolicyAction) GetPropertyDoActionLogging() (value bool, err error) {
	retValue, err := instance.GetProperty("DoActionLogging")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyActionName sets the value of PolicyActionName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyActionName(value string) (err error) {
	return instance.SetProperty("PolicyActionName", value)
}

// GetPolicyActionName gets the value of PolicyActionName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyActionName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyActionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyRuleCreationClassName sets the value of PolicyRuleCreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyRuleCreationClassName(value string) (err error) {
	return instance.SetProperty("PolicyRuleCreationClassName", value)
}

// GetPolicyRuleCreationClassName gets the value of PolicyRuleCreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyRuleCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyRuleCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyRuleName sets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyAction) SetPropertyPolicyRuleName(value string) (err error) {
	return instance.SetProperty("PolicyRuleName", value)
}

// GetPolicyRuleName gets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyAction) GetPropertyPolicyRuleName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyRuleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyAction) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyAction) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_PolicyAction) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_PolicyAction) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
