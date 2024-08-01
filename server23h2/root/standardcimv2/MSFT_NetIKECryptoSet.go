// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIKECryptoSet struct
type MSFT_NetIKECryptoSet struct {
	*CIM_IKEAction

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
	Proposals []MSFT_NetIKECryptoProposal

	//
	RuleGroup string

	//
	Status string

	//
	StatusCode uint32
}

func NewMSFT_NetIKECryptoSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKECryptoSet, err error) {
	tmp, err := NewCIM_IKEActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECryptoSet{
		CIM_IKEAction: tmp,
	}
	return
}

func NewMSFT_NetIKECryptoSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKECryptoSet, err error) {
	tmp, err := NewCIM_IKEActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECryptoSet{
		CIM_IKEAction: tmp,
	}
	return
}

// SetDisplayGroup sets the value of DisplayGroup for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyDisplayGroup(value string) (err error) {
	return instance.SetProperty("DisplayGroup", (value))
}

// GetDisplayGroup gets the value of DisplayGroup for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyDisplayGroup() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayGroup")
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

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyDisplayName() (value string, err error) {
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
func (instance *MSFT_NetIKECryptoSet) SetPropertyEnforcementStatus(value []uint16) (err error) {
	return instance.SetProperty("EnforcementStatus", (value))
}

// GetEnforcementStatus gets the value of EnforcementStatus for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyEnforcementStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("EnforcementStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetPolicyStoreSource sets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyPolicyStoreSource(value string) (err error) {
	return instance.SetProperty("PolicyStoreSource", (value))
}

// GetPolicyStoreSource gets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyPolicyStoreSource() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStoreSource")
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

// SetPolicyStoreSourceType sets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", (value))
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
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

// SetPrimaryStatus sets the value of PrimaryStatus for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyPrimaryStatus(value uint16) (err error) {
	return instance.SetProperty("PrimaryStatus", (value))
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyPrimaryStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("PrimaryStatus")
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

// SetProposals sets the value of Proposals for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyProposals(value []MSFT_NetIKECryptoProposal) (err error) {
	return instance.SetProperty("Proposals", (value))
}

// GetProposals gets the value of Proposals for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyProposals() (value []MSFT_NetIKECryptoProposal, err error) {
	retValue, err := instance.GetProperty("Proposals")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetIKECryptoProposal)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetIKECryptoProposal is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetIKECryptoProposal(valuetmp))
	}

	return
}

// SetRuleGroup sets the value of RuleGroup for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyRuleGroup(value string) (err error) {
	return instance.SetProperty("RuleGroup", (value))
}

// GetRuleGroup gets the value of RuleGroup for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyRuleGroup() (value string, err error) {
	retValue, err := instance.GetProperty("RuleGroup")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStatusCode sets the value of StatusCode for the instance
func (instance *MSFT_NetIKECryptoSet) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", (value))
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *MSFT_NetIKECryptoSet) GetPropertyStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
