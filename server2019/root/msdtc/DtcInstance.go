// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DtcInstance struct
type DtcInstance struct {
	*cim.WmiInstance

	//
	DtcName string

	//
	KtmRmEndpointCid string

	//
	OleTxEndpointCid string

	//
	Status string

	//
	UisEndpointCid string

	//
	VirtualServerName string

	//
	XAEndpointCid string
}

func NewDtcInstanceEx1(instance *cim.WmiInstance) (newInstance *DtcInstance, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcInstance{
		WmiInstance: tmp,
	}
	return
}

func NewDtcInstanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcInstance, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcInstance{
		WmiInstance: tmp,
	}
	return
}

// SetDtcName sets the value of DtcName for the instance
func (instance *DtcInstance) SetPropertyDtcName(value string) (err error) {
	return instance.SetProperty("DtcName", (value))
}

// GetDtcName gets the value of DtcName for the instance
func (instance *DtcInstance) GetPropertyDtcName() (value string, err error) {
	retValue, err := instance.GetProperty("DtcName")
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

// SetKtmRmEndpointCid sets the value of KtmRmEndpointCid for the instance
func (instance *DtcInstance) SetPropertyKtmRmEndpointCid(value string) (err error) {
	return instance.SetProperty("KtmRmEndpointCid", (value))
}

// GetKtmRmEndpointCid gets the value of KtmRmEndpointCid for the instance
func (instance *DtcInstance) GetPropertyKtmRmEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("KtmRmEndpointCid")
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

// SetOleTxEndpointCid sets the value of OleTxEndpointCid for the instance
func (instance *DtcInstance) SetPropertyOleTxEndpointCid(value string) (err error) {
	return instance.SetProperty("OleTxEndpointCid", (value))
}

// GetOleTxEndpointCid gets the value of OleTxEndpointCid for the instance
func (instance *DtcInstance) GetPropertyOleTxEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("OleTxEndpointCid")
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

// SetStatus sets the value of Status for the instance
func (instance *DtcInstance) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *DtcInstance) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetUisEndpointCid sets the value of UisEndpointCid for the instance
func (instance *DtcInstance) SetPropertyUisEndpointCid(value string) (err error) {
	return instance.SetProperty("UisEndpointCid", (value))
}

// GetUisEndpointCid gets the value of UisEndpointCid for the instance
func (instance *DtcInstance) GetPropertyUisEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("UisEndpointCid")
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

// SetVirtualServerName sets the value of VirtualServerName for the instance
func (instance *DtcInstance) SetPropertyVirtualServerName(value string) (err error) {
	return instance.SetProperty("VirtualServerName", (value))
}

// GetVirtualServerName gets the value of VirtualServerName for the instance
func (instance *DtcInstance) GetPropertyVirtualServerName() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualServerName")
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

// SetXAEndpointCid sets the value of XAEndpointCid for the instance
func (instance *DtcInstance) SetPropertyXAEndpointCid(value string) (err error) {
	return instance.SetProperty("XAEndpointCid", (value))
}

// GetXAEndpointCid gets the value of XAEndpointCid for the instance
func (instance *DtcInstance) GetPropertyXAEndpointCid() (value string, err error) {
	retValue, err := instance.GetProperty("XAEndpointCid")
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
