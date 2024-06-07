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
 "github.com/microsoft/wmi/pkg/base/instance"
)

// MSFT_DASettingsIndication struct
type MSFT_DASettingsIndication struct { 
	*cim.WmiInstance
}

	func NewMSFT_DASettingsIndicationEx1(instance *cim.WmiInstance) (newInstance *MSFT_DASettingsIndication, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_DASettingsIndication {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_DASettingsIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DASettingsIndication, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DASettingsIndication {
	WmiInstance: tmp,
	}
	return
	}
	

