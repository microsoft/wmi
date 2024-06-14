// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ISCSI_LoginOptions struct
type ISCSI_LoginOptions struct {
	*cim.WmiInstance

	//
	AuthType uint32

	//
	DataDigest LoginOptions_DataDigest

	//
	DefaultTime2Retain uint32

	//
	DefaultTime2Wait uint32

	//
	HeaderDigest LoginOptions_HeaderDigest

	//
	InformationSpecified uint32

	//
	LoginFlags uint32

	//
	MaximumConnections uint32
}

func NewISCSI_LoginOptionsEx1(instance *cim.WmiInstance) (newInstance *ISCSI_LoginOptions, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_LoginOptions{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_LoginOptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_LoginOptions, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_LoginOptions{
		WmiInstance: tmp,
	}
	return
}

// SetAuthType sets the value of AuthType for the instance
func (instance *ISCSI_LoginOptions) SetPropertyAuthType(value uint32) (err error) {
	return instance.SetProperty("AuthType", (value))
}

// GetAuthType gets the value of AuthType for the instance
func (instance *ISCSI_LoginOptions) GetPropertyAuthType() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthType")
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

// SetDataDigest sets the value of DataDigest for the instance
func (instance *ISCSI_LoginOptions) SetPropertyDataDigest(value LoginOptions_DataDigest) (err error) {
	return instance.SetProperty("DataDigest", (value))
}

// GetDataDigest gets the value of DataDigest for the instance
func (instance *ISCSI_LoginOptions) GetPropertyDataDigest() (value LoginOptions_DataDigest, err error) {
	retValue, err := instance.GetProperty("DataDigest")
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

	value = LoginOptions_DataDigest(valuetmp)

	return
}

// SetDefaultTime2Retain sets the value of DefaultTime2Retain for the instance
func (instance *ISCSI_LoginOptions) SetPropertyDefaultTime2Retain(value uint32) (err error) {
	return instance.SetProperty("DefaultTime2Retain", (value))
}

// GetDefaultTime2Retain gets the value of DefaultTime2Retain for the instance
func (instance *ISCSI_LoginOptions) GetPropertyDefaultTime2Retain() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultTime2Retain")
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

// SetDefaultTime2Wait sets the value of DefaultTime2Wait for the instance
func (instance *ISCSI_LoginOptions) SetPropertyDefaultTime2Wait(value uint32) (err error) {
	return instance.SetProperty("DefaultTime2Wait", (value))
}

// GetDefaultTime2Wait gets the value of DefaultTime2Wait for the instance
func (instance *ISCSI_LoginOptions) GetPropertyDefaultTime2Wait() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultTime2Wait")
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

// SetHeaderDigest sets the value of HeaderDigest for the instance
func (instance *ISCSI_LoginOptions) SetPropertyHeaderDigest(value LoginOptions_HeaderDigest) (err error) {
	return instance.SetProperty("HeaderDigest", (value))
}

// GetHeaderDigest gets the value of HeaderDigest for the instance
func (instance *ISCSI_LoginOptions) GetPropertyHeaderDigest() (value LoginOptions_HeaderDigest, err error) {
	retValue, err := instance.GetProperty("HeaderDigest")
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

	value = LoginOptions_HeaderDigest(valuetmp)

	return
}

// SetInformationSpecified sets the value of InformationSpecified for the instance
func (instance *ISCSI_LoginOptions) SetPropertyInformationSpecified(value uint32) (err error) {
	return instance.SetProperty("InformationSpecified", (value))
}

// GetInformationSpecified gets the value of InformationSpecified for the instance
func (instance *ISCSI_LoginOptions) GetPropertyInformationSpecified() (value uint32, err error) {
	retValue, err := instance.GetProperty("InformationSpecified")
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

// SetLoginFlags sets the value of LoginFlags for the instance
func (instance *ISCSI_LoginOptions) SetPropertyLoginFlags(value uint32) (err error) {
	return instance.SetProperty("LoginFlags", (value))
}

// GetLoginFlags gets the value of LoginFlags for the instance
func (instance *ISCSI_LoginOptions) GetPropertyLoginFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoginFlags")
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

// SetMaximumConnections sets the value of MaximumConnections for the instance
func (instance *ISCSI_LoginOptions) SetPropertyMaximumConnections(value uint32) (err error) {
	return instance.SetProperty("MaximumConnections", (value))
}

// GetMaximumConnections gets the value of MaximumConnections for the instance
func (instance *ISCSI_LoginOptions) GetPropertyMaximumConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumConnections")
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
