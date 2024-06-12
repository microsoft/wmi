// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS02DDE4F9_9A42_44F8_A277_072194B38AC9.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IPSECPolicySetting struct
type RSOP_IPSECPolicySetting struct {
	*RSOP_PolicySetting

	//
	ClassName string

	//
	description string

	//
	ipsecData []uint8

	//
	ipsecDataType uint32

	//
	ipsecFilterReference []string

	//
	ipsecID string

	//
	ipsecISAKMPReference string

	//
	ipsecName string

	//
	ipsecNegotiationPolicyAction string

	//
	ipsecNegotiationPolicyReference string

	//
	ipsecNegotiationPolicyType string

	//
	ipsecNFAReference []string

	//
	ipsecOwnersReference []string

	//
	whenChanged uint32
}

func NewRSOP_IPSECPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_IPSECPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IPSECPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_IPSECPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IPSECPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IPSECPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetClassName sets the value of ClassName for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyClassName(value string) (err error) {
	return instance.SetProperty("ClassName", (value))
}

// GetClassName gets the value of ClassName for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ClassName")
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

// Setdescription sets the value of description for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertydescription(value string) (err error) {
	return instance.SetProperty("description", (value))
}

// Getdescription gets the value of description for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertydescription() (value string, err error) {
	retValue, err := instance.GetProperty("description")
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

// SetipsecData sets the value of ipsecData for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecData(value []uint8) (err error) {
	return instance.SetProperty("ipsecData", (value))
}

// GetipsecData gets the value of ipsecData for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ipsecData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetipsecDataType sets the value of ipsecDataType for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecDataType(value uint32) (err error) {
	return instance.SetProperty("ipsecDataType", (value))
}

// GetipsecDataType gets the value of ipsecDataType for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecDataType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ipsecDataType")
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

// SetipsecFilterReference sets the value of ipsecFilterReference for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecFilterReference(value []string) (err error) {
	return instance.SetProperty("ipsecFilterReference", (value))
}

// GetipsecFilterReference gets the value of ipsecFilterReference for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecFilterReference() (value []string, err error) {
	retValue, err := instance.GetProperty("ipsecFilterReference")
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

// SetipsecID sets the value of ipsecID for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecID(value string) (err error) {
	return instance.SetProperty("ipsecID", (value))
}

// GetipsecID gets the value of ipsecID for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecID() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecID")
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

// SetipsecISAKMPReference sets the value of ipsecISAKMPReference for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecISAKMPReference(value string) (err error) {
	return instance.SetProperty("ipsecISAKMPReference", (value))
}

// GetipsecISAKMPReference gets the value of ipsecISAKMPReference for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecISAKMPReference() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecISAKMPReference")
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

// SetipsecName sets the value of ipsecName for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecName(value string) (err error) {
	return instance.SetProperty("ipsecName", (value))
}

// GetipsecName gets the value of ipsecName for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecName() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecName")
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

// SetipsecNegotiationPolicyAction sets the value of ipsecNegotiationPolicyAction for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecNegotiationPolicyAction(value string) (err error) {
	return instance.SetProperty("ipsecNegotiationPolicyAction", (value))
}

// GetipsecNegotiationPolicyAction gets the value of ipsecNegotiationPolicyAction for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecNegotiationPolicyAction() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecNegotiationPolicyAction")
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

// SetipsecNegotiationPolicyReference sets the value of ipsecNegotiationPolicyReference for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecNegotiationPolicyReference(value string) (err error) {
	return instance.SetProperty("ipsecNegotiationPolicyReference", (value))
}

// GetipsecNegotiationPolicyReference gets the value of ipsecNegotiationPolicyReference for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecNegotiationPolicyReference() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecNegotiationPolicyReference")
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

// SetipsecNegotiationPolicyType sets the value of ipsecNegotiationPolicyType for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecNegotiationPolicyType(value string) (err error) {
	return instance.SetProperty("ipsecNegotiationPolicyType", (value))
}

// GetipsecNegotiationPolicyType gets the value of ipsecNegotiationPolicyType for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecNegotiationPolicyType() (value string, err error) {
	retValue, err := instance.GetProperty("ipsecNegotiationPolicyType")
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

// SetipsecNFAReference sets the value of ipsecNFAReference for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecNFAReference(value []string) (err error) {
	return instance.SetProperty("ipsecNFAReference", (value))
}

// GetipsecNFAReference gets the value of ipsecNFAReference for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecNFAReference() (value []string, err error) {
	retValue, err := instance.GetProperty("ipsecNFAReference")
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

// SetipsecOwnersReference sets the value of ipsecOwnersReference for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertyipsecOwnersReference(value []string) (err error) {
	return instance.SetProperty("ipsecOwnersReference", (value))
}

// GetipsecOwnersReference gets the value of ipsecOwnersReference for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertyipsecOwnersReference() (value []string, err error) {
	retValue, err := instance.GetProperty("ipsecOwnersReference")
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

// SetwhenChanged sets the value of whenChanged for the instance
func (instance *RSOP_IPSECPolicySetting) SetPropertywhenChanged(value uint32) (err error) {
	return instance.SetProperty("whenChanged", (value))
}

// GetwhenChanged gets the value of whenChanged for the instance
func (instance *RSOP_IPSECPolicySetting) GetPropertywhenChanged() (value uint32, err error) {
	retValue, err := instance.GetProperty("whenChanged")
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
