// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS9E4D1FDD_7A96_4826_963E_F60A751B9B15
//////////////////////////////////////////////
package ns9e4d1fdd_7a96_4826_963e_f60a751b9b15

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// __thisNAMESPACE struct
type __thisNAMESPACE struct {
	*__SystemClass

	//
	SECURITY_DESCRIPTOR []uint8
}

func New__thisNAMESPACEEx1(instance *cim.WmiInstance) (newInstance *__thisNAMESPACE, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__thisNAMESPACE{
		__SystemClass: tmp,
	}
	return
}

func New__thisNAMESPACEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__thisNAMESPACE, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__thisNAMESPACE{
		__SystemClass: tmp,
	}
	return
}

// SetSECURITY_DESCRIPTOR sets the value of SECURITY_DESCRIPTOR for the instance
func (instance *__thisNAMESPACE) SetPropertySECURITY_DESCRIPTOR(value []uint8) (err error) {
	return instance.SetProperty("SECURITY_DESCRIPTOR", (value))
}

// GetSECURITY_DESCRIPTOR gets the value of SECURITY_DESCRIPTOR for the instance
func (instance *__thisNAMESPACE) GetPropertySECURITY_DESCRIPTOR() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SECURITY_DESCRIPTOR")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}
