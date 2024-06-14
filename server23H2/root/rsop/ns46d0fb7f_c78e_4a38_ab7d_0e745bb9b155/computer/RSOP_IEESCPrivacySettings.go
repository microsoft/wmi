// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.NS46D0FB7F_C78E_4A38_AB7D_0E745BB9B155.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEESCPrivacySettings struct
type RSOP_IEESCPrivacySettings struct {
	*RSOP_IEPrivacySettings

	//
	EscEnabled bool
}

func NewRSOP_IEESCPrivacySettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEESCPrivacySettings, err error) {
	tmp, err := NewRSOP_IEPrivacySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCPrivacySettings{
		RSOP_IEPrivacySettings: tmp,
	}
	return
}

func NewRSOP_IEESCPrivacySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEESCPrivacySettings, err error) {
	tmp, err := NewRSOP_IEPrivacySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEESCPrivacySettings{
		RSOP_IEPrivacySettings: tmp,
	}
	return
}

// SetEscEnabled sets the value of EscEnabled for the instance
func (instance *RSOP_IEESCPrivacySettings) SetPropertyEscEnabled(value bool) (err error) {
	return instance.SetProperty("EscEnabled", (value))
}

// GetEscEnabled gets the value of EscEnabled for the instance
func (instance *RSOP_IEESCPrivacySettings) GetPropertyEscEnabled() (value bool, err error) {
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
