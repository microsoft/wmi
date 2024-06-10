// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// Win32_RDFileAssociation struct
type Win32_RDFileAssociation struct {
	*cim.WmiInstance

	// Name of the extension (e.g. .txt)
	ExtName string

	// Contents of the icon for this file association
	IconContents []uint8

	// Index of the icon for this file association
	IconIndex int32

	// The path to the icon for this file association
	IconPath string

	// Whether this association is for a primary handler
	PrimaryHandler bool

	// Hint to help open documents with this file association
	ProgIdHint string
}

func NewWin32_RDFileAssociationEx1(instance *cim.WmiInstance) (newInstance *Win32_RDFileAssociation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_RDFileAssociation{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_RDFileAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_RDFileAssociation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_RDFileAssociation{
		WmiInstance: tmp,
	}
	return
}

// SetExtName sets the value of ExtName for the instance
func (instance *Win32_RDFileAssociation) SetPropertyExtName(value string) (err error) {
	return instance.SetProperty("ExtName", (value))
}

// GetExtName gets the value of ExtName for the instance
func (instance *Win32_RDFileAssociation) GetPropertyExtName() (value string, err error) {
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

// SetIconContents sets the value of IconContents for the instance
func (instance *Win32_RDFileAssociation) SetPropertyIconContents(value []uint8) (err error) {
	return instance.SetProperty("IconContents", (value))
}

// GetIconContents gets the value of IconContents for the instance
func (instance *Win32_RDFileAssociation) GetPropertyIconContents() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IconContents")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetIconIndex sets the value of IconIndex for the instance
func (instance *Win32_RDFileAssociation) SetPropertyIconIndex(value int32) (err error) {
	return instance.SetProperty("IconIndex", (value))
}

// GetIconIndex gets the value of IconIndex for the instance
func (instance *Win32_RDFileAssociation) GetPropertyIconIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("IconIndex")
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

// SetIconPath sets the value of IconPath for the instance
func (instance *Win32_RDFileAssociation) SetPropertyIconPath(value string) (err error) {
	return instance.SetProperty("IconPath", (value))
}

// GetIconPath gets the value of IconPath for the instance
func (instance *Win32_RDFileAssociation) GetPropertyIconPath() (value string, err error) {
	retValue, err := instance.GetProperty("IconPath")
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

// SetPrimaryHandler sets the value of PrimaryHandler for the instance
func (instance *Win32_RDFileAssociation) SetPropertyPrimaryHandler(value bool) (err error) {
	return instance.SetProperty("PrimaryHandler", (value))
}

// GetPrimaryHandler gets the value of PrimaryHandler for the instance
func (instance *Win32_RDFileAssociation) GetPropertyPrimaryHandler() (value bool, err error) {
	retValue, err := instance.GetProperty("PrimaryHandler")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetProgIdHint sets the value of ProgIdHint for the instance
func (instance *Win32_RDFileAssociation) SetPropertyProgIdHint(value string) (err error) {
	return instance.SetProperty("ProgIdHint", (value))
}

// GetProgIdHint gets the value of ProgIdHint for the instance
func (instance *Win32_RDFileAssociation) GetPropertyProgIdHint() (value string, err error) {
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
