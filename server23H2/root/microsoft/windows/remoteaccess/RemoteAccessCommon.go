// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessCommon struct
type RemoteAccessCommon struct {
	*RemoteAccessCore

	//
	DAStatus string

	//
	LoadBalancing string

	//
	RoutingStatus string

	//
	SstpProxyStatus string

	//
	UseHttp bool

	//
	VpnS2SStatus string

	//
	VpnStatus string
}

func NewRemoteAccessCommonEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessCommon, err error) {
	tmp, err := NewRemoteAccessCoreEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessCommon{
		RemoteAccessCore: tmp,
	}
	return
}

func NewRemoteAccessCommonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessCommon, err error) {
	tmp, err := NewRemoteAccessCoreEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessCommon{
		RemoteAccessCore: tmp,
	}
	return
}

// SetDAStatus sets the value of DAStatus for the instance
func (instance *RemoteAccessCommon) SetPropertyDAStatus(value string) (err error) {
	return instance.SetProperty("DAStatus", (value))
}

// GetDAStatus gets the value of DAStatus for the instance
func (instance *RemoteAccessCommon) GetPropertyDAStatus() (value string, err error) {
	retValue, err := instance.GetProperty("DAStatus")
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

// SetLoadBalancing sets the value of LoadBalancing for the instance
func (instance *RemoteAccessCommon) SetPropertyLoadBalancing(value string) (err error) {
	return instance.SetProperty("LoadBalancing", (value))
}

// GetLoadBalancing gets the value of LoadBalancing for the instance
func (instance *RemoteAccessCommon) GetPropertyLoadBalancing() (value string, err error) {
	retValue, err := instance.GetProperty("LoadBalancing")
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

// SetRoutingStatus sets the value of RoutingStatus for the instance
func (instance *RemoteAccessCommon) SetPropertyRoutingStatus(value string) (err error) {
	return instance.SetProperty("RoutingStatus", (value))
}

// GetRoutingStatus gets the value of RoutingStatus for the instance
func (instance *RemoteAccessCommon) GetPropertyRoutingStatus() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingStatus")
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

// SetSstpProxyStatus sets the value of SstpProxyStatus for the instance
func (instance *RemoteAccessCommon) SetPropertySstpProxyStatus(value string) (err error) {
	return instance.SetProperty("SstpProxyStatus", (value))
}

// GetSstpProxyStatus gets the value of SstpProxyStatus for the instance
func (instance *RemoteAccessCommon) GetPropertySstpProxyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("SstpProxyStatus")
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

// SetUseHttp sets the value of UseHttp for the instance
func (instance *RemoteAccessCommon) SetPropertyUseHttp(value bool) (err error) {
	return instance.SetProperty("UseHttp", (value))
}

// GetUseHttp gets the value of UseHttp for the instance
func (instance *RemoteAccessCommon) GetPropertyUseHttp() (value bool, err error) {
	retValue, err := instance.GetProperty("UseHttp")
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

// SetVpnS2SStatus sets the value of VpnS2SStatus for the instance
func (instance *RemoteAccessCommon) SetPropertyVpnS2SStatus(value string) (err error) {
	return instance.SetProperty("VpnS2SStatus", (value))
}

// GetVpnS2SStatus gets the value of VpnS2SStatus for the instance
func (instance *RemoteAccessCommon) GetPropertyVpnS2SStatus() (value string, err error) {
	retValue, err := instance.GetProperty("VpnS2SStatus")
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

// SetVpnStatus sets the value of VpnStatus for the instance
func (instance *RemoteAccessCommon) SetPropertyVpnStatus(value string) (err error) {
	return instance.SetProperty("VpnStatus", (value))
}

// GetVpnStatus gets the value of VpnStatus for the instance
func (instance *RemoteAccessCommon) GetPropertyVpnStatus() (value string, err error) {
	retValue, err := instance.GetProperty("VpnStatus")
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
