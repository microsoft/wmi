// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ConcreteCollection struct
type ConcreteCollection struct {
	*CIM_ConcreteCollection
}

func NewConcreteCollectionEx1(instance *cim.WmiInstance) (newInstance *ConcreteCollection, err error) {
	tmp, err := NewCIM_ConcreteCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConcreteCollection{
		CIM_ConcreteCollection: tmp,
	}
	return
}

func NewConcreteCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConcreteCollection, err error) {
	tmp, err := NewCIM_ConcreteCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConcreteCollection{
		CIM_ConcreteCollection: tmp,
	}
	return
}
