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

// VirtualDirectoryContainsConfigurationSection struct
type VirtualDirectoryContainsConfigurationSection struct {
	*ObjectConfigurationAssociation
}

func NewVirtualDirectoryContainsConfigurationSectionEx1(instance *cim.WmiInstance) (newInstance *VirtualDirectoryContainsConfigurationSection, err error) {
	tmp, err := NewObjectConfigurationAssociationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectoryContainsConfigurationSection{
		ObjectConfigurationAssociation: tmp,
	}
	return
}

func NewVirtualDirectoryContainsConfigurationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VirtualDirectoryContainsConfigurationSection, err error) {
	tmp, err := NewObjectConfigurationAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VirtualDirectoryContainsConfigurationSection{
		ObjectConfigurationAssociation: tmp,
	}
	return
}
