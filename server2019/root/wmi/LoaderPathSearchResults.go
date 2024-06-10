// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LoaderPathSearchResults struct
type LoaderPathSearchResults struct {
	*Image

	//
	AppDir string

	//
	Cwd string

	//
	DllDir string

	//
	DllLoadDir string

	//
	SearchInfo uint32
}

func NewLoaderPathSearchResultsEx1(instance *cim.WmiInstance) (newInstance *LoaderPathSearchResults, err error) {
	tmp, err := NewImageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LoaderPathSearchResults{
		Image: tmp,
	}
	return
}

func NewLoaderPathSearchResultsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LoaderPathSearchResults, err error) {
	tmp, err := NewImageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LoaderPathSearchResults{
		Image: tmp,
	}
	return
}

// SetAppDir sets the value of AppDir for the instance
func (instance *LoaderPathSearchResults) SetPropertyAppDir(value string) (err error) {
	return instance.SetProperty("AppDir", (value))
}

// GetAppDir gets the value of AppDir for the instance
func (instance *LoaderPathSearchResults) GetPropertyAppDir() (value string, err error) {
	retValue, err := instance.GetProperty("AppDir")
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

// SetCwd sets the value of Cwd for the instance
func (instance *LoaderPathSearchResults) SetPropertyCwd(value string) (err error) {
	return instance.SetProperty("Cwd", (value))
}

// GetCwd gets the value of Cwd for the instance
func (instance *LoaderPathSearchResults) GetPropertyCwd() (value string, err error) {
	retValue, err := instance.GetProperty("Cwd")
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

// SetDllDir sets the value of DllDir for the instance
func (instance *LoaderPathSearchResults) SetPropertyDllDir(value string) (err error) {
	return instance.SetProperty("DllDir", (value))
}

// GetDllDir gets the value of DllDir for the instance
func (instance *LoaderPathSearchResults) GetPropertyDllDir() (value string, err error) {
	retValue, err := instance.GetProperty("DllDir")
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

// SetDllLoadDir sets the value of DllLoadDir for the instance
func (instance *LoaderPathSearchResults) SetPropertyDllLoadDir(value string) (err error) {
	return instance.SetProperty("DllLoadDir", (value))
}

// GetDllLoadDir gets the value of DllLoadDir for the instance
func (instance *LoaderPathSearchResults) GetPropertyDllLoadDir() (value string, err error) {
	retValue, err := instance.GetProperty("DllLoadDir")
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

// SetSearchInfo sets the value of SearchInfo for the instance
func (instance *LoaderPathSearchResults) SetPropertySearchInfo(value uint32) (err error) {
	return instance.SetProperty("SearchInfo", (value))
}

// GetSearchInfo gets the value of SearchInfo for the instance
func (instance *LoaderPathSearchResults) GetPropertySearchInfo() (value uint32, err error) {
	retValue, err := instance.GetProperty("SearchInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
