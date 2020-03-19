// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SomFilterPutStatus struct
type SomFilterPutStatus struct {
	*__ExtendedStatus

	//
	RuleValidationResults []uint32
}

func NewSomFilterPutStatusEx1(instance *cim.WmiInstance) (newInstance *SomFilterPutStatus, err error) {
	tmp, err := New__ExtendedStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SomFilterPutStatus{
		__ExtendedStatus: tmp,
	}
	return
}

func NewSomFilterPutStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SomFilterPutStatus, err error) {
	tmp, err := New__ExtendedStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SomFilterPutStatus{
		__ExtendedStatus: tmp,
	}
	return
}

// SetRuleValidationResults sets the value of RuleValidationResults for the instance
func (instance *SomFilterPutStatus) SetPropertyRuleValidationResults(value []uint32) (err error) {
	return instance.SetProperty("RuleValidationResults", value)
}

// GetRuleValidationResults gets the value of RuleValidationResults for the instance
func (instance *SomFilterPutStatus) GetPropertyRuleValidationResults() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RuleValidationResults")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
