// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D
//////////////////////////////////////////////
package ns38920e29_35fe_4b68_a4c2_fe3436b5582d
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
	

