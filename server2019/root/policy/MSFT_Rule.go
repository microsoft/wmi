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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_Rule struct
type MSFT_Rule struct {
	*cim.WmiInstance

	//
	Query string

	//
	QueryLanguage string

	//
	TargetNameSpace string
}

func NewMSFT_RuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_Rule, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_Rule{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_RuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Rule, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Rule{
		WmiInstance: tmp,
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *MSFT_Rule) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", (value))
}

// GetQuery gets the value of Query for the instance
func (instance *MSFT_Rule) GetPropertyQuery() (value string, err error) {
	retValue, err := instance.GetProperty("Query")
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

// SetQueryLanguage sets the value of QueryLanguage for the instance
func (instance *MSFT_Rule) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", (value))
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *MSFT_Rule) GetPropertyQueryLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("QueryLanguage")
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

// SetTargetNameSpace sets the value of TargetNameSpace for the instance
func (instance *MSFT_Rule) SetPropertyTargetNameSpace(value string) (err error) {
	return instance.SetProperty("TargetNameSpace", (value))
}

// GetTargetNameSpace gets the value of TargetNameSpace for the instance
func (instance *MSFT_Rule) GetPropertyTargetNameSpace() (value string, err error) {
	retValue, err := instance.GetProperty("TargetNameSpace")
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
