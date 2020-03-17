// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// __AggregateEvent struct
type __AggregateEvent struct {
	__IndicationRelated

	//
	NumberOfEvents uint32

	//
	Representative interface{}
}

// SetNumberOfEvents sets the value of NumberOfEvents for the instance
func (instance *__AggregateEvent) SetPropertyNumberOfEvents(value uint32) (err error) {
	return instance.SetProperty("NumberOfEvents", value)
}

// GetNumberOfEvents gets the value of NumberOfEvents for the instance
func (instance *__AggregateEvent) GetPropertyNumberOfEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfEvents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRepresentative sets the value of Representative for the instance
func (instance *__AggregateEvent) SetPropertyRepresentative(value interface{}) (err error) {
	return instance.SetProperty("Representative", value)
}

// GetRepresentative gets the value of Representative for the instance
func (instance *__AggregateEvent) GetPropertyRepresentative() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Representative")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
