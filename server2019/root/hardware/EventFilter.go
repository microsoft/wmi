// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// __EventFilter struct
type __EventFilter struct {
	__IndicationRelated

	//
	CreatorSID []uint8

	//
	EventAccess string

	//
	EventNamespace string

	//
	Name string

	//
	Query string

	//
	QueryLanguage string
}

// SetCreatorSID sets the value of CreatorSID for the instance
func (instance *__EventFilter) SetPropertyCreatorSID(value []uint8) (err error) {
	return instance.SetProperty("CreatorSID", value)
}

// GetCreatorSID gets the value of CreatorSID for the instance
func (instance *__EventFilter) GetPropertyCreatorSID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CreatorSID")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventAccess sets the value of EventAccess for the instance
func (instance *__EventFilter) SetPropertyEventAccess(value string) (err error) {
	return instance.SetProperty("EventAccess", value)
}

// GetEventAccess gets the value of EventAccess for the instance
func (instance *__EventFilter) GetPropertyEventAccess() (value string, err error) {
	retValue, err := instance.GetProperty("EventAccess")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventNamespace sets the value of EventNamespace for the instance
func (instance *__EventFilter) SetPropertyEventNamespace(value string) (err error) {
	return instance.SetProperty("EventNamespace", value)
}

// GetEventNamespace gets the value of EventNamespace for the instance
func (instance *__EventFilter) GetPropertyEventNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("EventNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *__EventFilter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *__EventFilter) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuery sets the value of Query for the instance
func (instance *__EventFilter) SetPropertyQuery(value string) (err error) {
	return instance.SetProperty("Query", value)
}

// GetQuery gets the value of Query for the instance
func (instance *__EventFilter) GetPropertyQuery() (value string, err error) {
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
func (instance *__EventFilter) SetPropertyQueryLanguage(value string) (err error) {
	return instance.SetProperty("QueryLanguage", value)
}

// GetQueryLanguage gets the value of QueryLanguage for the instance
func (instance *__EventFilter) GetPropertyQueryLanguage() (value string, err error) {
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
