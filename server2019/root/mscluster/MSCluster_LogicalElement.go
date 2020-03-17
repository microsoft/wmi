// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_LogicalElement struct
type MSCluster_LogicalElement struct {
	CIM_LogicalElement

	//
	Characteristics uint32

	//
	Flags uint32
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *MSCluster_LogicalElement) SetPropertyCharacteristics(value uint32) (err error) {
	return instance.SetProperty("Characteristics", value)
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *MSCluster_LogicalElement) GetPropertyCharacteristics() (value uint32, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSCluster_LogicalElement) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", value)
}

// GetFlags gets the value of Flags for the instance
func (instance *MSCluster_LogicalElement) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
