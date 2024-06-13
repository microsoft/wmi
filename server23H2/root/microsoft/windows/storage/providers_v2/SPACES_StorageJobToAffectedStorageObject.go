// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageJobToAffectedStorageObject struct
type SPACES_StorageJobToAffectedStorageObject struct {
	*MSFT_StorageJobToAffectedStorageObject
}

func NewSPACES_StorageJobToAffectedStorageObjectEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageJobToAffectedStorageObject, err error) {
	tmp, err := NewMSFT_StorageJobToAffectedStorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageJobToAffectedStorageObject{
		MSFT_StorageJobToAffectedStorageObject: tmp,
	}
	return
}

func NewSPACES_StorageJobToAffectedStorageObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageJobToAffectedStorageObject, err error) {
	tmp, err := NewMSFT_StorageJobToAffectedStorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageJobToAffectedStorageObject{
		MSFT_StorageJobToAffectedStorageObject: tmp,
	}
	return
}
