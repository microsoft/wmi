// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_Rule struct
type MSFT_Rule struct {
	cim.WmiInstance

	//
	Query string

	//
	QueryLanguage string

	//
	TargetNameSpace string
}

// SetQuery sets the value of Query for the instance
func (instance *MSFT_Rule) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *MSFT_Rule) GetPropertyQuery() (value string, err error) {
	retValue, err := instance.GetProperty("Query")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueryLanguage sets the value of QueryLanguage for the instance
func (instance *MSFT_Rule) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *MSFT_Rule) GetPropertyQueryLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("QueryLanguage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetNameSpace sets the value of TargetNameSpace for the instance
func (instance *MSFT_Rule) SetPropertyTargetNameSpace(value string) (err error) {
	return instance.SetProperty("TargetNameSpace", value)
}

// GetTargetNameSpace gets the value of TargetNameSpace for the instance
func (instance *MSFT_Rule) GetPropertyTargetNameSpace() (value string, err error) {
	retValue, err := instance.GetProperty("TargetNameSpace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
