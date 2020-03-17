// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_CentralAccessPolicySetting struct
type RSOP_CentralAccessPolicySetting struct {
	RSOP_PolicySetting

	//
	CentralAccessPolicyName []string
}

// SetCentralAccessPolicyName sets the value of CentralAccessPolicyName for the instance
func (instance *RSOP_CentralAccessPolicySetting) SetPropertyCentralAccessPolicyName(value []string) (err error) {
	return instance.SetProperty("CentralAccessPolicyName", value)
}

// GetCentralAccessPolicyName gets the value of CentralAccessPolicyName for the instance
func (instance *RSOP_CentralAccessPolicySetting) GetPropertyCentralAccessPolicyName() (value []string, err error) {
	retValue, err := instance.GetProperty("CentralAccessPolicyName")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
