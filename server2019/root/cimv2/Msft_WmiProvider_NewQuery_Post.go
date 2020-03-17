// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Msft_WmiProvider_NewQuery_Post struct
type Msft_WmiProvider_NewQuery_Post struct {
	Msft_WmiProvider_OperationEvent_Post

	//
	Query string

	//
	QueryId uint32

	//
	QueryLanguage string

	//
	Result uint32
}

// SetQuery sets the value of Query for the instance
func (instance *Msft_WmiProvider_NewQuery_Post) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *Msft_WmiProvider_NewQuery_Post) GetPropertyQuery() (value string, err error) {
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
func (instance *Msft_WmiProvider_NewQuery_Post) SetPropertyQueryId(value uint32) (err error) {
	return instance.SetProperty("QueryId", value)
}

// GetQueryId gets the value of QueryId for the instance
func (instance *Msft_WmiProvider_NewQuery_Post) GetPropertyQueryId() (value uint32, err error) {
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
func (instance *Msft_WmiProvider_NewQuery_Post) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *Msft_WmiProvider_NewQuery_Post) GetPropertyQueryLanguage() (value string, err error) {
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
func (instance *Msft_WmiProvider_NewQuery_Post) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *Msft_WmiProvider_NewQuery_Post) GetPropertyResult() (value uint32, err error) {
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
