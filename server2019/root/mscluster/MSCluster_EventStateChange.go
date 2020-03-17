// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_EventStateChange struct
type MSCluster_EventStateChange struct {
	MSCluster_Event

	//
	EventNewState uint32
}

// SetEventNewState sets the value of EventNewState for the instance
func (instance *MSCluster_EventStateChange) SetPropertyEventNewState(value uint32) (err error) {
	return instance.SetProperty("EventNewState", value)
}

// GetEventNewState gets the value of EventNewState for the instance
func (instance *MSCluster_EventStateChange) GetPropertyEventNewState() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventNewState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
