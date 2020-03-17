// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_EventGroupStateChange struct
type MSCluster_EventGroupStateChange struct {
	MSCluster_EventStateChange

	//
	EventNode string
}

// SetEventNode sets the value of EventNode for the instance
func (instance *MSCluster_EventGroupStateChange) SetPropertyEventNode(value string) (err error) {
	return instance.SetProperty("EventNode", value)
}

// GetEventNode gets the value of EventNode for the instance
func (instance *MSCluster_EventGroupStateChange) GetPropertyEventNode() (value string, err error) {
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
