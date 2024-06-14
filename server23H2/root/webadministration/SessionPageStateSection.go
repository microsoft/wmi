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

// SessionPageStateSection struct
type SessionPageStateSection struct {
	*ConfigurationSection

	//
	HistorySize int32
}

func NewSessionPageStateSectionEx1(instance *cim.WmiInstance) (newInstance *SessionPageStateSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SessionPageStateSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewSessionPageStateSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SessionPageStateSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SessionPageStateSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetHistorySize sets the value of HistorySize for the instance
func (instance *SessionPageStateSection) SetPropertyHistorySize(value int32) (err error) {
	return instance.SetProperty("HistorySize", (value))
}

// GetHistorySize gets the value of HistorySize for the instance
func (instance *SessionPageStateSection) GetPropertyHistorySize() (value int32, err error) {
	retValue, err := instance.GetProperty("HistorySize")
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
