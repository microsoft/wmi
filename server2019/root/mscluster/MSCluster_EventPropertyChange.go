// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_EventPropertyChange struct
type MSCluster_EventPropertyChange struct {
	MSCluster_Event

	//
	EventProperty string
}

// SetEventProperty sets the value of EventProperty for the instance
func (instance *MSCluster_EventPropertyChange) SetPropertyEventProperty(value string) (err error) {
	return instance.SetProperty("EventProperty", value)
}

// GetEventProperty gets the value of EventProperty for the instance
func (instance *MSCluster_EventPropertyChange) GetPropertyEventProperty() (value string, err error) {
	retValue, err := instance.GetProperty("EventProperty")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
