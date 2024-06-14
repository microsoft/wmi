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

// ApplicationContainsConfigurationSection struct
type ApplicationContainsConfigurationSection struct {
	*ObjectConfigurationAssociation
}

func NewApplicationContainsConfigurationSectionEx1(instance *cim.WmiInstance) (newInstance *ApplicationContainsConfigurationSection, err error) {
	tmp, err := NewObjectConfigurationAssociationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ApplicationContainsConfigurationSection{
		ObjectConfigurationAssociation: tmp,
	}
	return
}

func NewApplicationContainsConfigurationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ApplicationContainsConfigurationSection, err error) {
	tmp, err := NewObjectConfigurationAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ApplicationContainsConfigurationSection{
		ObjectConfigurationAssociation: tmp,
	}
	return
}
