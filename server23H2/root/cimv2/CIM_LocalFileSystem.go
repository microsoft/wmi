// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_LocalFileSystem struct
type CIM_LocalFileSystem struct {
	*CIM_FileSystem
}

func NewCIM_LocalFileSystemEx1(instance *cim.WmiInstance) (newInstance *CIM_LocalFileSystem, err error) {
	tmp, err := NewCIM_FileSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_LocalFileSystem{
		CIM_FileSystem: tmp,
	}
	return
}

func NewCIM_LocalFileSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LocalFileSystem, err error) {
	tmp, err := NewCIM_FileSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LocalFileSystem{
		CIM_FileSystem: tmp,
	}
	return
}