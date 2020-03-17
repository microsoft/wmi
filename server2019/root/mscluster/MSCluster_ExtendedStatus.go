// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_ExtendedStatus struct
type MSCluster_ExtendedStatus struct {
	__ExtendedStatus

	//
	ErrorType uint32
}

// SetErrorType sets the value of ErrorType for the instance
func (instance *MSCluster_ExtendedStatus) SetPropertyErrorType(value uint32) (err error) {
	return instance.SetProperty("ErrorType", value)
}

// GetErrorType gets the value of ErrorType for the instance
func (instance *MSCluster_ExtendedStatus) GetPropertyErrorType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
