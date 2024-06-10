// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_FwIcmpSettings struct
type HNet_FwIcmpSettings struct {
	*cim.WmiInstance

	//
	AllowInboundEchoRequest bool

	//
	AllowInboundMaskRequest bool

	//
	AllowInboundRouterRequest bool

	//
	AllowInboundTimestampRequest bool

	//
	AllowOutboundDestinationUnreachable bool

	//
	AllowOutboundParameterProblem bool

	//
	AllowOutboundSourceQuench bool

	//
	AllowOutboundTimeExceeded bool

	//
	AllowRedirect bool

	//
	Name string
}

func NewHNet_FwIcmpSettingsEx1(instance *cim.WmiInstance) (newInstance *HNet_FwIcmpSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_FwIcmpSettings{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_FwIcmpSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_FwIcmpSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_FwIcmpSettings{
		WmiInstance: tmp,
	}
	return
}

// SetAllowInboundEchoRequest sets the value of AllowInboundEchoRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundEchoRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundEchoRequest", (value))
}

// GetAllowInboundEchoRequest gets the value of AllowInboundEchoRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundEchoRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundEchoRequest")
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

// SetAllowInboundMaskRequest sets the value of AllowInboundMaskRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundMaskRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundMaskRequest", (value))
}

// GetAllowInboundMaskRequest gets the value of AllowInboundMaskRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundMaskRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundMaskRequest")
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

// SetAllowInboundRouterRequest sets the value of AllowInboundRouterRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundRouterRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundRouterRequest", (value))
}

// GetAllowInboundRouterRequest gets the value of AllowInboundRouterRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundRouterRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundRouterRequest")
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

// SetAllowInboundTimestampRequest sets the value of AllowInboundTimestampRequest for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowInboundTimestampRequest(value bool) (err error) {
	return instance.SetProperty("AllowInboundTimestampRequest", (value))
}

// GetAllowInboundTimestampRequest gets the value of AllowInboundTimestampRequest for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowInboundTimestampRequest() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInboundTimestampRequest")
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

// SetAllowOutboundDestinationUnreachable sets the value of AllowOutboundDestinationUnreachable for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundDestinationUnreachable(value bool) (err error) {
	return instance.SetProperty("AllowOutboundDestinationUnreachable", (value))
}

// GetAllowOutboundDestinationUnreachable gets the value of AllowOutboundDestinationUnreachable for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundDestinationUnreachable() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundDestinationUnreachable")
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

// SetAllowOutboundParameterProblem sets the value of AllowOutboundParameterProblem for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundParameterProblem(value bool) (err error) {
	return instance.SetProperty("AllowOutboundParameterProblem", (value))
}

// GetAllowOutboundParameterProblem gets the value of AllowOutboundParameterProblem for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundParameterProblem() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundParameterProblem")
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

// SetAllowOutboundSourceQuench sets the value of AllowOutboundSourceQuench for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundSourceQuench(value bool) (err error) {
	return instance.SetProperty("AllowOutboundSourceQuench", (value))
}

// GetAllowOutboundSourceQuench gets the value of AllowOutboundSourceQuench for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundSourceQuench() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundSourceQuench")
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

// SetAllowOutboundTimeExceeded sets the value of AllowOutboundTimeExceeded for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowOutboundTimeExceeded(value bool) (err error) {
	return instance.SetProperty("AllowOutboundTimeExceeded", (value))
}

// GetAllowOutboundTimeExceeded gets the value of AllowOutboundTimeExceeded for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowOutboundTimeExceeded() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowOutboundTimeExceeded")
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

// SetAllowRedirect sets the value of AllowRedirect for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyAllowRedirect(value bool) (err error) {
	return instance.SetProperty("AllowRedirect", (value))
}

// GetAllowRedirect gets the value of AllowRedirect for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyAllowRedirect() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowRedirect")
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

// SetName sets the value of Name for the instance
func (instance *HNet_FwIcmpSettings) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *HNet_FwIcmpSettings) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
