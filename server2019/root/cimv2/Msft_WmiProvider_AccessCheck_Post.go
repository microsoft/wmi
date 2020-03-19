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

// Msft_WmiProvider_AccessCheck_Post struct
type Msft_WmiProvider_AccessCheck_Post struct {
	*Msft_WmiProvider_OperationEvent_Post

	//
	Query string

	//
	QueryLanguage string

	//
	Result uint32

	//
	Sid []uint8
}

func NewMsft_WmiProvider_AccessCheck_PostEx1(instance *cim.WmiInstance) (newInstance *Msft_WmiProvider_AccessCheck_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_AccessCheck_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

func NewMsft_WmiProvider_AccessCheck_PostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_WmiProvider_AccessCheck_Post, err error) {
	tmp, err := NewMsft_WmiProvider_OperationEvent_PostEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_WmiProvider_AccessCheck_Post{
		Msft_WmiProvider_OperationEvent_Post: tmp,
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) GetPropertyQuery() (value string, err error) {
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
func (instance *Msft_WmiProvider_AccessCheck_Post) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) GetPropertyQueryLanguage() (value string, err error) {
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

// SetResult sets the value of Result for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) GetPropertyResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("Result")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSid sets the value of Sid for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) SetPropertySid(value []uint8) (err error) {
	return instance.SetProperty("Sid", value)
}

// GetSid gets the value of Sid for the instance
func (instance *Msft_WmiProvider_AccessCheck_Post) GetPropertySid() (value []uint8, err error) {
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
