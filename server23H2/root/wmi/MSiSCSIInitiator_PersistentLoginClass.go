// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MSiSCSIInitiator_PersistentLoginClass struct
type MSiSCSIInitiator_PersistentLoginClass struct {
	*cim.WmiInstance

	//
	InitiatorInstance string

	//
	InitiatorPortNumber uint32

	//
	IsInformationalSession bool

	//
	LoginOptions MSiSCSIInitiator_TargetLoginOptions

	//
	Mappings MSiSCSIInitiator_TargetMappings

	//
	SecurityFlags uint64

	//
	TargetName string

	//
	TargetPortal MSiSCSIInitiator_Portal
}

func NewMSiSCSIInitiator_PersistentLoginClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_PersistentLoginClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_PersistentLoginClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_PersistentLoginClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_PersistentLoginClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_PersistentLoginClass{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorInstance sets the value of InitiatorInstance for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyInitiatorInstance(value string) (err error) {
	return instance.SetProperty("InitiatorInstance", (value))
}

// GetInitiatorInstance gets the value of InitiatorInstance for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyInitiatorInstance() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorInstance")
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

// SetInitiatorPortNumber sets the value of InitiatorPortNumber for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyInitiatorPortNumber(value uint32) (err error) {
	return instance.SetProperty("InitiatorPortNumber", (value))
}

// GetInitiatorPortNumber gets the value of InitiatorPortNumber for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyInitiatorPortNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitiatorPortNumber")
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

// SetIsInformationalSession sets the value of IsInformationalSession for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyIsInformationalSession(value bool) (err error) {
	return instance.SetProperty("IsInformationalSession", (value))
}

// GetIsInformationalSession gets the value of IsInformationalSession for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyIsInformationalSession() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInformationalSession")
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

// SetLoginOptions sets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyLoginOptions(value MSiSCSIInitiator_TargetLoginOptions) (err error) {
	return instance.SetProperty("LoginOptions", (value))
}

// GetLoginOptions gets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyLoginOptions() (value MSiSCSIInitiator_TargetLoginOptions, err error) {
	retValue, err := instance.GetProperty("LoginOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSiSCSIInitiator_TargetLoginOptions)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_TargetLoginOptions is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSiSCSIInitiator_TargetLoginOptions(valuetmp)

	return
}

// SetMappings sets the value of Mappings for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyMappings(value MSiSCSIInitiator_TargetMappings) (err error) {
	return instance.SetProperty("Mappings", (value))
}

// GetMappings gets the value of Mappings for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyMappings() (value MSiSCSIInitiator_TargetMappings, err error) {
	retValue, err := instance.GetProperty("Mappings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSiSCSIInitiator_TargetMappings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_TargetMappings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSiSCSIInitiator_TargetMappings(valuetmp)

	return
}

// SetSecurityFlags sets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertySecurityFlags(value uint64) (err error) {
	return instance.SetProperty("SecurityFlags", (value))
}

// GetSecurityFlags gets the value of SecurityFlags for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertySecurityFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("SecurityFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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

// SetTargetPortal sets the value of TargetPortal for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) SetPropertyTargetPortal(value MSiSCSIInitiator_Portal) (err error) {
	return instance.SetProperty("TargetPortal", (value))
}

// GetTargetPortal gets the value of TargetPortal for the instance
func (instance *MSiSCSIInitiator_PersistentLoginClass) GetPropertyTargetPortal() (value MSiSCSIInitiator_Portal, err error) {
	retValue, err := instance.GetProperty("TargetPortal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSiSCSIInitiator_Portal)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_Portal is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSiSCSIInitiator_Portal(valuetmp)

	return
}
