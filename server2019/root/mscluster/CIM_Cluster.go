// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// CIM_Cluster struct
type CIM_Cluster struct {
	CIM_ComputerSystem

	//
	MaxNumberOfNodes uint32
}

// SetMaxNumberOfNodes sets the value of MaxNumberOfNodes for the instance
func (instance *CIM_Cluster) SetPropertyMaxNumberOfNodes(value uint32) (err error) {
	return instance.SetProperty("MaxNumberOfNodes", value)
}

// GetMaxNumberOfNodes gets the value of MaxNumberOfNodes for the instance
func (instance *CIM_Cluster) GetPropertyMaxNumberOfNodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfNodes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
