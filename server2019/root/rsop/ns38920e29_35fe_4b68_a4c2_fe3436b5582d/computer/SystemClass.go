// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D.Computer
//////////////////////////////////////////////
package computer
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// __SystemClass struct
type __SystemClass struct { 
	*cim.WmiInstance
}

	func New__SystemClassEx1(instance *cim.WmiInstance) (newInstance *__SystemClass, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &__SystemClass {
	WmiInstance: tmp,
	}
	return
	}
	

	func New__SystemClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__SystemClass, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__SystemClass {
	WmiInstance: tmp,
	}
	return
	}
	

