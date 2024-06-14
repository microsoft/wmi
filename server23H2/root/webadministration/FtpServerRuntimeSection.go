// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpServerRuntimeSection struct
type FtpServerRuntimeSection struct {
	*ConfigurationSection

	//
	HostNameSupport FtpHostNameSupport
}

func NewFtpServerRuntimeSectionEx1(instance *cim.WmiInstance) (newInstance *FtpServerRuntimeSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpServerRuntimeSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewFtpServerRuntimeSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpServerRuntimeSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpServerRuntimeSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetHostNameSupport sets the value of HostNameSupport for the instance
func (instance *FtpServerRuntimeSection) SetPropertyHostNameSupport(value FtpHostNameSupport) (err error) {
	return instance.SetProperty("HostNameSupport", (value))
}

// GetHostNameSupport gets the value of HostNameSupport for the instance
func (instance *FtpServerRuntimeSection) GetPropertyHostNameSupport() (value FtpHostNameSupport, err error) {
	retValue, err := instance.GetProperty("HostNameSupport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpHostNameSupport)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpHostNameSupport is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpHostNameSupport(valuetmp)

	return
}
