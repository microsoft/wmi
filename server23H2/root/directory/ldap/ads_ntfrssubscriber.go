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

// ads_ntfrssubscriber struct
type ads_ntfrssubscriber struct {
	*ds_top

	//
	DS_fRSExtensions Uint8Array

	//
	DS_fRSFaultCondition string

	//
	DS_fRSFlags int32

	//
	DS_fRSMemberReference string

	//
	DS_fRSRootPath string

	//
	DS_fRSServiceCommand string

	//
	DS_fRSServiceCommandStatus string

	//
	DS_fRSStagingPath string

	//
	DS_fRSTimeLastCommand string

	//
	DS_fRSTimeLastConfigChange string

	//
	DS_fRSUpdateTimeout int32

	//
	DS_schedule Uint8Array
}

func Newads_ntfrssubscriberEx1(instance *cim.WmiInstance) (newInstance *ads_ntfrssubscriber, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrssubscriber{
		ds_top: tmp,
	}
	return
}

func Newads_ntfrssubscriberEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntfrssubscriber, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntfrssubscriber{
		ds_top: tmp,
	}
	return
}

// SetDS_fRSExtensions sets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSExtensions(value Uint8Array) (err error) {
	return instance.SetProperty("DS_fRSExtensions", (value))
}

// GetDS_fRSExtensions gets the value of DS_fRSExtensions for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSExtensions() (value Uint8Array, err error) {
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

// SetDS_fRSFaultCondition sets the value of DS_fRSFaultCondition for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSFaultCondition(value string) (err error) {
	return instance.SetProperty("DS_fRSFaultCondition", (value))
}

// GetDS_fRSFaultCondition gets the value of DS_fRSFaultCondition for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSFaultCondition() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSFaultCondition")
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

// SetDS_fRSFlags sets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSFlags(value int32) (err error) {
	return instance.SetProperty("DS_fRSFlags", (value))
}

// GetDS_fRSFlags gets the value of DS_fRSFlags for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSFlags() (value int32, err error) {
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

// SetDS_fRSMemberReference sets the value of DS_fRSMemberReference for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSMemberReference(value string) (err error) {
	return instance.SetProperty("DS_fRSMemberReference", (value))
}

// GetDS_fRSMemberReference gets the value of DS_fRSMemberReference for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSMemberReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSMemberReference")
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

// SetDS_fRSRootPath sets the value of DS_fRSRootPath for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSRootPath(value string) (err error) {
	return instance.SetProperty("DS_fRSRootPath", (value))
}

// GetDS_fRSRootPath gets the value of DS_fRSRootPath for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSRootPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSRootPath")
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

// SetDS_fRSServiceCommand sets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSServiceCommand(value string) (err error) {
	return instance.SetProperty("DS_fRSServiceCommand", (value))
}

// GetDS_fRSServiceCommand gets the value of DS_fRSServiceCommand for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSServiceCommand() (value string, err error) {
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

// SetDS_fRSServiceCommandStatus sets the value of DS_fRSServiceCommandStatus for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSServiceCommandStatus(value string) (err error) {
	return instance.SetProperty("DS_fRSServiceCommandStatus", (value))
}

// GetDS_fRSServiceCommandStatus gets the value of DS_fRSServiceCommandStatus for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSServiceCommandStatus() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSServiceCommandStatus")
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

// SetDS_fRSStagingPath sets the value of DS_fRSStagingPath for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSStagingPath(value string) (err error) {
	return instance.SetProperty("DS_fRSStagingPath", (value))
}

// GetDS_fRSStagingPath gets the value of DS_fRSStagingPath for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSStagingPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSStagingPath")
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

// SetDS_fRSTimeLastCommand sets the value of DS_fRSTimeLastCommand for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSTimeLastCommand(value string) (err error) {
	return instance.SetProperty("DS_fRSTimeLastCommand", (value))
}

// GetDS_fRSTimeLastCommand gets the value of DS_fRSTimeLastCommand for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSTimeLastCommand() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSTimeLastCommand")
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

// SetDS_fRSTimeLastConfigChange sets the value of DS_fRSTimeLastConfigChange for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSTimeLastConfigChange(value string) (err error) {
	return instance.SetProperty("DS_fRSTimeLastConfigChange", (value))
}

// GetDS_fRSTimeLastConfigChange gets the value of DS_fRSTimeLastConfigChange for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSTimeLastConfigChange() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSTimeLastConfigChange")
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
func (instance *ads_ntfrssubscriber) SetPropertyDS_fRSUpdateTimeout(value int32) (err error) {
	return instance.SetProperty("DS_fRSUpdateTimeout", (value))
}

// GetDS_fRSUpdateTimeout gets the value of DS_fRSUpdateTimeout for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_fRSUpdateTimeout() (value int32, err error) {
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

// SetDS_schedule sets the value of DS_schedule for the instance
func (instance *ads_ntfrssubscriber) SetPropertyDS_schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schedule", (value))
}

// GetDS_schedule gets the value of DS_schedule for the instance
func (instance *ads_ntfrssubscriber) GetPropertyDS_schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schedule")
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
