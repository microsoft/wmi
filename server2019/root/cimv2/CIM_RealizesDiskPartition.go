// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_RealizesDiskPartition struct
type CIM_RealizesDiskPartition struct {
	CIM_Realizes

	//
	StartingAddress uint64
}

// SetStartingAddress sets the value of StartingAddress for the instance
func (instance *CIM_RealizesDiskPartition) SetPropertyStartingAddress(value uint64) (err error) {
	return instance.SetProperty("StartingAddress", value)
}

// GetStartingAddress gets the value of StartingAddress for the instance
func (instance *CIM_RealizesDiskPartition) GetPropertyStartingAddress() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartingAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
