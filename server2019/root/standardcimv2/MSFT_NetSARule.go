// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetSARule struct
type MSFT_NetSARule struct {
	*CIM_SARule

	//
	DisplayGroup string

	//
	DisplayName string

	//
	EnforcementStatus []uint16

	//
	MainModeCryptoSet string

	//
	Phase1AuthSet string

	//
	Phase2AuthSet string

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
	QuickModeCryptoSet string

	//
	RuleGroup string

	//
	Status string

	//
	StatusCode uint32
}

func NewMSFT_NetSARuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSARule, err error) {
	tmp, err := NewCIM_SARuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSARule{
		CIM_SARule: tmp,
	}
	return
}

func NewMSFT_NetSARuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetSARule, err error) {
	tmp, err := NewCIM_SARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetSARule{
		CIM_SARule: tmp,
	}
	return
}

// SetDisplayGroup sets the value of DisplayGroup for the instance
func (instance *MSFT_NetSARule) SetPropertyDisplayGroup(value string) (err error) {
	return instance.SetProperty("DisplayGroup", (value))
}

// GetDisplayGroup gets the value of DisplayGroup for the instance
func (instance *MSFT_NetSARule) GetPropertyDisplayGroup() (value string, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetSARule) GetPropertyDisplayName() (value string, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyEnforcementStatus(value []uint16) (err error) {
	return instance.SetProperty("EnforcementStatus", (value))
}

// GetEnforcementStatus gets the value of EnforcementStatus for the instance
func (instance *MSFT_NetSARule) GetPropertyEnforcementStatus() (value []uint16, err error) {
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

// SetMainModeCryptoSet sets the value of MainModeCryptoSet for the instance
func (instance *MSFT_NetSARule) SetPropertyMainModeCryptoSet(value string) (err error) {
	return instance.SetProperty("MainModeCryptoSet", (value))
}

// GetMainModeCryptoSet gets the value of MainModeCryptoSet for the instance
func (instance *MSFT_NetSARule) GetPropertyMainModeCryptoSet() (value string, err error) {
	retValue, err := instance.GetProperty("MainModeCryptoSet")
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

// SetPhase1AuthSet sets the value of Phase1AuthSet for the instance
func (instance *MSFT_NetSARule) SetPropertyPhase1AuthSet(value string) (err error) {
	return instance.SetProperty("Phase1AuthSet", (value))
}

// GetPhase1AuthSet gets the value of Phase1AuthSet for the instance
func (instance *MSFT_NetSARule) GetPropertyPhase1AuthSet() (value string, err error) {
	retValue, err := instance.GetProperty("Phase1AuthSet")
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

// SetPhase2AuthSet sets the value of Phase2AuthSet for the instance
func (instance *MSFT_NetSARule) SetPropertyPhase2AuthSet(value string) (err error) {
	return instance.SetProperty("Phase2AuthSet", (value))
}

// GetPhase2AuthSet gets the value of Phase2AuthSet for the instance
func (instance *MSFT_NetSARule) GetPropertyPhase2AuthSet() (value string, err error) {
	retValue, err := instance.GetProperty("Phase2AuthSet")
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

// SetPlatforms sets the value of Platforms for the instance
func (instance *MSFT_NetSARule) SetPropertyPlatforms(value []string) (err error) {
	return instance.SetProperty("Platforms", (value))
}

// GetPlatforms gets the value of Platforms for the instance
func (instance *MSFT_NetSARule) GetPropertyPlatforms() (value []string, err error) {
	retValue, err := instance.GetProperty("Platforms")
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

// SetPolicyStoreSource sets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetSARule) SetPropertyPolicyStoreSource(value string) (err error) {
	return instance.SetProperty("PolicyStoreSource", (value))
}

// GetPolicyStoreSource gets the value of PolicyStoreSource for the instance
func (instance *MSFT_NetSARule) GetPropertyPolicyStoreSource() (value string, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyPolicyStoreSourceType(value uint16) (err error) {
	return instance.SetProperty("PolicyStoreSourceType", (value))
}

// GetPolicyStoreSourceType gets the value of PolicyStoreSourceType for the instance
func (instance *MSFT_NetSARule) GetPropertyPolicyStoreSourceType() (value uint16, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyPrimaryStatus(value uint16) (err error) {
	return instance.SetProperty("PrimaryStatus", (value))
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *MSFT_NetSARule) GetPropertyPrimaryStatus() (value uint16, err error) {
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

// SetProfiles sets the value of Profiles for the instance
func (instance *MSFT_NetSARule) SetPropertyProfiles(value uint16) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *MSFT_NetSARule) GetPropertyProfiles() (value uint16, err error) {
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

// SetQuickModeCryptoSet sets the value of QuickModeCryptoSet for the instance
func (instance *MSFT_NetSARule) SetPropertyQuickModeCryptoSet(value string) (err error) {
	return instance.SetProperty("QuickModeCryptoSet", (value))
}

// GetQuickModeCryptoSet gets the value of QuickModeCryptoSet for the instance
func (instance *MSFT_NetSARule) GetPropertyQuickModeCryptoSet() (value string, err error) {
	retValue, err := instance.GetProperty("QuickModeCryptoSet")
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

// SetRuleGroup sets the value of RuleGroup for the instance
func (instance *MSFT_NetSARule) SetPropertyRuleGroup(value string) (err error) {
	return instance.SetProperty("RuleGroup", (value))
}

// GetRuleGroup gets the value of RuleGroup for the instance
func (instance *MSFT_NetSARule) GetPropertyRuleGroup() (value string, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_NetSARule) GetPropertyStatus() (value string, err error) {
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
func (instance *MSFT_NetSARule) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", (value))
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *MSFT_NetSARule) GetPropertyStatusCode() (value uint32, err error) {
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
