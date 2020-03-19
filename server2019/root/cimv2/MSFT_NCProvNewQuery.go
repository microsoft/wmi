// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NCProvNewQuery struct
type MSFT_NCProvNewQuery struct {
	*MSFT_NCProvEvent

	//
	ID uint32

	//
	Query string

	//
	QueryLanguage string
}

func NewMSFT_NCProvNewQueryEx1(instance *cim.WmiInstance) (newInstance *MSFT_NCProvNewQuery, err error) {
	tmp, err := NewMSFT_NCProvEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvNewQuery{
		MSFT_NCProvEvent: tmp,
	}
	return
}

func NewMSFT_NCProvNewQueryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NCProvNewQuery, err error) {
	tmp, err := NewMSFT_NCProvEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvNewQuery{
		MSFT_NCProvEvent: tmp,
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MSFT_NCProvNewQuery) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MSFT_NCProvNewQuery) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *MSFT_NCProvNewQuery) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *MSFT_NCProvNewQuery) GetPropertyQuery() (value string, err error) {
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
func (instance *MSFT_NCProvNewQuery) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *MSFT_NCProvNewQuery) GetPropertyQueryLanguage() (value string, err error) {
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
