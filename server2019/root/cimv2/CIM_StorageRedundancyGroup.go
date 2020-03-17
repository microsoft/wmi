// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_StorageRedundancyGroup struct
type CIM_StorageRedundancyGroup struct {
	CIM_RedundancyGroup

	//
	TypeOfAlgorithm uint16
}

// SetTypeOfAlgorithm sets the value of TypeOfAlgorithm for the instance
func (instance *CIM_StorageRedundancyGroup) SetPropertyTypeOfAlgorithm(value uint16) (err error) {
	return instance.SetProperty("TypeOfAlgorithm", value)
}

// GetTypeOfAlgorithm gets the value of TypeOfAlgorithm for the instance
func (instance *CIM_StorageRedundancyGroup) GetPropertyTypeOfAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("TypeOfAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
