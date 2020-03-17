// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// SPACES_StorageJob struct
type SPACES_StorageJob struct {
	MSFT_StorageJob

	// ObjectId for internal use only.
	UpdatedObjectId string
}

// SetUpdatedObjectId sets the value of UpdatedObjectId for the instance
func (instance *SPACES_StorageJob) SetPropertyUpdatedObjectId(value string) (err error) {
	return instance.SetProperty("UpdatedObjectId", value)
}

// GetUpdatedObjectId gets the value of UpdatedObjectId for the instance
func (instance *SPACES_StorageJob) GetPropertyUpdatedObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("UpdatedObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
