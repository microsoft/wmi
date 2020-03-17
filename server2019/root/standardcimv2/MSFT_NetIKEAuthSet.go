// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKEAuthSet struct
type MSFT_NetIKEAuthSet struct {
	CIM_IKEAction

	//
	DisplayGroup string

	//
	DisplayName string

	//
	EnforcementStatus []uint16

	//
	PolicyStoreSource string

	//
	PolicyStoreSourceType uint16

	//
	PrimaryStatus uint16

	//
	Proposals []MSFT_NetIKEAuthProposal

	//
	RuleGroup string

	//
	Status string

	//
	StatusCode uint32
}

// SetDisplayGroup sets the value of DisplayGroup for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyDisplayGroup(value string) (err error) {
	return instance.SetProperty("DisplayGroup", value)
}

// GetDisplayGroup gets the value of DisplayGroup for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyDisplayGroup() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnforcementStatus sets the value of EnforcementStatus for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyEnforcementStatus(value []uint16) (err error) {
	return instance.SetProperty("EnforcementStatus", value)
}

// GetEnforcementStatus gets the value of EnforcementStatus for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyEnforcementStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("EnforcementStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStoreSource sets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyPolicyStoreSource(value string) (err error) {
	return instance.SetProperty("PolicyStoreSource", value)
}

// GetPolicyStoreSource gets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyPolicyStoreSource() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSource")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStoreSourceType sets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", value)
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryStatus sets the value of PrimaryStatus for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyPrimaryStatus(value uint16) (err error) {
	return instance.SetProperty("PrimaryStatus", value)
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyPrimaryStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("PrimaryStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProposals sets the value of Proposals for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyProposals(value []MSFT_NetIKEAuthProposal) (err error) {
	return instance.SetProperty("Proposals", value)
}

// GetProposals gets the value of Proposals for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyProposals() (value []MSFT_NetIKEAuthProposal, err error) {
	retValue, err := instance.GetProperty("Proposals")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_NetIKEAuthProposal)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRuleGroup sets the value of RuleGroup for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyRuleGroup(value string) (err error) {
	return instance.SetProperty("RuleGroup", value)
}

// GetRuleGroup gets the value of RuleGroup for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyRuleGroup() (value string, err error) {
	retValue, err := instance.GetProperty("RuleGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusCode sets the value of StatusCode for the instance
func (instance *MSFT_NetIKEAuthSet) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", value)
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *MSFT_NetIKEAuthSet) GetPropertyStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
