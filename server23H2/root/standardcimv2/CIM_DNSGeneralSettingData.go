// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DNSGeneralSettingData struct
type CIM_DNSGeneralSettingData struct {
	*CIM_IPAssignmentSettingData

	// 698
	AppendParentSuffixes bool

	// 697
	AppendPrimarySuffixes bool

	// 699
	DNSSuffixesToAppend []string
}

func NewCIM_DNSGeneralSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_DNSGeneralSettingData, err error) {
	tmp, err := NewCIM_IPAssignmentSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DNSGeneralSettingData{
		CIM_IPAssignmentSettingData: tmp,
	}
	return
}

func NewCIM_DNSGeneralSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DNSGeneralSettingData, err error) {
	tmp, err := NewCIM_IPAssignmentSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DNSGeneralSettingData{
		CIM_IPAssignmentSettingData: tmp,
	}
	return
}

// SetAppendParentSuffixes sets the value of AppendParentSuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyAppendParentSuffixes(value bool) (err error) {
	return instance.SetProperty("AppendParentSuffixes", (value))
}

// GetAppendParentSuffixes gets the value of AppendParentSuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyAppendParentSuffixes() (value bool, err error) {
	retValue, err := instance.GetProperty("AppendParentSuffixes")
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

// SetAppendPrimarySuffixes sets the value of AppendPrimarySuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyAppendPrimarySuffixes(value bool) (err error) {
	return instance.SetProperty("AppendPrimarySuffixes", (value))
}

// GetAppendPrimarySuffixes gets the value of AppendPrimarySuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyAppendPrimarySuffixes() (value bool, err error) {
	retValue, err := instance.GetProperty("AppendPrimarySuffixes")
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

// SetDNSSuffixesToAppend sets the value of DNSSuffixesToAppend for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyDNSSuffixesToAppend(value []string) (err error) {
	return instance.SetProperty("DNSSuffixesToAppend", (value))
}

// GetDNSSuffixesToAppend gets the value of DNSSuffixesToAppend for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyDNSSuffixesToAppend() (value []string, err error) {
	retValue, err := instance.GetProperty("DNSSuffixesToAppend")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
