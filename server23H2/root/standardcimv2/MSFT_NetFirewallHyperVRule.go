// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetFirewallHyperVRule struct
type MSFT_NetFirewallHyperVRule struct {
	*CIM_PolicyRule

	//
	Action uint16

	//
	Direction uint16

	//
	DisplayName string

	//
	EnforcementStatus uint16

	//
	LocalAddresses []string

	//
	LocalPorts []string

	//
	PolicyStoreSourceType uint16

	//
	PortStatuses []MSFT_NetFirewallHyperVRulePortStatus

	//
	Profiles uint16

	//
	Protocol string

	//
	RemoteAddresses []string

	//
	RemotePorts []string

	//
	RulePriority uint16

	//
	VMCreatorId string
}

func NewMSFT_NetFirewallHyperVRuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetFirewallHyperVRule, err error) {
	tmp, err := NewCIM_PolicyRuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVRule{
		CIM_PolicyRule: tmp,
	}
	return
}

func NewMSFT_NetFirewallHyperVRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetFirewallHyperVRule, err error) {
	tmp, err := NewCIM_PolicyRuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetFirewallHyperVRule{
		CIM_PolicyRule: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyAction(value uint16) (err error) {
	return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDirection sets the value of Direction for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyDirection(value uint16) (err error) {
	return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyDirection() (value uint16, err error) {
	retValue, err := instance.GetProperty("Direction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetEnforcementStatus sets the value of EnforcementStatus for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyEnforcementStatus(value uint16) (err error) {
	return instance.SetProperty("EnforcementStatus", (value))
}

// GetEnforcementStatus gets the value of EnforcementStatus for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyEnforcementStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnforcementStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLocalAddresses sets the value of LocalAddresses for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyLocalAddresses(value []string) (err error) {
	return instance.SetProperty("LocalAddresses", (value))
}

// GetLocalAddresses gets the value of LocalAddresses for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyLocalAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetLocalPorts sets the value of LocalPorts for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyLocalPorts(value []string) (err error) {
	return instance.SetProperty("LocalPorts", (value))
}

// GetLocalPorts gets the value of LocalPorts for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyLocalPorts() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalPorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetPolicyStoreSourceType sets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", (value))
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSourceType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPortStatuses sets the value of PortStatuses for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyPortStatuses(value []MSFT_NetFirewallHyperVRulePortStatus) (err error) {
	return instance.SetProperty("PortStatuses", (value))
}

// GetPortStatuses gets the value of PortStatuses for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyPortStatuses() (value []MSFT_NetFirewallHyperVRulePortStatus, err error) {
	retValue, err := instance.GetProperty("PortStatuses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetFirewallHyperVRulePortStatus)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetFirewallHyperVRulePortStatus is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetFirewallHyperVRulePortStatus(valuetmp))
	}

	return
}

// SetProfiles sets the value of Profiles for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyProfiles(value uint16) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyProfiles() (value uint16, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemoteAddresses sets the value of RemoteAddresses for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyRemoteAddresses(value []string) (err error) {
	return instance.SetProperty("RemoteAddresses", (value))
}

// GetRemoteAddresses gets the value of RemoteAddresses for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyRemoteAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoteAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetRemotePorts sets the value of RemotePorts for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyRemotePorts(value []string) (err error) {
	return instance.SetProperty("RemotePorts", (value))
}

// GetRemotePorts gets the value of RemotePorts for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyRemotePorts() (value []string, err error) {
	retValue, err := instance.GetProperty("RemotePorts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetRulePriority sets the value of RulePriority for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyRulePriority(value uint16) (err error) {
	return instance.SetProperty("RulePriority", (value))
}

// GetRulePriority gets the value of RulePriority for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyRulePriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("RulePriority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetVMCreatorId sets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVRule) SetPropertyVMCreatorId(value string) (err error) {
	return instance.SetProperty("VMCreatorId", (value))
}

// GetVMCreatorId gets the value of VMCreatorId for the instance
func (instance *MSFT_NetFirewallHyperVRule) GetPropertyVMCreatorId() (value string, err error) {
	retValue, err := instance.GetProperty("VMCreatorId")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallHyperVRule) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallHyperVRule) Disable() (result uint32, err error) {
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
func (instance *MSFT_NetFirewallHyperVRule) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Dependents" type="CIM_ManagedSystemElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetFirewallHyperVRule) EnumerateFull( /* OUT */ Dependents []CIM_ManagedSystemElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnumerateFull")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
