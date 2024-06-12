// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ntfrsmember struct
type ads_ntfrsmember struct {
	*ds_top

	//
	DS_frsComputerReference string

	//
	DS_fRSControlDataCreation string

	//
	DS_fRSControlInboundBacklog string

	//
	DS_fRSControlOutboundBacklog string

	//
	DS_fRSExtensions Uint8Array

	//
	DS_fRSFlags int32

	//
	DS_fRSPartnerAuthLevel int32

	//
	DS_fRSRootSecurity Uint8Array

	//
	DS_fRSServiceCommand string

	//
	DS_fRSUpdateTimeout int32

	//
	DS_serverReference string
}

func Newads_ntfrsmemberEx1(instance *cim.WmiInstance) (newInstance *ads_ntfrsmember, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrsmember{
		ds_top: tmp,
	}
	return
}

func Newads_ntfrsmemberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntfrsmember, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrsmember{
		ds_top: tmp,
	}
	return
}

// SetDS_frsComputerReference sets the value of DS_frsComputerReference for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_frsComputerReference(value string) (err error) {
	return instance.SetProperty("DS_frsComputerReference", (value))
}

// GetDS_frsComputerReference gets the value of DS_frsComputerReference for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_frsComputerReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_frsComputerReference")
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

// SetDS_fRSControlDataCreation sets the value of DS_fRSControlDataCreation for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSControlDataCreation(value string) (err error) {
	return instance.SetProperty("DS_fRSControlDataCreation", (value))
}

// GetDS_fRSControlDataCreation gets the value of DS_fRSControlDataCreation for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSControlDataCreation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSControlDataCreation")
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

// SetDS_fRSControlInboundBacklog sets the value of DS_fRSControlInboundBacklog for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSControlInboundBacklog(value string) (err error) {
	return instance.SetProperty("DS_fRSControlInboundBacklog", (value))
}

// GetDS_fRSControlInboundBacklog gets the value of DS_fRSControlInboundBacklog for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSControlInboundBacklog() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSControlInboundBacklog")
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

// SetDS_fRSControlOutboundBacklog sets the value of DS_fRSControlOutboundBacklog for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSControlOutboundBacklog(value string) (err error) {
	return instance.SetProperty("DS_fRSControlOutboundBacklog", (value))
}

// GetDS_fRSControlOutboundBacklog gets the value of DS_fRSControlOutboundBacklog for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSControlOutboundBacklog() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSControlOutboundBacklog")
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

// SetDS_fRSExtensions sets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSExtensions(value Uint8Array) (err error) {
	return instance.SetProperty("DS_fRSExtensions", (value))
}

// GetDS_fRSExtensions gets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSExtensions() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_fRSExtensions")
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

// SetDS_fRSFlags sets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSFlags(value int32) (err error) {
	return instance.SetProperty("DS_fRSFlags", (value))
}

// GetDS_fRSFlags gets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_fRSFlags")
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

// SetDS_fRSPartnerAuthLevel sets the value of DS_fRSPartnerAuthLevel for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSPartnerAuthLevel(value int32) (err error) {
	return instance.SetProperty("DS_fRSPartnerAuthLevel", (value))
}

// GetDS_fRSPartnerAuthLevel gets the value of DS_fRSPartnerAuthLevel for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSPartnerAuthLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_fRSPartnerAuthLevel")
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

// SetDS_fRSRootSecurity sets the value of DS_fRSRootSecurity for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSRootSecurity(value Uint8Array) (err error) {
	return instance.SetProperty("DS_fRSRootSecurity", (value))
}

// GetDS_fRSRootSecurity gets the value of DS_fRSRootSecurity for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSRootSecurity() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_fRSRootSecurity")
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

// SetDS_fRSServiceCommand sets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSServiceCommand(value string) (err error) {
	return instance.SetProperty("DS_fRSServiceCommand", (value))
}

// GetDS_fRSServiceCommand gets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSServiceCommand() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSServiceCommand")
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

// SetDS_fRSUpdateTimeout sets the value of DS_fRSUpdateTimeout for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_fRSUpdateTimeout(value int32) (err error) {
	return instance.SetProperty("DS_fRSUpdateTimeout", (value))
}

// GetDS_fRSUpdateTimeout gets the value of DS_fRSUpdateTimeout for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_fRSUpdateTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_fRSUpdateTimeout")
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

// SetDS_serverReference sets the value of DS_serverReference for the instance
func (instance *ads_ntfrsmember) SetPropertyDS_serverReference(value string) (err error) {
	return instance.SetProperty("DS_serverReference", (value))
}

// GetDS_serverReference gets the value of DS_serverReference for the instance
func (instance *ads_ntfrsmember) GetPropertyDS_serverReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serverReference")
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
