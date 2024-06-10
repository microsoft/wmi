// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Policy
//
// ////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("RuleValidationResults", (value))
}

// GetRuleValidationResults gets the value of RuleValidationResults for the instance
func (instance *SomFilterPutStatus) GetPropertyRuleValidationResults() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RuleValidationResults")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}
