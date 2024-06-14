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

// ServerSideIncludeSection struct
type ServerSideIncludeSection struct {
	*ConfigurationSection

	//
	SsiExecDisable bool
}

func NewServerSideIncludeSectionEx1(instance *cim.WmiInstance) (newInstance *ServerSideIncludeSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServerSideIncludeSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewServerSideIncludeSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServerSideIncludeSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServerSideIncludeSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetSsiExecDisable sets the value of SsiExecDisable for the instance
func (instance *ServerSideIncludeSection) SetPropertySsiExecDisable(value bool) (err error) {
	return instance.SetProperty("SsiExecDisable", (value))
}

// GetSsiExecDisable gets the value of SsiExecDisable for the instance
func (instance *ServerSideIncludeSection) GetPropertySsiExecDisable() (value bool, err error) {
	retValue, err := instance.GetProperty("SsiExecDisable")
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
