// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SynchronousReceiveBehavior struct
type SynchronousReceiveBehavior struct {
	*Behavior
}

func NewSynchronousReceiveBehaviorEx1(instance *cim.WmiInstance) (newInstance *SynchronousReceiveBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SynchronousReceiveBehavior{
		Behavior: tmp,
	}
	return
}

func NewSynchronousReceiveBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SynchronousReceiveBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SynchronousReceiveBehavior{
		Behavior: tmp,
	}
	return
}
