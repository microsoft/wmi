// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// OMI_ResourceModuleManager struct
type OMI_ResourceModuleManager struct {
	*OMI_MetaConfigurationResource
}

func NewOMI_ResourceModuleManagerEx1(instance *cim.WmiInstance) (newInstance *OMI_ResourceModuleManager, err error) {
	tmp, err := NewOMI_MetaConfigurationResourceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OMI_ResourceModuleManager{
		OMI_MetaConfigurationResource: tmp,
	}
	return
}

func NewOMI_ResourceModuleManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OMI_ResourceModuleManager, err error) {
	tmp, err := NewOMI_MetaConfigurationResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OMI_ResourceModuleManager{
		OMI_MetaConfigurationResource: tmp,
	}
	return
}
