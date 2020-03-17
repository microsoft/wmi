// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

// SomFilterPutStatus struct
type SomFilterPutStatus struct {
	__ExtendedStatus

	//
	RuleValidationResults []uint32
}

// SetRuleValidationResults sets the value of RuleValidationResults for the instance
func (instance *SomFilterPutStatus) SetPropertyRuleValidationResults(value []uint32) (err error) {
	return instance.SetProperty("RuleValidationResults", value)
}

// GetRuleValidationResults gets the value of RuleValidationResults for the instance
func (instance *SomFilterPutStatus) GetPropertyRuleValidationResults() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RuleValidationResults")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
