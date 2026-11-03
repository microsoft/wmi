// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.SecurityCenter2
//////////////////////////////////////////////
package securitycenter2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WSC_CallerAMPPLData struct
type WSC_CallerAMPPLData struct {
	*cim.WmiInstance

	//
	displayName string

	//
	hresultCode int32

	//
	iniString string

	//
	instanceGuid string

	//
	pathToCaller string

	//
	pathToSignedProductExe string

	//
	statusCode int32
}

func NewWSC_CallerAMPPLDataEx1(instance *cim.WmiInstance) (newInstance *WSC_CallerAMPPLData, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &WSC_CallerAMPPLData{
		WmiInstance: tmp,
	}
	return
}

func NewWSC_CallerAMPPLDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSC_CallerAMPPLData, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSC_CallerAMPPLData{
		WmiInstance: tmp,
	}
	return
}

// SetdisplayName sets the value of displayName for the instance
func (instance *WSC_CallerAMPPLData) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", (value))
}

// GetdisplayName gets the value of displayName for the instance
func (instance *WSC_CallerAMPPLData) GetPropertydisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("displayName")
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

// SethresultCode sets the value of hresultCode for the instance
func (instance *WSC_CallerAMPPLData) SetPropertyhresultCode(value int32) (err error) {
	return instance.SetProperty("hresultCode", (value))
}

// GethresultCode gets the value of hresultCode for the instance
func (instance *WSC_CallerAMPPLData) GetPropertyhresultCode() (value int32, err error) {
	retValue, err := instance.GetProperty("hresultCode")
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

// SetiniString sets the value of iniString for the instance
func (instance *WSC_CallerAMPPLData) SetPropertyiniString(value string) (err error) {
	return instance.SetProperty("iniString", (value))
}

// GetiniString gets the value of iniString for the instance
func (instance *WSC_CallerAMPPLData) GetPropertyiniString() (value string, err error) {
	retValue, err := instance.GetProperty("iniString")
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

// SetinstanceGuid sets the value of instanceGuid for the instance
func (instance *WSC_CallerAMPPLData) SetPropertyinstanceGuid(value string) (err error) {
	return instance.SetProperty("instanceGuid", (value))
}

// GetinstanceGuid gets the value of instanceGuid for the instance
func (instance *WSC_CallerAMPPLData) GetPropertyinstanceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("instanceGuid")
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

// SetpathToCaller sets the value of pathToCaller for the instance
func (instance *WSC_CallerAMPPLData) SetPropertypathToCaller(value string) (err error) {
	return instance.SetProperty("pathToCaller", (value))
}

// GetpathToCaller gets the value of pathToCaller for the instance
func (instance *WSC_CallerAMPPLData) GetPropertypathToCaller() (value string, err error) {
	retValue, err := instance.GetProperty("pathToCaller")
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

// SetpathToSignedProductExe sets the value of pathToSignedProductExe for the instance
func (instance *WSC_CallerAMPPLData) SetPropertypathToSignedProductExe(value string) (err error) {
	return instance.SetProperty("pathToSignedProductExe", (value))
}

// GetpathToSignedProductExe gets the value of pathToSignedProductExe for the instance
func (instance *WSC_CallerAMPPLData) GetPropertypathToSignedProductExe() (value string, err error) {
	retValue, err := instance.GetProperty("pathToSignedProductExe")
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

// SetstatusCode sets the value of statusCode for the instance
func (instance *WSC_CallerAMPPLData) SetPropertystatusCode(value int32) (err error) {
	return instance.SetProperty("statusCode", (value))
}

// GetstatusCode gets the value of statusCode for the instance
func (instance *WSC_CallerAMPPLData) GetPropertystatusCode() (value int32, err error) {
	retValue, err := instance.GetProperty("statusCode")
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
