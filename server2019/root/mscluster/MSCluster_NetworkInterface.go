// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_NetworkInterface struct
type MSCluster_NetworkInterface struct {
	*CIM_LogicalDevice

	//
	Adapter string

	//
	AdapterId string

	//
	Address string

	//
	Characteristics uint32

	//
	DhcpEnabled bool

	//
	Flags uint32

	//
	Id string

	//
	IPv4Addresses []string

	//
	IPv6Addresses []string

	//
	Network string

	//
	Node string

	//
	PrivateProperties MSCluster_Property

	//
	State uint32
}

func NewMSCluster_NetworkInterfaceEx1(instance *cim.WmiInstance) (newInstance *MSCluster_NetworkInterface, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NetworkInterface{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMSCluster_NetworkInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_NetworkInterface, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_NetworkInterface{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAdapter sets the value of Adapter for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyAdapter(value string) (err error) {
	return instance.SetProperty("Adapter", (value))
}

// GetAdapter gets the value of Adapter for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyAdapter() (value string, err error) {
	retValue, err := instance.GetProperty("Adapter")
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

// SetAdapterId sets the value of AdapterId for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyAdapterId(value string) (err error) {
	return instance.SetProperty("AdapterId", (value))
}

// GetAdapterId gets the value of AdapterId for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyAdapterId() (value string, err error) {
	retValue, err := instance.GetProperty("AdapterId")
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

// SetAddress sets the value of Address for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyCharacteristics(value uint32) (err error) {
	return instance.SetProperty("Characteristics", (value))
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyCharacteristics() (value uint32, err error) {
	retValue, err := instance.GetProperty("Characteristics")
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

// SetDhcpEnabled sets the value of DhcpEnabled for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyDhcpEnabled(value bool) (err error) {
	return instance.SetProperty("DhcpEnabled", (value))
}

// GetDhcpEnabled gets the value of DhcpEnabled for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyDhcpEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DhcpEnabled")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetId sets the value of Id for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetIPv4Addresses sets the value of IPv4Addresses for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyIPv4Addresses(value []string) (err error) {
	return instance.SetProperty("IPv4Addresses", (value))
}

// GetIPv4Addresses gets the value of IPv4Addresses for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyIPv4Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4Addresses")
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

// SetIPv6Addresses sets the value of IPv6Addresses for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyIPv6Addresses(value []string) (err error) {
	return instance.SetProperty("IPv6Addresses", (value))
}

// GetIPv6Addresses gets the value of IPv6Addresses for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyIPv6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6Addresses")
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

// SetNetwork sets the value of Network for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyNetwork(value string) (err error) {
	return instance.SetProperty("Network", (value))
}

// GetNetwork gets the value of Network for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("Network")
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

// SetNode sets the value of Node for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyNode(value string) (err error) {
	return instance.SetProperty("Node", (value))
}

// GetNode gets the value of Node for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyNode() (value string, err error) {
	retValue, err := instance.GetProperty("Node")
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

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", (value))
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
	retValue, err := instance.GetProperty("PrivateProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSCluster_Property)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSCluster_Property is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSCluster_Property(valuetmp)

	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_NetworkInterface) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSCluster_NetworkInterface) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>
// <param name="Reason" type="string "></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_NetworkInterface) ExecuteNetworkInterfaceControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32,
	/* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethod("ExecuteNetworkInterfaceControl", ControlCode, InputBuffer, Reason)
	if err != nil {
		return
	}
	return

}
