// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting struct
type MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting struct {
	*MSFT_NetAdapterElementSettingData
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting, err error) {
	tmp, err := NewMSFT_NetAdapterElementSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadElementSetting{
		MSFT_NetAdapterElementSettingData: tmp,
	}
	return
}
