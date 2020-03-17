// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_NewQuery_Pre struct
type Msft_WmiProvider_NewQuery_Pre struct {
	Msft_WmiProvider_OperationEvent_Pre

	//
	Query string

	//
	QueryId uint32

	//
	QueryLanguage string
}

// SetQuery sets the value of Query for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) GetPropertyQuery() (value string, err error) {
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

// SetQueryId sets the value of QueryId for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) SetPropertyQueryId(value uint32) (err error) {
	return instance.SetProperty("QueryId", value)
}

// GetQueryId gets the value of QueryId for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) GetPropertyQueryId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueryId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueryLanguage sets the value of QueryLanguage for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *Msft_WmiProvider_NewQuery_Pre) GetPropertyQueryLanguage() (value string, err error) {
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
