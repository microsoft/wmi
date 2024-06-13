// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// DS_LDAP_Root_Class struct
type DS_LDAP_Root_Class struct {
	*cim.WmiInstance
}

func NewDS_LDAP_Root_ClassEx1(instance *cim.WmiInstance) (newInstance *DS_LDAP_Root_Class, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Root_Class{
		WmiInstance: tmp,
	}
	return
}

func NewDS_LDAP_Root_ClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DS_LDAP_Root_Class, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DS_LDAP_Root_Class{
		WmiInstance: tmp,
	}
	return
}
