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

// MSFT_NetEventVmSwitch struct
type MSFT_NetEventVmSwitch struct {
	*MSFT_NetEventPacketCaptureTarget
}

func NewMSFT_NetEventVmSwitchEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventVmSwitch, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVmSwitch{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}

func NewMSFT_NetEventVmSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventVmSwitch, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVmSwitch{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}
