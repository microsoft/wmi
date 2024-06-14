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

// KerbLogonUser_Start struct
type KerbLogonUser_Start struct {
	*KerbLogonUser
}

func NewKerbLogonUser_StartEx1(instance *cim.WmiInstance) (newInstance *KerbLogonUser_Start, err error) {
	tmp, err := NewKerbLogonUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbLogonUser_Start{
		KerbLogonUser: tmp,
	}
	return
}

func NewKerbLogonUser_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbLogonUser_Start, err error) {
	tmp, err := NewKerbLogonUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbLogonUser_Start{
		KerbLogonUser: tmp,
	}
	return
}
