// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_Service struct
type MSCluster_Service struct {
	CIM_ClusteringService

	//
	NodeHighestVersion uint32

	//
	NodeLowestVersion uint32

	//
	State string
}

// SetNodeHighestVersion sets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Service) SetPropertyNodeHighestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeHighestVersion", value)
}

// GetNodeHighestVersion gets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Service) GetPropertyNodeHighestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeHighestVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeLowestVersion sets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Service) SetPropertyNodeLowestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeLowestVersion", value)
}

// GetNodeLowestVersion gets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Service) GetPropertyNodeLowestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeLowestVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_Service) SetPropertyState(value string) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSCluster_Service) GetPropertyState() (value string, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//
func (instance *MSCluster_Service) Start() (err error) {
	_, err = instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Service) Stop() (err error) {
	_, err = instance.InvokeMethodWithReturn("Stop")
	if err != nil {
		return
	}
	return

}
