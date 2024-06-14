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
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpDirectoryBrowse struct
type FtpDirectoryBrowse struct {
	*EmbeddedObject

	//
	ShowFlags int32

	//
	VirtualDirectoryTimeout int32
}

func NewFtpDirectoryBrowseEx1(instance *cim.WmiInstance) (newInstance *FtpDirectoryBrowse, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpDirectoryBrowse{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpDirectoryBrowseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpDirectoryBrowse, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpDirectoryBrowse{
		EmbeddedObject: tmp,
	}
	return
}

// SetShowFlags sets the value of ShowFlags for the instance
func (instance *FtpDirectoryBrowse) SetPropertyShowFlags(value int32) (err error) {
	return instance.SetProperty("ShowFlags", (value))
}

// GetShowFlags gets the value of ShowFlags for the instance
func (instance *FtpDirectoryBrowse) GetPropertyShowFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("ShowFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetVirtualDirectoryTimeout sets the value of VirtualDirectoryTimeout for the instance
func (instance *FtpDirectoryBrowse) SetPropertyVirtualDirectoryTimeout(value int32) (err error) {
	return instance.SetProperty("VirtualDirectoryTimeout", (value))
}

// GetVirtualDirectoryTimeout gets the value of VirtualDirectoryTimeout for the instance
func (instance *FtpDirectoryBrowse) GetPropertyVirtualDirectoryTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("VirtualDirectoryTimeout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
