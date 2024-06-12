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

// KerbSetPassword_Start struct
type KerbSetPassword_Start struct {
	*KerbSetPassword
}

func NewKerbSetPassword_StartEx1(instance *cim.WmiInstance) (newInstance *KerbSetPassword_Start, err error) {
	tmp, err := NewKerbSetPasswordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword_Start{
		KerbSetPassword: tmp,
	}
	return
}

func NewKerbSetPassword_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbSetPassword_Start, err error) {
	tmp, err := NewKerbSetPasswordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword_Start{
		KerbSetPassword: tmp,
	}
	return
}
