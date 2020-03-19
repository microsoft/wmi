// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageJob struct
type SPACES_StorageJob struct {
	*MSFT_StorageJob

	// ObjectId for internal use only.
	UpdatedObjectId string
}

func NewSPACES_StorageJobEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageJob, err error) {
	tmp, err := NewMSFT_StorageJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageJob{
		MSFT_StorageJob: tmp,
	}
	return
}

func NewSPACES_StorageJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageJob, err error) {
	tmp, err := NewMSFT_StorageJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageJob{
		MSFT_StorageJob: tmp,
	}
	return
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
