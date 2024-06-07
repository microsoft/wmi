// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS4D70B888_0C83_4EC8_A075_63C360913417
//////////////////////////////////////////////
package ns4d70b888_0c83_4ec8_a075_63c360913417
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// __SecurityRelatedClass struct
type __SecurityRelatedClass struct { 
	*cim.WmiInstance
}

	func New__SecurityRelatedClassEx1(instance *cim.WmiInstance) (newInstance *__SecurityRelatedClass, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &__SecurityRelatedClass {
	WmiInstance: tmp,
	}
	return
	}
	

	func New__SecurityRelatedClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__SecurityRelatedClass, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__SecurityRelatedClass {
	WmiInstance: tmp,
	}
	return
	}
	

