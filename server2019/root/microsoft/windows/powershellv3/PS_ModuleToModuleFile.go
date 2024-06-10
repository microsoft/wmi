// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Powershellv3
//
// ////////////////////////////////////////////
package powershellv3

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_ModuleToModuleFile struct
type PS_ModuleToModuleFile struct {
	*CIM_Dependency
}

func NewPS_ModuleToModuleFileEx1(instance *cim.WmiInstance) (newInstance *PS_ModuleToModuleFile, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PS_ModuleToModuleFile{
		CIM_Dependency: tmp,
	}
	return
}

func NewPS_ModuleToModuleFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_ModuleToModuleFile, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_ModuleToModuleFile{
		CIM_Dependency: tmp,
	}
	return
}
