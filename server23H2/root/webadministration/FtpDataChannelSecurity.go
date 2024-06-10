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

// FtpDataChannelSecurity struct
type FtpDataChannelSecurity struct {
	*EmbeddedObject

	//
	MatchClientAddressForPasv bool

	//
	MatchClientAddressForPort bool
}

func NewFtpDataChannelSecurityEx1(instance *cim.WmiInstance) (newInstance *FtpDataChannelSecurity, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpDataChannelSecurity{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpDataChannelSecurityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpDataChannelSecurity, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpDataChannelSecurity{
		EmbeddedObject: tmp,
	}
	return
}

// SetMatchClientAddressForPasv sets the value of MatchClientAddressForPasv for the instance
func (instance *FtpDataChannelSecurity) SetPropertyMatchClientAddressForPasv(value bool) (err error) {
	return instance.SetProperty("MatchClientAddressForPasv", (value))
}

// GetMatchClientAddressForPasv gets the value of MatchClientAddressForPasv for the instance
func (instance *FtpDataChannelSecurity) GetPropertyMatchClientAddressForPasv() (value bool, err error) {
	retValue, err := instance.GetProperty("MatchClientAddressForPasv")
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

// SetMatchClientAddressForPort sets the value of MatchClientAddressForPort for the instance
func (instance *FtpDataChannelSecurity) SetPropertyMatchClientAddressForPort(value bool) (err error) {
	return instance.SetProperty("MatchClientAddressForPort", (value))
}

// GetMatchClientAddressForPort gets the value of MatchClientAddressForPort for the instance
func (instance *FtpDataChannelSecurity) GetPropertyMatchClientAddressForPort() (value bool, err error) {
	retValue, err := instance.GetProperty("MatchClientAddressForPort")
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
