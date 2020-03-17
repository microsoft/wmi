// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.EventTracingManagement
//////////////////////////////////////////////
package eventtracingmanagement

// MSFT_EtwTraceProvider struct
type MSFT_EtwTraceProvider struct {
	CIM_LogicalElement

	//
	AutologgerName string

	//
	Guid string

	//
	Level uint8

	//
	MatchAllKeyword uint64

	//
	MatchAnyKeyword uint64

	//
	Property uint32

	//
	SessionName string
}

// SetAutologgerName sets the value of AutologgerName for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyAutologgerName(value string) (err error) {
	return instance.SetProperty("AutologgerName", value)
}

// GetAutologgerName gets the value of AutologgerName for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyAutologgerName() (value string, err error) {
	retValue, err := instance.GetProperty("AutologgerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLevel sets the value of Level for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyLevel(value uint8) (err error) {
	return instance.SetProperty("Level", value)
}

// GetLevel gets the value of Level for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMatchAllKeyword sets the value of MatchAllKeyword for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyMatchAllKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAllKeyword", value)
}

// GetMatchAllKeyword gets the value of MatchAllKeyword for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyMatchAllKeyword() (value uint64, err error) {
	retValue, err := instance.GetProperty("MatchAllKeyword")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMatchAnyKeyword sets the value of MatchAnyKeyword for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyMatchAnyKeyword(value uint64) (err error) {
	return instance.SetProperty("MatchAnyKeyword", value)
}

// GetMatchAnyKeyword gets the value of MatchAnyKeyword for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyMatchAnyKeyword() (value uint64, err error) {
	retValue, err := instance.GetProperty("MatchAnyKeyword")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProperty sets the value of Property for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertyProperty(value uint32) (err error) {
	return instance.SetProperty("Property", value)
}

// GetProperty gets the value of Property for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertyProperty() (value uint32, err error) {
	retValue, err := instance.GetProperty("Property")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSessionName sets the value of SessionName for the instance
func (instance *MSFT_EtwTraceProvider) SetPropertySessionName(value string) (err error) {
	return instance.SetProperty("SessionName", value)
}

// GetSessionName gets the value of SessionName for the instance
func (instance *MSFT_EtwTraceProvider) GetPropertySessionName() (value string, err error) {
	retValue, err := instance.GetProperty("SessionName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
