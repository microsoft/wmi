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

// SPACES_MaskingSet struct
type SPACES_MaskingSet struct {
	*MSFT_MaskingSet
}

func NewSPACES_MaskingSetEx1(instance *cim.WmiInstance) (newInstance *SPACES_MaskingSet, err error) {
	tmp, err := NewMSFT_MaskingSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SPACES_MaskingSet{
		MSFT_MaskingSet: tmp,
	}
	return
}

func NewSPACES_MaskingSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SPACES_MaskingSet, err error) {
	tmp, err := NewMSFT_MaskingSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SPACES_MaskingSet{
		MSFT_MaskingSet: tmp,
	}
	return
}
