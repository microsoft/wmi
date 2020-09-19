// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DNSClientGlobalSetting struct
type MSFT_DNSClientGlobalSetting struct {
	*CIM_DNSGeneralSettingData

	// 705
	DevolutionLevel uint32

	// 704
	SuffixSearchList []string

	// 703
	UseDevolution bool

	// 702
	UseSuffixSearchList bool
}

func NewMSFT_DNSClientGlobalSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_DNSClientGlobalSetting, err error) {
	tmp, err := NewCIM_DNSGeneralSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientGlobalSetting{
		CIM_DNSGeneralSettingData: tmp,
	}
	return
}

func NewMSFT_DNSClientGlobalSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DNSClientGlobalSetting, err error) {
	tmp, err := NewCIM_DNSGeneralSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DNSClientGlobalSetting{
		CIM_DNSGeneralSettingData: tmp,
	}
	return
}

// SetDevolutionLevel sets the value of DevolutionLevel for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyDevolutionLevel(value uint32) (err error) {
	return instance.SetProperty("DevolutionLevel", (value))
}

// GetDevolutionLevel gets the value of DevolutionLevel for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyDevolutionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("DevolutionLevel")
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

// SetSuffixSearchList sets the value of SuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertySuffixSearchList(value []string) (err error) {
	return instance.SetProperty("SuffixSearchList", (value))
}

// GetSuffixSearchList gets the value of SuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertySuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("SuffixSearchList")
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

// SetUseDevolution sets the value of UseDevolution for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyUseDevolution(value bool) (err error) {
	return instance.SetProperty("UseDevolution", (value))
}

// GetUseDevolution gets the value of UseDevolution for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyUseDevolution() (value bool, err error) {
	retValue, err := instance.GetProperty("UseDevolution")
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

// SetUseSuffixSearchList sets the value of UseSuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyUseSuffixSearchList(value bool) (err error) {
	return instance.SetProperty("UseSuffixSearchList", (value))
}

// GetUseSuffixSearchList gets the value of UseSuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyUseSuffixSearchList() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSuffixSearchList")
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
