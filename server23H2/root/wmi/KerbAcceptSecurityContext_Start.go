// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// KerbAcceptSecurityContext_Start struct
type KerbAcceptSecurityContext_Start struct {
	*KerbAcceptSecurityContext
}

func NewKerbAcceptSecurityContext_StartEx1(instance *cim.WmiInstance) (newInstance *KerbAcceptSecurityContext_Start, err error) {
	tmp, err := NewKerbAcceptSecurityContextEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbAcceptSecurityContext_Start{
		KerbAcceptSecurityContext: tmp,
	}
	return
}

func NewKerbAcceptSecurityContext_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbAcceptSecurityContext_Start, err error) {
	tmp, err := NewKerbAcceptSecurityContextEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbAcceptSecurityContext_Start{
		KerbAcceptSecurityContext: tmp,
	}
	return
}
