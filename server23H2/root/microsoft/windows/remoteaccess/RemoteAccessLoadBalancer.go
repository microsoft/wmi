// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessLoadBalancer struct
type RemoteAccessLoadBalancer struct {
	*cim.WmiInstance

	//
	InternalVirtualIPAddress []string

	//
	InternetVirtualIPAddress []string

	//
	NlbNodeStatus []RemoteAccessLoadBalancerNode

	//
	ThirdPartyLoadBalancer string
}

func NewRemoteAccessLoadBalancerEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessLoadBalancer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessLoadBalancer{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessLoadBalancerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessLoadBalancer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessLoadBalancer{
		WmiInstance: tmp,
	}
	return
}

// SetInternalVirtualIPAddress sets the value of InternalVirtualIPAddress for the instance
func (instance *RemoteAccessLoadBalancer) SetPropertyInternalVirtualIPAddress(value []string) (err error) {
	return instance.SetProperty("InternalVirtualIPAddress", (value))
}

// GetInternalVirtualIPAddress gets the value of InternalVirtualIPAddress for the instance
func (instance *RemoteAccessLoadBalancer) GetPropertyInternalVirtualIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("InternalVirtualIPAddress")
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

// SetInternetVirtualIPAddress sets the value of InternetVirtualIPAddress for the instance
func (instance *RemoteAccessLoadBalancer) SetPropertyInternetVirtualIPAddress(value []string) (err error) {
	return instance.SetProperty("InternetVirtualIPAddress", (value))
}

// GetInternetVirtualIPAddress gets the value of InternetVirtualIPAddress for the instance
func (instance *RemoteAccessLoadBalancer) GetPropertyInternetVirtualIPAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("InternetVirtualIPAddress")
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

// SetNlbNodeStatus sets the value of NlbNodeStatus for the instance
func (instance *RemoteAccessLoadBalancer) SetPropertyNlbNodeStatus(value []RemoteAccessLoadBalancerNode) (err error) {
	return instance.SetProperty("NlbNodeStatus", (value))
}

// GetNlbNodeStatus gets the value of NlbNodeStatus for the instance
func (instance *RemoteAccessLoadBalancer) GetPropertyNlbNodeStatus() (value []RemoteAccessLoadBalancerNode, err error) {
	retValue, err := instance.GetProperty("NlbNodeStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(RemoteAccessLoadBalancerNode)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " RemoteAccessLoadBalancerNode is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, RemoteAccessLoadBalancerNode(valuetmp))
	}

	return
}

// SetThirdPartyLoadBalancer sets the value of ThirdPartyLoadBalancer for the instance
func (instance *RemoteAccessLoadBalancer) SetPropertyThirdPartyLoadBalancer(value string) (err error) {
	return instance.SetProperty("ThirdPartyLoadBalancer", (value))
}

// GetThirdPartyLoadBalancer gets the value of ThirdPartyLoadBalancer for the instance
func (instance *RemoteAccessLoadBalancer) GetPropertyThirdPartyLoadBalancer() (value string, err error) {
	retValue, err := instance.GetProperty("ThirdPartyLoadBalancer")
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
