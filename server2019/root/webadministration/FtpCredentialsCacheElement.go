// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpCredentialsCacheElement struct
type FtpCredentialsCacheElement struct {
	*CollectionElement

	//
	Enabled bool

	//
	FlushInterval uint32
}

func NewFtpCredentialsCacheElementEx1(instance *cim.WmiInstance) (newInstance *FtpCredentialsCacheElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCredentialsCacheElement{
		CollectionElement: tmp,
	}
	return
}

func NewFtpCredentialsCacheElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCredentialsCacheElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCredentialsCacheElement{
		CollectionElement: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *FtpCredentialsCacheElement) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *FtpCredentialsCacheElement) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetFlushInterval sets the value of FlushInterval for the instance
func (instance *FtpCredentialsCacheElement) SetPropertyFlushInterval(value uint32) (err error) {
	return instance.SetProperty("FlushInterval", (value))
}

// GetFlushInterval gets the value of FlushInterval for the instance
func (instance *FtpCredentialsCacheElement) GetPropertyFlushInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("FlushInterval")
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
