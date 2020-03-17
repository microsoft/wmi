// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_EventResourceStateChange struct
type MSCluster_EventResourceStateChange struct {
	MSCluster_EventStateChange

	//
	EventGroup string

	//
	EventNode string
}

// SetEventGroup sets the value of EventGroup for the instance
func (instance *MSCluster_EventResourceStateChange) SetPropertyEventGroup(value string) (err error) {
	return instance.SetProperty("EventGroup", value)
}

// GetEventGroup gets the value of EventGroup for the instance
func (instance *MSCluster_EventResourceStateChange) GetPropertyEventGroup() (value string, err error) {
	retValue, err := instance.GetProperty("EventGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventNode sets the value of EventNode for the instance
func (instance *MSCluster_EventResourceStateChange) SetPropertyEventNode(value string) (err error) {
	return instance.SetProperty("EventNode", value)
}

// GetEventNode gets the value of EventNode for the instance
func (instance *MSCluster_EventResourceStateChange) GetPropertyEventNode() (value string, err error) {
	retValue, err := instance.GetProperty("EventNode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
