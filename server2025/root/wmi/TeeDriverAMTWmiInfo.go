// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TeeDriverAMTWmiInfo struct
type TeeDriverAMTWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	// AMT Adt Rec Info
	AmtAdtRecInfo TeeDriverAmtAdtRecWmiInfo

	// AMT capability supported/not supported
	AmtCapablity bool

	// AMT globaly enabled/ disabled
	AmtEnabled bool

	// Prov state Wmi Info
	AmtProvStateInfo TeeDriverAmtProvStateNvarWmiInfo

	// AMT Trusted DNS Suffix
	AmtTrustedDNSSuffix []byte

	//
	InstanceName string
}

func NewTeeDriverAMTWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverAMTWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverAMTWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverAMTWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverAMTWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverAMTWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetAmtAdtRecInfo sets the value of AmtAdtRecInfo for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyAmtAdtRecInfo(value TeeDriverAmtAdtRecWmiInfo) (err error) {
	return instance.SetProperty("AmtAdtRecInfo", (value))
}

// GetAmtAdtRecInfo gets the value of AmtAdtRecInfo for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyAmtAdtRecInfo() (value TeeDriverAmtAdtRecWmiInfo, err error) {
	retValue, err := instance.GetProperty("AmtAdtRecInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TeeDriverAmtAdtRecWmiInfo)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TeeDriverAmtAdtRecWmiInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TeeDriverAmtAdtRecWmiInfo(valuetmp)

	return
}

// SetAmtCapablity sets the value of AmtCapablity for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyAmtCapablity(value bool) (err error) {
	return instance.SetProperty("AmtCapablity", (value))
}

// GetAmtCapablity gets the value of AmtCapablity for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyAmtCapablity() (value bool, err error) {
	retValue, err := instance.GetProperty("AmtCapablity")
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

// SetAmtEnabled sets the value of AmtEnabled for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyAmtEnabled(value bool) (err error) {
	return instance.SetProperty("AmtEnabled", (value))
}

// GetAmtEnabled gets the value of AmtEnabled for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyAmtEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AmtEnabled")
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

// SetAmtProvStateInfo sets the value of AmtProvStateInfo for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyAmtProvStateInfo(value TeeDriverAmtProvStateNvarWmiInfo) (err error) {
	return instance.SetProperty("AmtProvStateInfo", (value))
}

// GetAmtProvStateInfo gets the value of AmtProvStateInfo for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyAmtProvStateInfo() (value TeeDriverAmtProvStateNvarWmiInfo, err error) {
	retValue, err := instance.GetProperty("AmtProvStateInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TeeDriverAmtProvStateNvarWmiInfo)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TeeDriverAmtProvStateNvarWmiInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TeeDriverAmtProvStateNvarWmiInfo(valuetmp)

	return
}

// SetAmtTrustedDNSSuffix sets the value of AmtTrustedDNSSuffix for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyAmtTrustedDNSSuffix(value []byte) (err error) {
	return instance.SetProperty("AmtTrustedDNSSuffix", (value))
}

// GetAmtTrustedDNSSuffix gets the value of AmtTrustedDNSSuffix for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyAmtTrustedDNSSuffix() (value []byte, err error) {
	retValue, err := instance.GetProperty("AmtTrustedDNSSuffix")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(byte)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, byte(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverAMTWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverAMTWmiInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
