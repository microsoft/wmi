// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_PolicyRule struct
type CIM_PolicyRule struct {
	CIM_PolicySet

	//
	ConditionListType uint16

	//
	CreationClassName string

	//
	ExecutionStrategy uint16

	//
	Mandatory bool

	//
	PolicyRuleName string

	//
	Priority uint16

	//
	RuleUsage string

	//
	SequencedActions uint16

	//
	SystemCreationClassName string

	//
	SystemName string
}

// SetConditionListType sets the value of ConditionListType for the instance
func (instance *CIM_PolicyRule) SetPropertyConditionListType(value uint16) (err error) {
	return instance.SetProperty("ConditionListType", value)
}

// GetConditionListType gets the value of ConditionListType for the instance
func (instance *CIM_PolicyRule) GetPropertyConditionListType() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConditionListType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_PolicyRule) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_PolicyRule) GetPropertyCreationClassName() (value string, err error) {
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

// SetExecutionStrategy sets the value of ExecutionStrategy for the instance
func (instance *CIM_PolicyRule) SetPropertyExecutionStrategy(value uint16) (err error) {
	return instance.SetProperty("ExecutionStrategy", value)
}

// GetExecutionStrategy gets the value of ExecutionStrategy for the instance
func (instance *CIM_PolicyRule) GetPropertyExecutionStrategy() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExecutionStrategy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMandatory sets the value of Mandatory for the instance
func (instance *CIM_PolicyRule) SetPropertyMandatory(value bool) (err error) {
	return instance.SetProperty("Mandatory", value)
}

// GetMandatory gets the value of Mandatory for the instance
func (instance *CIM_PolicyRule) GetPropertyMandatory() (value bool, err error) {
	retValue, err := instance.GetProperty("Mandatory")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyRuleName sets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyRule) SetPropertyPolicyRuleName(value string) (err error) {
	return instance.SetProperty("PolicyRuleName", value)
}

// GetPolicyRuleName gets the value of PolicyRuleName for the instance
func (instance *CIM_PolicyRule) GetPropertyPolicyRuleName() (value string, err error) {
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

// SetPriority sets the value of Priority for the instance
func (instance *CIM_PolicyRule) SetPropertyPriority(value uint16) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *CIM_PolicyRule) GetPropertyPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRuleUsage sets the value of RuleUsage for the instance
func (instance *CIM_PolicyRule) SetPropertyRuleUsage(value string) (err error) {
	return instance.SetProperty("RuleUsage", value)
}

// GetRuleUsage gets the value of RuleUsage for the instance
func (instance *CIM_PolicyRule) GetPropertyRuleUsage() (value string, err error) {
	retValue, err := instance.GetProperty("RuleUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequencedActions sets the value of SequencedActions for the instance
func (instance *CIM_PolicyRule) SetPropertySequencedActions(value uint16) (err error) {
	return instance.SetProperty("SequencedActions", value)
}

// GetSequencedActions gets the value of SequencedActions for the instance
func (instance *CIM_PolicyRule) GetPropertySequencedActions() (value uint16, err error) {
	retValue, err := instance.GetProperty("SequencedActions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyRule) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_PolicyRule) GetPropertySystemCreationClassName() (value string, err error) {
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
func (instance *CIM_PolicyRule) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_PolicyRule) GetPropertySystemName() (value string, err error) {
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
