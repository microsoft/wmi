// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS691FCD74_C5E4_4C01_A889_46480EAB1D29.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_CentralAccessPolicySetting struct
type RSOP_CentralAccessPolicySetting struct {
	*RSOP_PolicySetting

	//
	CentralAccessPolicyName []string
}

func NewRSOP_CentralAccessPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_CentralAccessPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_CentralAccessPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_CentralAccessPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_CentralAccessPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_CentralAccessPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetCentralAccessPolicyName sets the value of CentralAccessPolicyName for the instance
func (instance *RSOP_CentralAccessPolicySetting) SetPropertyCentralAccessPolicyName(value []string) (err error) {
	return instance.SetProperty("CentralAccessPolicyName", (value))
}

// GetCentralAccessPolicyName gets the value of CentralAccessPolicyName for the instance
func (instance *RSOP_CentralAccessPolicySetting) GetPropertyCentralAccessPolicyName() (value []string, err error) {
	retValue, err := instance.GetProperty("CentralAccessPolicyName")
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
