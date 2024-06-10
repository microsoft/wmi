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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSIInitiator_TargetLoginOptions struct
type MSiSCSIInitiator_TargetLoginOptions struct {
	*cim.WmiInstance

	//
	AuthType uint32

	//
	DataDigest TargetLoginOptions_DataDigest

	//
	DefaultTime2Retain uint32

	//
	DefaultTime2Wait uint32

	//
	HeaderDigest TargetLoginOptions_HeaderDigest

	//
	InformationSpecified uint32

	//
	LoginFlags uint32

	//
	MaximumConnections uint32

	//
	Password []uint8

	//
	Username []uint8

	//
	Version uint32
}

func NewMSiSCSIInitiator_TargetLoginOptionsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_TargetLoginOptions, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetLoginOptions{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_TargetLoginOptionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_TargetLoginOptions, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_TargetLoginOptions{
		WmiInstance: tmp,
	}
	return
}

// SetAuthType sets the value of AuthType for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyAuthType(value uint32) (err error) {
	return instance.SetProperty("AuthType", (value))
}

// GetAuthType gets the value of AuthType for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyAuthType() (value uint32, err error) {
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
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyDataDigest(value TargetLoginOptions_DataDigest) (err error) {
	return instance.SetProperty("DataDigest", (value))
}

// GetDataDigest gets the value of DataDigest for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyDataDigest() (value TargetLoginOptions_DataDigest, err error) {
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

	value = TargetLoginOptions_DataDigest(valuetmp)

	return
}

// SetDefaultTime2Retain sets the value of DefaultTime2Retain for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyDefaultTime2Retain(value uint32) (err error) {
	return instance.SetProperty("DefaultTime2Retain", (value))
}

// GetDefaultTime2Retain gets the value of DefaultTime2Retain for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyDefaultTime2Retain() (value uint32, err error) {
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
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyDefaultTime2Wait(value uint32) (err error) {
	return instance.SetProperty("DefaultTime2Wait", (value))
}

// GetDefaultTime2Wait gets the value of DefaultTime2Wait for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyDefaultTime2Wait() (value uint32, err error) {
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
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyHeaderDigest(value TargetLoginOptions_HeaderDigest) (err error) {
	return instance.SetProperty("HeaderDigest", (value))
}

// GetHeaderDigest gets the value of HeaderDigest for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyHeaderDigest() (value TargetLoginOptions_HeaderDigest, err error) {
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

	value = TargetLoginOptions_HeaderDigest(valuetmp)

	return
}

// SetInformationSpecified sets the value of InformationSpecified for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyInformationSpecified(value uint32) (err error) {
	return instance.SetProperty("InformationSpecified", (value))
}

// GetInformationSpecified gets the value of InformationSpecified for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyInformationSpecified() (value uint32, err error) {
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
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyLoginFlags(value uint32) (err error) {
	return instance.SetProperty("LoginFlags", (value))
}

// GetLoginFlags gets the value of LoginFlags for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyLoginFlags() (value uint32, err error) {
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
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyMaximumConnections(value uint32) (err error) {
	return instance.SetProperty("MaximumConnections", (value))
}

// GetMaximumConnections gets the value of MaximumConnections for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyMaximumConnections() (value uint32, err error) {
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

// SetPassword sets the value of Password for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyPassword(value []uint8) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyPassword() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetUsername sets the value of Username for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyUsername(value []uint8) (err error) {
	return instance.SetProperty("Username", (value))
}

// GetUsername gets the value of Username for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyUsername() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Username")
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

// SetVersion sets the value of Version for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) SetPropertyVersion(value uint32) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MSiSCSIInitiator_TargetLoginOptions) GetPropertyVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("Version")
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
