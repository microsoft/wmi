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

// RemoteAccessRadiusServer struct
type RemoteAccessRadiusServer struct {
	*cim.WmiInstance

	//
	AccountingOnOffMsg string

	//
	MsgAuthenticator string

	//
	Port uint32

	//
	Score uint32

	//
	ServerPurpose RemoteAccessRadiusServerPurpose

	//
	SharedSecret string

	//
	Timeout uint32
}

func NewRemoteAccessRadiusServerEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessRadiusServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusServer{
		WmiInstance: tmp,
	}
	return
}

func NewRemoteAccessRadiusServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessRadiusServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessRadiusServer{
		WmiInstance: tmp,
	}
	return
}

// SetAccountingOnOffMsg sets the value of AccountingOnOffMsg for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyAccountingOnOffMsg(value string) (err error) {
	return instance.SetProperty("AccountingOnOffMsg", (value))
}

// GetAccountingOnOffMsg gets the value of AccountingOnOffMsg for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyAccountingOnOffMsg() (value string, err error) {
	retValue, err := instance.GetProperty("AccountingOnOffMsg")
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

// SetMsgAuthenticator sets the value of MsgAuthenticator for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyMsgAuthenticator(value string) (err error) {
	return instance.SetProperty("MsgAuthenticator", (value))
}

// GetMsgAuthenticator gets the value of MsgAuthenticator for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyMsgAuthenticator() (value string, err error) {
	retValue, err := instance.GetProperty("MsgAuthenticator")
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

// SetPort sets the value of Port for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyPort(value uint32) (err error) {
	return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("Port")
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

// SetScore sets the value of Score for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyScore(value uint32) (err error) {
	return instance.SetProperty("Score", (value))
}

// GetScore gets the value of Score for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyScore() (value uint32, err error) {
	retValue, err := instance.GetProperty("Score")
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

// SetServerPurpose sets the value of ServerPurpose for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyServerPurpose(value RemoteAccessRadiusServerPurpose) (err error) {
	return instance.SetProperty("ServerPurpose", (value))
}

// GetServerPurpose gets the value of ServerPurpose for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyServerPurpose() (value RemoteAccessRadiusServerPurpose, err error) {
	retValue, err := instance.GetProperty("ServerPurpose")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RemoteAccessRadiusServerPurpose)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RemoteAccessRadiusServerPurpose is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RemoteAccessRadiusServerPurpose(valuetmp)

	return
}

// SetSharedSecret sets the value of SharedSecret for the instance
func (instance *RemoteAccessRadiusServer) SetPropertySharedSecret(value string) (err error) {
	return instance.SetProperty("SharedSecret", (value))
}

// GetSharedSecret gets the value of SharedSecret for the instance
func (instance *RemoteAccessRadiusServer) GetPropertySharedSecret() (value string, err error) {
	retValue, err := instance.GetProperty("SharedSecret")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *RemoteAccessRadiusServer) SetPropertyTimeout(value uint32) (err error) {
	return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *RemoteAccessRadiusServer) GetPropertyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timeout")
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
