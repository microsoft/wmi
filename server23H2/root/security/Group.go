// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.SECURITY
//////////////////////////////////////////////
package security

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __Group struct
type __Group struct {
	*__Subject
}

func New__GroupEx1(instance *cim.WmiInstance) (newInstance *__Group, err error) {
	tmp, err := New__SubjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__Group{
		__Subject: tmp,
	}
	return
}

func New__GroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__Group, err error) {
	tmp, err := New__SubjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__Group{
		__Subject: tmp,
	}
	return
}
