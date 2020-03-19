// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_SARule struct
type CIM_SARule struct {
	*CIM_PolicyRule

	//
	LimitNegotiation uint16
}

func NewCIM_SARuleEx1(instance *cim.WmiInstance) (newInstance *CIM_SARule, err error) {
	tmp, err := NewCIM_PolicyRuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SARule{
		CIM_PolicyRule: tmp,
	}
	return
}

func NewCIM_SARuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SARule, err error) {
	tmp, err := NewCIM_PolicyRuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SARule{
		CIM_PolicyRule: tmp,
	}
	return
}

// SetLimitNegotiation sets the value of LimitNegotiation for the instance
func (instance *CIM_SARule) SetPropertyLimitNegotiation(value uint16) (err error) {
	return instance.SetProperty("LimitNegotiation", value)
}

// GetLimitNegotiation gets the value of LimitNegotiation for the instance
func (instance *CIM_SARule) GetPropertyLimitNegotiation() (value uint16, err error) {
	retValue, err := instance.GetProperty("LimitNegotiation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
