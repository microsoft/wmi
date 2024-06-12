// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSBED60BE6_4FD5_4343_AC9D_FE4F9D9536E5.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEESCSecurityZoneSettings struct
type RSOP_IEESCSecurityZoneSettings struct {
	*RSOP_IESecurityZoneSettings

	//
	EscEnabled bool
}

func NewRSOP_IEESCSecurityZoneSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEESCSecurityZoneSettings, err error) {
	tmp, err := NewRSOP_IESecurityZoneSettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCSecurityZoneSettings{
		RSOP_IESecurityZoneSettings: tmp,
	}
	return
}

func NewRSOP_IEESCSecurityZoneSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEESCSecurityZoneSettings, err error) {
	tmp, err := NewRSOP_IESecurityZoneSettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCSecurityZoneSettings{
		RSOP_IESecurityZoneSettings: tmp,
	}
	return
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", (value))
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESCSecurityZoneSettings) GetPropertyEscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("EscEnabled")
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
