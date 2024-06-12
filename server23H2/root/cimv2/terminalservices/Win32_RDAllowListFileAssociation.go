// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_RDAllowListFileAssociation struct
type Win32_RDAllowListFileAssociation struct {
	*cim.WmiInstance

	// Alias of this file association's RemoteApp
	AppAlias string

	// Name of the extension (e.g. .txt)
	ExtName string

	// Hint to help open documents with this file association
	ProgIdHint string
}

func NewWin32_RDAllowListFileAssociationEx1(instance *cim.WmiInstance) (newInstance *Win32_RDAllowListFileAssociation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_RDAllowListFileAssociation{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_RDAllowListFileAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_RDAllowListFileAssociation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_RDAllowListFileAssociation{
		WmiInstance: tmp,
	}
	return
}

// SetAppAlias sets the value of AppAlias for the instance
func (instance *Win32_RDAllowListFileAssociation) SetPropertyAppAlias(value string) (err error) {
	return instance.SetProperty("AppAlias", (value))
}

// GetAppAlias gets the value of AppAlias for the instance
func (instance *Win32_RDAllowListFileAssociation) GetPropertyAppAlias() (value string, err error) {
	retValue, err := instance.GetProperty("AppAlias")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetExtName sets the value of ExtName for the instance
func (instance *Win32_RDAllowListFileAssociation) SetPropertyExtName(value string) (err error) {
	return instance.SetProperty("ExtName", (value))
}

// GetExtName gets the value of ExtName for the instance
func (instance *Win32_RDAllowListFileAssociation) GetPropertyExtName() (value string, err error) {
	retValue, err := instance.GetProperty("ExtName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProgIdHint sets the value of ProgIdHint for the instance
func (instance *Win32_RDAllowListFileAssociation) SetPropertyProgIdHint(value string) (err error) {
	return instance.SetProperty("ProgIdHint", (value))
}

// GetProgIdHint gets the value of ProgIdHint for the instance
func (instance *Win32_RDAllowListFileAssociation) GetPropertyProgIdHint() (value string, err error) {
	retValue, err := instance.GetProperty("ProgIdHint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
