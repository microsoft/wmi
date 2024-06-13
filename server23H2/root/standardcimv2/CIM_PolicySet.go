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

// CIM_PolicySet struct
type CIM_PolicySet struct {
	*CIM_Policy

	//
	Enabled uint16

	//
	PolicyDecisionStrategy uint16

	//
	PolicyRoles []string
}

func NewCIM_PolicySetEx1(instance *cim.WmiInstance) (newInstance *CIM_PolicySet, err error) {
	tmp, err := NewCIM_PolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicySet{
		CIM_Policy: tmp,
	}
	return
}

func NewCIM_PolicySetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PolicySet, err error) {
	tmp, err := NewCIM_PolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PolicySet{
		CIM_Policy: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *CIM_PolicySet) SetPropertyEnabled(value uint16) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *CIM_PolicySet) GetPropertyEnabled() (value uint16, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetPolicyDecisionStrategy sets the value of PolicyDecisionStrategy for the instance
func (instance *CIM_PolicySet) SetPropertyPolicyDecisionStrategy(value uint16) (err error) {
	return instance.SetProperty("PolicyDecisionStrategy", (value))
}

// GetPolicyDecisionStrategy gets the value of PolicyDecisionStrategy for the instance
func (instance *CIM_PolicySet) GetPropertyPolicyDecisionStrategy() (value uint16, err error) {
	retValue, err := instance.GetProperty("PolicyDecisionStrategy")
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

// SetPolicyRoles sets the value of PolicyRoles for the instance
func (instance *CIM_PolicySet) SetPropertyPolicyRoles(value []string) (err error) {
	return instance.SetProperty("PolicyRoles", (value))
}

// GetPolicyRoles gets the value of PolicyRoles for the instance
func (instance *CIM_PolicySet) GetPropertyPolicyRoles() (value []string, err error) {
	retValue, err := instance.GetProperty("PolicyRoles")
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
