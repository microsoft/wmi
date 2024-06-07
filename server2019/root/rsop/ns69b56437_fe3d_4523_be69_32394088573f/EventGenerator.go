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

// __EventGenerator struct
type __EventGenerator struct { 
	*__IndicationRelated
}

	func New__EventGeneratorEx1(instance *cim.WmiInstance) (newInstance *__EventGenerator, err error) {tmp, err := New__IndicationRelatedEx1(instance)
		
	if err != nil { return }
	newInstance = &__EventGenerator {
	__IndicationRelated: tmp,
	}
	return
	}
	

	func New__EventGeneratorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__EventGenerator, err error) {tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__EventGenerator {
	__IndicationRelated: tmp,
	}
	return
	}
	

