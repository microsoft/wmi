// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SPACES_StorageScaleUnit struct
type SPACES_StorageScaleUnit struct {
	*MSFT_StorageScaleUnit
}

func NewSPACES_StorageScaleUnitEx1(instance *cim.WmiInstance) (newInstance *SPACES_StorageScaleUnit, err error) {
	tmp, err := NewMSFT_StorageScaleUnitEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageScaleUnit{
		MSFT_StorageScaleUnit: tmp,
	}
	return
}

func NewSPACES_StorageScaleUnitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_StorageScaleUnit, err error) {
	tmp, err := NewMSFT_StorageScaleUnitEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_StorageScaleUnit{
		MSFT_StorageScaleUnit: tmp,
	}
	return
}
