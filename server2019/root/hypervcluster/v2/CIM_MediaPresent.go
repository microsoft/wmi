// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// CIM_MediaPresent struct
type CIM_MediaPresent struct {
	CIM_Dependency

	// Boolean indicating that the accessed StorageExtent is fixed in the MediaAccessDevice and can not be ejected.
	FixedMedia bool
}

// SetFixedMedia sets the value of FixedMedia for the instance
func (instance *CIM_MediaPresent) SetPropertyFixedMedia(value bool) (err error) {
	return instance.SetProperty("FixedMedia", value)
}

// GetFixedMedia gets the value of FixedMedia for the instance
func (instance *CIM_MediaPresent) GetPropertyFixedMedia() (value bool, err error) {
	retValue, err := instance.GetProperty("FixedMedia")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
