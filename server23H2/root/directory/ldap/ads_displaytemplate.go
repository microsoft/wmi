// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_displaytemplate struct
type ads_displaytemplate struct {
	*ds_top

	//
	DS_addressEntryDisplayTable Uint8Array

	//
	DS_addressEntryDisplayTableMSDOS Uint8Array

	//
	DS_helpData16 Uint8Array

	//
	DS_helpData32 Uint8Array

	//
	DS_helpFileName string

	//
	DS_originalDisplayTable Uint8Array

	//
	DS_originalDisplayTableMSDOS Uint8Array
}

func Newads_displaytemplateEx1(instance *cim.WmiInstance) (newInstance *ads_displaytemplate, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_displaytemplate{
		ds_top: tmp,
	}
	return
}

func Newads_displaytemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_displaytemplate, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_displaytemplate{
		ds_top: tmp,
	}
	return
}

// SetDS_addressEntryDisplayTable sets the value of DS_addressEntryDisplayTable for the instance
func (instance *ads_displaytemplate) SetPropertyDS_addressEntryDisplayTable(value Uint8Array) (err error) {
	return instance.SetProperty("DS_addressEntryDisplayTable", (value))
}

// GetDS_addressEntryDisplayTable gets the value of DS_addressEntryDisplayTable for the instance
func (instance *ads_displaytemplate) GetPropertyDS_addressEntryDisplayTable() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_addressEntryDisplayTable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_addressEntryDisplayTableMSDOS sets the value of DS_addressEntryDisplayTableMSDOS for the instance
func (instance *ads_displaytemplate) SetPropertyDS_addressEntryDisplayTableMSDOS(value Uint8Array) (err error) {
	return instance.SetProperty("DS_addressEntryDisplayTableMSDOS", (value))
}

// GetDS_addressEntryDisplayTableMSDOS gets the value of DS_addressEntryDisplayTableMSDOS for the instance
func (instance *ads_displaytemplate) GetPropertyDS_addressEntryDisplayTableMSDOS() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_addressEntryDisplayTableMSDOS")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_helpData16 sets the value of DS_helpData16 for the instance
func (instance *ads_displaytemplate) SetPropertyDS_helpData16(value Uint8Array) (err error) {
	return instance.SetProperty("DS_helpData16", (value))
}

// GetDS_helpData16 gets the value of DS_helpData16 for the instance
func (instance *ads_displaytemplate) GetPropertyDS_helpData16() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_helpData16")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_helpData32 sets the value of DS_helpData32 for the instance
func (instance *ads_displaytemplate) SetPropertyDS_helpData32(value Uint8Array) (err error) {
	return instance.SetProperty("DS_helpData32", (value))
}

// GetDS_helpData32 gets the value of DS_helpData32 for the instance
func (instance *ads_displaytemplate) GetPropertyDS_helpData32() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_helpData32")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_helpFileName sets the value of DS_helpFileName for the instance
func (instance *ads_displaytemplate) SetPropertyDS_helpFileName(value string) (err error) {
	return instance.SetProperty("DS_helpFileName", (value))
}

// GetDS_helpFileName gets the value of DS_helpFileName for the instance
func (instance *ads_displaytemplate) GetPropertyDS_helpFileName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_helpFileName")
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

// SetDS_originalDisplayTable sets the value of DS_originalDisplayTable for the instance
func (instance *ads_displaytemplate) SetPropertyDS_originalDisplayTable(value Uint8Array) (err error) {
	return instance.SetProperty("DS_originalDisplayTable", (value))
}

// GetDS_originalDisplayTable gets the value of DS_originalDisplayTable for the instance
func (instance *ads_displaytemplate) GetPropertyDS_originalDisplayTable() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_originalDisplayTable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_originalDisplayTableMSDOS sets the value of DS_originalDisplayTableMSDOS for the instance
func (instance *ads_displaytemplate) SetPropertyDS_originalDisplayTableMSDOS(value Uint8Array) (err error) {
	return instance.SetProperty("DS_originalDisplayTableMSDOS", (value))
}

// GetDS_originalDisplayTableMSDOS gets the value of DS_originalDisplayTableMSDOS for the instance
func (instance *ads_displaytemplate) GetPropertyDS_originalDisplayTableMSDOS() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_originalDisplayTableMSDOS")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}
