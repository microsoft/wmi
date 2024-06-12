// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// KerbChangePassword_Start struct
type KerbChangePassword_Start struct {
	*KerbChangePassword
}

func NewKerbChangePassword_StartEx1(instance *cim.WmiInstance) (newInstance *KerbChangePassword_Start, err error) {
	tmp, err := NewKerbChangePasswordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword_Start{
		KerbChangePassword: tmp,
	}
	return
}

func NewKerbChangePassword_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbChangePassword_Start, err error) {
	tmp, err := NewKerbChangePasswordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword_Start{
		KerbChangePassword: tmp,
	}
	return
}
