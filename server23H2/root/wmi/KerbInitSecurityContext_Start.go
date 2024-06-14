// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// KerbInitSecurityContext_Start struct
type KerbInitSecurityContext_Start struct {
	*KerbInitSecurityContext
}

func NewKerbInitSecurityContext_StartEx1(instance *cim.WmiInstance) (newInstance *KerbInitSecurityContext_Start, err error) {
	tmp, err := NewKerbInitSecurityContextEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext_Start{
		KerbInitSecurityContext: tmp,
	}
	return
}

func NewKerbInitSecurityContext_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbInitSecurityContext_Start, err error) {
	tmp, err := NewKerbInitSecurityContextEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext_Start{
		KerbInitSecurityContext: tmp,
	}
	return
}
