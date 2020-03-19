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

// Msft_WmiProvider_AccessCheck_Pre struct
type Msft_WmiProvider_AccessCheck_Pre struct {
	*Msft_WmiProvider_OperationEvent_Pre

	//
	Query string

	//
	QueryLanguage string

	//
	Sid []uint8
}

func NewMsft_WmiProvider_AccessCheck_PreEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_AccessCheck_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_AccessCheck_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

func NewMsft_WmiProvider_AccessCheck_PreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_AccessCheck_Pre, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PreEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_AccessCheck_Pre{
		Msft_WmiProvider_OperationEvent_Pre: tmp,
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *Msft_WmiProvider_AccessCheck_Pre) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *Msft_WmiProvider_AccessCheck_Pre) GetPropertyQuery() (value string, err error) {
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
func (instance *Msft_WmiProvider_AccessCheck_Pre) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *Msft_WmiProvider_AccessCheck_Pre) GetPropertyQueryLanguage() (value string, err error) {
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
func (instance *Msft_WmiProvider_AccessCheck_Pre) SetPropertySid(value []uint8) (err error) {
	return instance.SetProperty("Sid", value)
}

// GetSid gets the value of Sid for the instance
func (instance *Msft_WmiProvider_AccessCheck_Pre) GetPropertySid() (value []uint8, err error) {
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
