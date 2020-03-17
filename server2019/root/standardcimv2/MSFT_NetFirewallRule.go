// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetFirewallRule struct
type MSFT_NetFirewallRule struct {
	CIM_PolicyRule

	//
	Action uint16

	//
	Direction uint16

	//
	DisplayGroup string

	//
	DisplayName string

	//
	EdgeTraversalPolicy uint16

	//
	EnforcementStatus []uint16

	//
	LocalOnlyMapping bool

	//
	LooseSourceMapping bool

	//
	Owner string

	//
	Platforms []string

	//
	PolicyStoreSource string

	//
	PolicyStoreSourceType uint16

	//
	PrimaryStatus uint16

	//
	Profiles uint16

	//
	RuleGroup string

	//
	Status string

	//
	StatusCode uint32
}

// SetAction sets the value of Action for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyAction(value uint16) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirection sets the value of Direction for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyDirection(value uint16) (err error) {
	return instance.SetProperty("Direction", value)
}

// GetDirection gets the value of Direction for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyDirection() (value uint16, err error) {
	retValue, err := instance.GetProperty("Direction")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayGroup sets the value of DisplayGroup for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyDisplayGroup(value string) (err error) {
	return instance.SetProperty("DisplayGroup", value)
}

// GetDisplayGroup gets the value of DisplayGroup for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyDisplayGroup() (value string, err error) {
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
func (instance *MSFT_NetFirewallRule) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyDisplayName() (value string, err error) {
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

// SetEdgeTraversalPolicy sets the value of EdgeTraversalPolicy for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyEdgeTraversalPolicy(value uint16) (err error) {
	return instance.SetProperty("EdgeTraversalPolicy", value)
}

// GetEdgeTraversalPolicy gets the value of EdgeTraversalPolicy for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyEdgeTraversalPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("EdgeTraversalPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnforcementStatus sets the value of EnforcementStatus for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyEnforcementStatus(value []uint16) (err error) {
	return instance.SetProperty("EnforcementStatus", value)
}

// GetEnforcementStatus gets the value of EnforcementStatus for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyEnforcementStatus() (value []uint16, err error) {
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

// SetLocalOnlyMapping sets the value of LocalOnlyMapping for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyLocalOnlyMapping(value bool) (err error) {
	return instance.SetProperty("LocalOnlyMapping", value)
}

// GetLocalOnlyMapping gets the value of LocalOnlyMapping for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyLocalOnlyMapping() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalOnlyMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLooseSourceMapping sets the value of LooseSourceMapping for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyLooseSourceMapping(value bool) (err error) {
	return instance.SetProperty("LooseSourceMapping", value)
}

// GetLooseSourceMapping gets the value of LooseSourceMapping for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyLooseSourceMapping() (value bool, err error) {
	retValue, err := instance.GetProperty("LooseSourceMapping")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwner sets the value of Owner for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyOwner(value string) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyOwner() (value string, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatforms sets the value of Platforms for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyPlatforms(value []string) (err error) {
	return instance.SetProperty("Platforms", value)
}

// GetPlatforms gets the value of Platforms for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyPlatforms() (value []string, err error) {
	retValue, err := instance.GetProperty("Platforms")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStoreSource sets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyPolicyStoreSource(value string) (err error) {
	return instance.SetProperty("PolicyStoreSource", value)
}

// GetPolicyStoreSource gets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyPolicyStoreSource() (value string, err error) {
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
func (instance *MSFT_NetFirewallRule) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", value)
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
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
func (instance *MSFT_NetFirewallRule) SetPropertyPrimaryStatus(value uint16) (err error) {
	return instance.SetProperty("PrimaryStatus", value)
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyPrimaryStatus() (value uint16, err error) {
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

// SetProfiles sets the value of Profiles for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyProfiles(value uint16) (err error) {
	return instance.SetProperty("Profiles", value)
}

// GetProfiles gets the value of Profiles for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyProfiles() (value uint16, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRuleGroup sets the value of RuleGroup for the instance
func (instance *MSFT_NetFirewallRule) SetPropertyRuleGroup(value string) (err error) {
	return instance.SetProperty("RuleGroup", value)
}

// GetRuleGroup gets the value of RuleGroup for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyRuleGroup() (value string, err error) {
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
func (instance *MSFT_NetFirewallRule) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyStatus() (value string, err error) {
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
func (instance *MSFT_NetFirewallRule) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", value)
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *MSFT_NetFirewallRule) GetPropertyStatusCode() (value uint32, err error) {
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallRule) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallRule) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallRule) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallRule) CloneObject( /* IN */ NewName string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Dependents" type="CIM_ManagedSystemElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallRule) EnumerateFull( /* OUT */ Dependents []CIM_ManagedSystemElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnumerateFull")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
