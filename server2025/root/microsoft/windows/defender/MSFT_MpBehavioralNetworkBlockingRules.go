// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MpBehavioralNetworkBlockingRules struct
type MSFT_MpBehavioralNetworkBlockingRules struct {
	*BaseStatus

	//
	CreationTime string

	//
	Direction string

	//
	ExpirationTime string

	//
	FilterGUID string

	//
	IpAddress string

	//
	LocalPorts string

	//
	Protocol string

	//
	RemotePorts string

	//
	Type string
}

func NewMSFT_MpBehavioralNetworkBlockingRulesEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpBehavioralNetworkBlockingRules, err error) {
	tmp, err := NewBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpBehavioralNetworkBlockingRules{
		BaseStatus: tmp,
	}
	return
}

func NewMSFT_MpBehavioralNetworkBlockingRulesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpBehavioralNetworkBlockingRules, err error) {
	tmp, err := NewBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpBehavioralNetworkBlockingRules{
		BaseStatus: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetDirection sets the value of Direction for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyDirection(value string) (err error) {
	return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyDirection() (value string, err error) {
	retValue, err := instance.GetProperty("Direction")
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

// SetExpirationTime sets the value of ExpirationTime for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyExpirationTime(value string) (err error) {
	return instance.SetProperty("ExpirationTime", (value))
}

// GetExpirationTime gets the value of ExpirationTime for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyExpirationTime() (value string, err error) {
	retValue, err := instance.GetProperty("ExpirationTime")
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

// SetFilterGUID sets the value of FilterGUID for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyFilterGUID(value string) (err error) {
	return instance.SetProperty("FilterGUID", (value))
}

// GetFilterGUID gets the value of FilterGUID for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyFilterGUID() (value string, err error) {
	retValue, err := instance.GetProperty("FilterGUID")
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

// SetIpAddress sets the value of IpAddress for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyIpAddress(value string) (err error) {
	return instance.SetProperty("IpAddress", (value))
}

// GetIpAddress gets the value of IpAddress for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IpAddress")
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

// SetLocalPorts sets the value of LocalPorts for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyLocalPorts(value string) (err error) {
	return instance.SetProperty("LocalPorts", (value))
}

// GetLocalPorts gets the value of LocalPorts for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyLocalPorts() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPorts")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemotePorts sets the value of RemotePorts for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyRemotePorts(value string) (err error) {
	return instance.SetProperty("RemotePorts", (value))
}

// GetRemotePorts gets the value of RemotePorts for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyRemotePorts() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePorts")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_MpBehavioralNetworkBlockingRules) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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

//

// <param name="FilterGUID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpBehavioralNetworkBlockingRules) Remove( /* IN */ FilterGUID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", FilterGUID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
