// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// WmiFWVersion struct
type WmiFWVersion struct {
	*cim.WmiInstance

	// Build
	build uint32

	// Hotfix
	hotfix uint32

	// Major
	major uint32

	// Minor
	minor uint32

	// Server Segment Code
	serversegmentcode uint32
}

func NewWmiFWVersionEx1(instance *cim.WmiInstance) (newInstance *WmiFWVersion, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &WmiFWVersion{
		WmiInstance: tmp,
	}
	return
}

func NewWmiFWVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WmiFWVersion, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WmiFWVersion{
		WmiInstance: tmp,
	}
	return
}

// Setbuild sets the value of build for the instance
func (instance *WmiFWVersion) SetPropertybuild(value uint32) (err error) {
	return instance.SetProperty("build", (value))
}

// Getbuild gets the value of build for the instance
func (instance *WmiFWVersion) GetPropertybuild() (value uint32, err error) {
	retValue, err := instance.GetProperty("build")
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

// Sethotfix sets the value of hotfix for the instance
func (instance *WmiFWVersion) SetPropertyhotfix(value uint32) (err error) {
	return instance.SetProperty("hotfix", (value))
}

// Gethotfix gets the value of hotfix for the instance
func (instance *WmiFWVersion) GetPropertyhotfix() (value uint32, err error) {
	retValue, err := instance.GetProperty("hotfix")
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

// Setmajor sets the value of major for the instance
func (instance *WmiFWVersion) SetPropertymajor(value uint32) (err error) {
	return instance.SetProperty("major", (value))
}

// Getmajor gets the value of major for the instance
func (instance *WmiFWVersion) GetPropertymajor() (value uint32, err error) {
	retValue, err := instance.GetProperty("major")
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

// Setminor sets the value of minor for the instance
func (instance *WmiFWVersion) SetPropertyminor(value uint32) (err error) {
	return instance.SetProperty("minor", (value))
}

// Getminor gets the value of minor for the instance
func (instance *WmiFWVersion) GetPropertyminor() (value uint32, err error) {
	retValue, err := instance.GetProperty("minor")
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

// Setserversegmentcode sets the value of serversegmentcode for the instance
func (instance *WmiFWVersion) SetPropertyserversegmentcode(value uint32) (err error) {
	return instance.SetProperty("serversegmentcode", (value))
}

// Getserversegmentcode gets the value of serversegmentcode for the instance
func (instance *WmiFWVersion) GetPropertyserversegmentcode() (value uint32, err error) {
	retValue, err := instance.GetProperty("serversegmentcode")
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
