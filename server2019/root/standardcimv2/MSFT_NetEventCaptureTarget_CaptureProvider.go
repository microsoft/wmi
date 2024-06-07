// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetEventCaptureTarget_CaptureProvider struct
type MSFT_NetEventCaptureTarget_CaptureProvider struct { 
	*CIM_Component
}

	func NewMSFT_NetEventCaptureTarget_CaptureProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventCaptureTarget_CaptureProvider, err error) {tmp, err := NewCIM_ComponentEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventCaptureTarget_CaptureProvider {
	CIM_Component: tmp,
	}
	return
	}
	

	func NewMSFT_NetEventCaptureTarget_CaptureProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetEventCaptureTarget_CaptureProvider, err error) {tmp, err := NewCIM_ComponentEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventCaptureTarget_CaptureProvider {
	CIM_Component: tmp,
	}
	return
	}
	

