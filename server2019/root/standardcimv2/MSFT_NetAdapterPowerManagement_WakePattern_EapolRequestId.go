// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId struct
type MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId struct {
	*MSFT_NetAdapterPowerManagement_WakePattern
}

func NewMSFT_NetAdapterPowerManagement_WakePattern_EapolRequestIdEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId{
		MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_WakePattern_EapolRequestIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_EapolRequestId{
		MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
}
