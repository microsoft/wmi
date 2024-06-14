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

// DANetworkLocationServer struct
type DANetworkLocationServer struct {
	*cim.WmiInstance

	//
	Certificate []uint8

	//
	NlsLocation string

	//
	Reachability bool

	//
	Url string
}

func NewDANetworkLocationServerEx1(instance *cim.WmiInstance) (newInstance *DANetworkLocationServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DANetworkLocationServer{
		WmiInstance: tmp,
	}
	return
}

func NewDANetworkLocationServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DANetworkLocationServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DANetworkLocationServer{
		WmiInstance: tmp,
	}
	return
}

// SetCertificate sets the value of Certificate for the instance
func (instance *DANetworkLocationServer) SetPropertyCertificate(value []uint8) (err error) {
	return instance.SetProperty("Certificate", (value))
}

// GetCertificate gets the value of Certificate for the instance
func (instance *DANetworkLocationServer) GetPropertyCertificate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Certificate")
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

// SetNlsLocation sets the value of NlsLocation for the instance
func (instance *DANetworkLocationServer) SetPropertyNlsLocation(value string) (err error) {
	return instance.SetProperty("NlsLocation", (value))
}

// GetNlsLocation gets the value of NlsLocation for the instance
func (instance *DANetworkLocationServer) GetPropertyNlsLocation() (value string, err error) {
	retValue, err := instance.GetProperty("NlsLocation")
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

// SetReachability sets the value of Reachability for the instance
func (instance *DANetworkLocationServer) SetPropertyReachability(value bool) (err error) {
	return instance.SetProperty("Reachability", (value))
}

// GetReachability gets the value of Reachability for the instance
func (instance *DANetworkLocationServer) GetPropertyReachability() (value bool, err error) {
	retValue, err := instance.GetProperty("Reachability")
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

// SetUrl sets the value of Url for the instance
func (instance *DANetworkLocationServer) SetPropertyUrl(value string) (err error) {
	return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *DANetworkLocationServer) GetPropertyUrl() (value string, err error) {
	retValue, err := instance.GetProperty("Url")
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
