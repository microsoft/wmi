// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_LinkStateData struct
type MSNdis_LinkStateData struct {
	*MSNdis

	//
	AutoNegotiationFlags uint32

	//
	Header MSNdis_ObjectHeader

	//
	MediaConnectState uint32

	//
	MediaDuplexState uint32

	//
	PauseFunctions uint32

	//
	RcvLinkSpeed uint64

	//
	XmitLinkSpeed uint64
}

func NewMSNdis_LinkStateDataEx1(instance *cim.WmiInstance) (newInstance *MSNdis_LinkStateData, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_LinkStateData{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_LinkStateDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_LinkStateData, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_LinkStateData{
		MSNdis: tmp,
	}
	return
}

// SetAutoNegotiationFlags sets the value of AutoNegotiationFlags for the instance
func (instance *MSNdis_LinkStateData) SetPropertyAutoNegotiationFlags(value uint32) (err error) {
	return instance.SetProperty("AutoNegotiationFlags", (value))
}

// GetAutoNegotiationFlags gets the value of AutoNegotiationFlags for the instance
func (instance *MSNdis_LinkStateData) GetPropertyAutoNegotiationFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoNegotiationFlags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_LinkStateData) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_LinkStateData) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetMediaConnectState sets the value of MediaConnectState for the instance
func (instance *MSNdis_LinkStateData) SetPropertyMediaConnectState(value uint32) (err error) {
	return instance.SetProperty("MediaConnectState", (value))
}

// GetMediaConnectState gets the value of MediaConnectState for the instance
func (instance *MSNdis_LinkStateData) GetPropertyMediaConnectState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaConnectState")
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

// SetMediaDuplexState sets the value of MediaDuplexState for the instance
func (instance *MSNdis_LinkStateData) SetPropertyMediaDuplexState(value uint32) (err error) {
	return instance.SetProperty("MediaDuplexState", (value))
}

// GetMediaDuplexState gets the value of MediaDuplexState for the instance
func (instance *MSNdis_LinkStateData) GetPropertyMediaDuplexState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaDuplexState")
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

// SetPauseFunctions sets the value of PauseFunctions for the instance
func (instance *MSNdis_LinkStateData) SetPropertyPauseFunctions(value uint32) (err error) {
	return instance.SetProperty("PauseFunctions", (value))
}

// GetPauseFunctions gets the value of PauseFunctions for the instance
func (instance *MSNdis_LinkStateData) GetPropertyPauseFunctions() (value uint32, err error) {
	retValue, err := instance.GetProperty("PauseFunctions")
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

// SetRcvLinkSpeed sets the value of RcvLinkSpeed for the instance
func (instance *MSNdis_LinkStateData) SetPropertyRcvLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("RcvLinkSpeed", (value))
}

// GetRcvLinkSpeed gets the value of RcvLinkSpeed for the instance
func (instance *MSNdis_LinkStateData) GetPropertyRcvLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("RcvLinkSpeed")
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

// SetXmitLinkSpeed sets the value of XmitLinkSpeed for the instance
func (instance *MSNdis_LinkStateData) SetPropertyXmitLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("XmitLinkSpeed", (value))
}

// GetXmitLinkSpeed gets the value of XmitLinkSpeed for the instance
func (instance *MSNdis_LinkStateData) GetPropertyXmitLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("XmitLinkSpeed")
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
