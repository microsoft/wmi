// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_DirectorySpecificationFile struct
type CIM_DirectorySpecificationFile struct {
	*cim.WmiInstance

	//
	DirectorySpecification CIM_DirectorySpecification

	//
	FileSpecification CIM_FileSpecification
}

func NewCIM_DirectorySpecificationFileEx1(instance *cim.WmiInstance) (newInstance *CIM_DirectorySpecificationFile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_DirectorySpecificationFile{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_DirectorySpecificationFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DirectorySpecificationFile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DirectorySpecificationFile{
		WmiInstance: tmp,
	}
	return
}

// SetDirectorySpecification sets the value of DirectorySpecification for the instance
func (instance *CIM_DirectorySpecificationFile) SetPropertyDirectorySpecification(value CIM_DirectorySpecification) (err error) {
	return instance.SetProperty("DirectorySpecification", value)
}

// GetDirectorySpecification gets the value of DirectorySpecification for the instance
func (instance *CIM_DirectorySpecificationFile) GetPropertyDirectorySpecification() (value CIM_DirectorySpecification, err error) {
	retValue, err := instance.GetProperty("DirectorySpecification")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_DirectorySpecification)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileSpecification sets the value of FileSpecification for the instance
func (instance *CIM_DirectorySpecificationFile) SetPropertyFileSpecification(value CIM_FileSpecification) (err error) {
	return instance.SetProperty("FileSpecification", value)
}

// GetFileSpecification gets the value of FileSpecification for the instance
func (instance *CIM_DirectorySpecificationFile) GetPropertyFileSpecification() (value CIM_FileSpecification, err error) {
	retValue, err := instance.GetProperty("FileSpecification")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_FileSpecification)
	if !ok {
		// TODO: Set an error
	}
	return
}
