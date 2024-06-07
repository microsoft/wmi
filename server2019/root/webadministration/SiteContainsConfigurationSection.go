// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SiteContainsConfigurationSection struct
type SiteContainsConfigurationSection struct { 
	*ObjectConfigurationAssociation
}

	func NewSiteContainsConfigurationSectionEx1(instance *cim.WmiInstance) (newInstance *SiteContainsConfigurationSection, err error) {tmp, err := NewObjectConfigurationAssociationEx1(instance)
		
	if err != nil { return }
	newInstance = &SiteContainsConfigurationSection {
	ObjectConfigurationAssociation: tmp,
	}
	return
	}
	

	func NewSiteContainsConfigurationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SiteContainsConfigurationSection, err error) {tmp, err := NewObjectConfigurationAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SiteContainsConfigurationSection {
	ObjectConfigurationAssociation: tmp,
	}
	return
	}
	

