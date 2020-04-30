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

// MSFT_NCProvAccessCheck struct
type MSFT_NCProvAccessCheck struct {
	*MSFT_NCProvEvent

	//
	Query string

	//
	QueryLanguage string

	//
	Sid []uint8
}

func NewMSFT_NCProvAccessCheckEx1(instance *cim.WmiInstance) (newInstance *MSFT_NCProvAccessCheck, err error) {
	tmp, err := NewMSFT_NCProvEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvAccessCheck{
		MSFT_NCProvEvent: tmp,
	}
	return
}

func NewMSFT_NCProvAccessCheckEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NCProvAccessCheck, err error) {
	tmp, err := NewMSFT_NCProvEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCProvAccessCheck{
		MSFT_NCProvEvent: tmp,
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *MSFT_NCProvAccessCheck) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *MSFT_NCProvAccessCheck) GetPropertyQuery() (value string, err error) {
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
func (instance *MSFT_NCProvAccessCheck) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *MSFT_NCProvAccessCheck) GetPropertyQueryLanguage() (value string, err error) {
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

// SetSid sets the value of Sid for the instance
func (instance *MSFT_NCProvAccessCheck) SetPropertySid(value []uint8) (err error) {
	return instance.SetProperty("Sid", value)
}

// GetSid gets the value of Sid for the instance
func (instance *MSFT_NCProvAccessCheck) GetPropertySid() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Sid")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}