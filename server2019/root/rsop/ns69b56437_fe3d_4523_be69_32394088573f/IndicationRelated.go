// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS69B56437_FE3D_4523_BE69_32394088573F
//////////////////////////////////////////////
package ns69b56437_fe3d_4523_be69_32394088573f
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __IndicationRelated struct
type __IndicationRelated struct { 
	*__SystemClass
}

	func New__IndicationRelatedEx1(instance *cim.WmiInstance) (newInstance *__IndicationRelated, err error) {tmp, err := New__SystemClassEx1(instance)
		
	if err != nil { return }
	newInstance = &__IndicationRelated {
	__SystemClass: tmp,
	}
	return
	}
	

	func New__IndicationRelatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__IndicationRelated, err error) {tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__IndicationRelated {
	__SystemClass: tmp,
	}
	return
	}
	

