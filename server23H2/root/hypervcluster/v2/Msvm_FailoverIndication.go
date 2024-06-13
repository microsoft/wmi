// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_FailoverIndication struct
type Msvm_FailoverIndication struct {
	*CIM_ProcessIndication

	//
	FailoverType uint16

	//
	HostedElement string

	//
	HostedElementFormat uint16

	//
	HostingSystem string

	//
	HostingSystemFormat uint16

	//
	OtherFailoverType string

	//
	OtherHostedElementFormat string

	//
	OtherHostingSystemFormat string

	//
	OtherPerceivedSeverity string

	//
	OtherPreviousHostingSystemFormat string

	//
	PreviousHostingSystem string

	//
	PreviousHostingSystemFormat uint16
}

func NewMsvm_FailoverIndicationEx1(instance *cim.WmiInstance) (newInstance *Msvm_FailoverIndication, err error) {
	tmp, err := NewCIM_ProcessIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_FailoverIndication{
		CIM_ProcessIndication: tmp,
	}
	return
}

func NewMsvm_FailoverIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_FailoverIndication, err error) {
	tmp, err := NewCIM_ProcessIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_FailoverIndication{
		CIM_ProcessIndication: tmp,
	}
	return
}

// SetFailoverType sets the value of FailoverType for the instance
func (instance *Msvm_FailoverIndication) SetPropertyFailoverType(value uint16) (err error) {
	return instance.SetProperty("FailoverType", (value))
}

// GetFailoverType gets the value of FailoverType for the instance
func (instance *Msvm_FailoverIndication) GetPropertyFailoverType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailoverType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHostedElement sets the value of HostedElement for the instance
func (instance *Msvm_FailoverIndication) SetPropertyHostedElement(value string) (err error) {
	return instance.SetProperty("HostedElement", (value))
}

// GetHostedElement gets the value of HostedElement for the instance
func (instance *Msvm_FailoverIndication) GetPropertyHostedElement() (value string, err error) {
	retValue, err := instance.GetProperty("HostedElement")
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

// SetHostedElementFormat sets the value of HostedElementFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyHostedElementFormat(value uint16) (err error) {
	return instance.SetProperty("HostedElementFormat", (value))
}

// GetHostedElementFormat gets the value of HostedElementFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyHostedElementFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostedElementFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetHostingSystem sets the value of HostingSystem for the instance
func (instance *Msvm_FailoverIndication) SetPropertyHostingSystem(value string) (err error) {
	return instance.SetProperty("HostingSystem", (value))
}

// GetHostingSystem gets the value of HostingSystem for the instance
func (instance *Msvm_FailoverIndication) GetPropertyHostingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("HostingSystem")
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

// SetHostingSystemFormat sets the value of HostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyHostingSystemFormat(value uint16) (err error) {
	return instance.SetProperty("HostingSystemFormat", (value))
}

// GetHostingSystemFormat gets the value of HostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyHostingSystemFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostingSystemFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetOtherFailoverType sets the value of OtherFailoverType for the instance
func (instance *Msvm_FailoverIndication) SetPropertyOtherFailoverType(value string) (err error) {
	return instance.SetProperty("OtherFailoverType", (value))
}

// GetOtherFailoverType gets the value of OtherFailoverType for the instance
func (instance *Msvm_FailoverIndication) GetPropertyOtherFailoverType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherFailoverType")
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

// SetOtherHostedElementFormat sets the value of OtherHostedElementFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyOtherHostedElementFormat(value string) (err error) {
	return instance.SetProperty("OtherHostedElementFormat", (value))
}

// GetOtherHostedElementFormat gets the value of OtherHostedElementFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyOtherHostedElementFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHostedElementFormat")
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

// SetOtherHostingSystemFormat sets the value of OtherHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyOtherHostingSystemFormat(value string) (err error) {
	return instance.SetProperty("OtherHostingSystemFormat", (value))
}

// GetOtherHostingSystemFormat gets the value of OtherHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyOtherHostingSystemFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHostingSystemFormat")
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

// SetOtherPerceivedSeverity sets the value of OtherPerceivedSeverity for the instance
func (instance *Msvm_FailoverIndication) SetPropertyOtherPerceivedSeverity(value string) (err error) {
	return instance.SetProperty("OtherPerceivedSeverity", (value))
}

// GetOtherPerceivedSeverity gets the value of OtherPerceivedSeverity for the instance
func (instance *Msvm_FailoverIndication) GetPropertyOtherPerceivedSeverity() (value string, err error) {
	retValue, err := instance.GetProperty("OtherPerceivedSeverity")
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

// SetOtherPreviousHostingSystemFormat sets the value of OtherPreviousHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyOtherPreviousHostingSystemFormat(value string) (err error) {
	return instance.SetProperty("OtherPreviousHostingSystemFormat", (value))
}

// GetOtherPreviousHostingSystemFormat gets the value of OtherPreviousHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyOtherPreviousHostingSystemFormat() (value string, err error) {
	retValue, err := instance.GetProperty("OtherPreviousHostingSystemFormat")
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

// SetPreviousHostingSystem sets the value of PreviousHostingSystem for the instance
func (instance *Msvm_FailoverIndication) SetPropertyPreviousHostingSystem(value string) (err error) {
	return instance.SetProperty("PreviousHostingSystem", (value))
}

// GetPreviousHostingSystem gets the value of PreviousHostingSystem for the instance
func (instance *Msvm_FailoverIndication) GetPropertyPreviousHostingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("PreviousHostingSystem")
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

// SetPreviousHostingSystemFormat sets the value of PreviousHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) SetPropertyPreviousHostingSystemFormat(value uint16) (err error) {
	return instance.SetProperty("PreviousHostingSystemFormat", (value))
}

// GetPreviousHostingSystemFormat gets the value of PreviousHostingSystemFormat for the instance
func (instance *Msvm_FailoverIndication) GetPropertyPreviousHostingSystemFormat() (value uint16, err error) {
	retValue, err := instance.GetProperty("PreviousHostingSystemFormat")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
