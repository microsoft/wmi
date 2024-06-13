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

// ads_grouppolicycontainer struct
type ads_grouppolicycontainer struct {
	*ads_container

	//
	DS_gPCFileSysPath string

	//
	DS_gPCFunctionalityVersion int32

	//
	DS_gPCMachineExtensionNames string

	//
	DS_gPCUserExtensionNames string

	//
	DS_gPCWQLFilter string

	//
	DS_versionNumber int32
}

func Newads_grouppolicycontainerEx1(instance *cim.WmiInstance) (newInstance *ads_grouppolicycontainer, err error) {
	tmp, err := Newads_containerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_grouppolicycontainer{
		ads_container: tmp,
	}
	return
}

func Newads_grouppolicycontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_grouppolicycontainer, err error) {
	tmp, err := Newads_containerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_grouppolicycontainer{
		ads_container: tmp,
	}
	return
}

// SetDS_gPCFileSysPath sets the value of DS_gPCFileSysPath for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_gPCFileSysPath(value string) (err error) {
	return instance.SetProperty("DS_gPCFileSysPath", (value))
}

// GetDS_gPCFileSysPath gets the value of DS_gPCFileSysPath for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_gPCFileSysPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPCFileSysPath")
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

// SetDS_gPCFunctionalityVersion sets the value of DS_gPCFunctionalityVersion for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_gPCFunctionalityVersion(value int32) (err error) {
	return instance.SetProperty("DS_gPCFunctionalityVersion", (value))
}

// GetDS_gPCFunctionalityVersion gets the value of DS_gPCFunctionalityVersion for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_gPCFunctionalityVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_gPCFunctionalityVersion")
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

// SetDS_gPCMachineExtensionNames sets the value of DS_gPCMachineExtensionNames for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_gPCMachineExtensionNames(value string) (err error) {
	return instance.SetProperty("DS_gPCMachineExtensionNames", (value))
}

// GetDS_gPCMachineExtensionNames gets the value of DS_gPCMachineExtensionNames for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_gPCMachineExtensionNames() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPCMachineExtensionNames")
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

// SetDS_gPCUserExtensionNames sets the value of DS_gPCUserExtensionNames for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_gPCUserExtensionNames(value string) (err error) {
	return instance.SetProperty("DS_gPCUserExtensionNames", (value))
}

// GetDS_gPCUserExtensionNames gets the value of DS_gPCUserExtensionNames for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_gPCUserExtensionNames() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPCUserExtensionNames")
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

// SetDS_gPCWQLFilter sets the value of DS_gPCWQLFilter for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_gPCWQLFilter(value string) (err error) {
	return instance.SetProperty("DS_gPCWQLFilter", (value))
}

// GetDS_gPCWQLFilter gets the value of DS_gPCWQLFilter for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_gPCWQLFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPCWQLFilter")
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

// SetDS_versionNumber sets the value of DS_versionNumber for the instance
func (instance *ads_grouppolicycontainer) SetPropertyDS_versionNumber(value int32) (err error) {
	return instance.SetProperty("DS_versionNumber", (value))
}

// GetDS_versionNumber gets the value of DS_versionNumber for the instance
func (instance *ads_grouppolicycontainer) GetPropertyDS_versionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumber")
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
