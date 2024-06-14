// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ServerContainsApplicationPool struct
type ServerContainsApplicationPool struct {
	*ObjectContainerAssociation
}

func NewServerContainsApplicationPoolEx1(instance *cim.WmiInstance) (newInstance *ServerContainsApplicationPool, err error) {
	tmp, err := NewObjectContainerAssociationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerContainsApplicationPool{
		ObjectContainerAssociation: tmp,
	}
	return
}

func NewServerContainsApplicationPoolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerContainsApplicationPool, err error) {
	tmp, err := NewObjectContainerAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerContainsApplicationPool{
		ObjectContainerAssociation: tmp,
	}
	return
}
