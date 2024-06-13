// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ResourceInDesiredState struct
type MSFT_ResourceInDesiredState struct {
	*MSFT_DSCResource
}

func NewMSFT_ResourceInDesiredStateEx1(instance *cim.WmiInstance) (newInstance *MSFT_ResourceInDesiredState, err error) {
	tmp, err := NewMSFT_DSCResourceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_ResourceInDesiredState{
		MSFT_DSCResource: tmp,
	}
	return
}

func NewMSFT_ResourceInDesiredStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ResourceInDesiredState, err error) {
	tmp, err := NewMSFT_DSCResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ResourceInDesiredState{
		MSFT_DSCResource: tmp,
	}
	return
}
