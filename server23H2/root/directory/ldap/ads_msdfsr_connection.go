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

// ads_msdfsr_connection struct
type ads_msdfsr_connection struct {
	*ds_top

	//
	DS_fromServer string

	//
	DS_msDFSR_DisablePacketPrivacy bool

	//
	DS_msDFSR_Enabled bool

	//
	DS_msDFSR_Extension Uint8Array

	//
	DS_msDFSR_Flags int32

	//
	DS_msDFSR_Keywords string

	//
	DS_msDFSR_Options int32

	//
	DS_msDFSR_Options2 int32

	//
	DS_msDFSR_Priority int32

	//
	DS_msDFSR_RdcEnabled bool

	//
	DS_msDFSR_RdcMinFileSizeInKb int64

	//
	DS_msDFSR_Schedule Uint8Array
}

func Newads_msdfsr_connectionEx1(instance *cim.WmiInstance) (newInstance *ads_msdfsr_connection, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_connection{
		ds_top: tmp,
	}
	return
}

func Newads_msdfsr_connectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfsr_connection, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfsr_connection{
		ds_top: tmp,
	}
	return
}

// SetDS_fromServer sets the value of DS_fromServer for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_fromServer(value string) (err error) {
	return instance.SetProperty("DS_fromServer", (value))
}

// GetDS_fromServer gets the value of DS_fromServer for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_fromServer() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fromServer")
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

// SetDS_msDFSR_DisablePacketPrivacy sets the value of DS_msDFSR_DisablePacketPrivacy for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_DisablePacketPrivacy(value bool) (err error) {
	return instance.SetProperty("DS_msDFSR_DisablePacketPrivacy", (value))
}

// GetDS_msDFSR_DisablePacketPrivacy gets the value of DS_msDFSR_DisablePacketPrivacy for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_DisablePacketPrivacy() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_DisablePacketPrivacy")
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

// SetDS_msDFSR_Enabled sets the value of DS_msDFSR_Enabled for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Enabled(value bool) (err error) {
	return instance.SetProperty("DS_msDFSR_Enabled", (value))
}

// GetDS_msDFSR_Enabled gets the value of DS_msDFSR_Enabled for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Enabled")
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

// SetDS_msDFSR_Extension sets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Extension(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Extension", (value))
}

// GetDS_msDFSR_Extension gets the value of DS_msDFSR_Extension for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Extension() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Extension")
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

// SetDS_msDFSR_Flags sets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Flags(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Flags", (value))
}

// GetDS_msDFSR_Flags gets the value of DS_msDFSR_Flags for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Flags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Flags")
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

// SetDS_msDFSR_Keywords sets the value of DS_msDFSR_Keywords for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Keywords(value string) (err error) {
	return instance.SetProperty("DS_msDFSR_Keywords", (value))
}

// GetDS_msDFSR_Keywords gets the value of DS_msDFSR_Keywords for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Keywords() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Keywords")
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

// SetDS_msDFSR_Options sets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Options(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options", (value))
}

// GetDS_msDFSR_Options gets the value of DS_msDFSR_Options for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Options")
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

// SetDS_msDFSR_Options2 sets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Options2(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Options2", (value))
}

// GetDS_msDFSR_Options2 gets the value of DS_msDFSR_Options2 for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Options2() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Options2")
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

// SetDS_msDFSR_Priority sets the value of DS_msDFSR_Priority for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Priority(value int32) (err error) {
	return instance.SetProperty("DS_msDFSR_Priority", (value))
}

// GetDS_msDFSR_Priority gets the value of DS_msDFSR_Priority for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Priority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Priority")
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

// SetDS_msDFSR_RdcEnabled sets the value of DS_msDFSR_RdcEnabled for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_RdcEnabled(value bool) (err error) {
	return instance.SetProperty("DS_msDFSR_RdcEnabled", (value))
}

// GetDS_msDFSR_RdcEnabled gets the value of DS_msDFSR_RdcEnabled for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_RdcEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_RdcEnabled")
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

// SetDS_msDFSR_RdcMinFileSizeInKb sets the value of DS_msDFSR_RdcMinFileSizeInKb for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_RdcMinFileSizeInKb(value int64) (err error) {
	return instance.SetProperty("DS_msDFSR_RdcMinFileSizeInKb", (value))
}

// GetDS_msDFSR_RdcMinFileSizeInKb gets the value of DS_msDFSR_RdcMinFileSizeInKb for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_RdcMinFileSizeInKb() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_RdcMinFileSizeInKb")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDFSR_Schedule sets the value of DS_msDFSR_Schedule for the instance
func (instance *ads_msdfsr_connection) SetPropertyDS_msDFSR_Schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFSR_Schedule", (value))
}

// GetDS_msDFSR_Schedule gets the value of DS_msDFSR_Schedule for the instance
func (instance *ads_msdfsr_connection) GetPropertyDS_msDFSR_Schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_Schedule")
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
