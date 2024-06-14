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

// FtpUserIsolation struct
type FtpUserIsolation struct {
	*EmbeddedObject

	//
	ActiveDirectory FtpActiveDirectory

	//
	Mode int32
}

func NewFtpUserIsolationEx1(instance *cim.WmiInstance) (newInstance *FtpUserIsolation, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpUserIsolation{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpUserIsolationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpUserIsolation, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpUserIsolation{
		EmbeddedObject: tmp,
	}
	return
}

// SetActiveDirectory sets the value of ActiveDirectory for the instance
func (instance *FtpUserIsolation) SetPropertyActiveDirectory(value FtpActiveDirectory) (err error) {
	return instance.SetProperty("ActiveDirectory", (value))
}

// GetActiveDirectory gets the value of ActiveDirectory for the instance
func (instance *FtpUserIsolation) GetPropertyActiveDirectory() (value FtpActiveDirectory, err error) {
	retValue, err := instance.GetProperty("ActiveDirectory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpActiveDirectory)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpActiveDirectory is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpActiveDirectory(valuetmp)

	return
}

// SetMode sets the value of Mode for the instance
func (instance *FtpUserIsolation) SetPropertyMode(value int32) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *FtpUserIsolation) GetPropertyMode() (value int32, err error) {
	retValue, err := instance.GetProperty("Mode")
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
