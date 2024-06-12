// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Service struct
type Service struct {
	*cim.WmiInstance

	// The base addresses used by the service.
	BaseAddresses []string

	// The behaviors associated with this service.
	Behaviors []Behavior

	// ServiceElement_BehaviorConfiguration
	ConfigurationName string

	// Instance name of the instance of the performance counters of the service.
	CounterInstanceName string

	// Service name at the address.
	DistinguishedName string

	// The instance contexts for the extensions of the service instance.
	Extensions []string

	// The service metadata settings.
	Metadata []string

	// The unique name of this service.
	Name string

	// The namespace of the service.
	Namespace string

	// The time the service was opened.
	Opened string

	// The channels that are outgoing from the service instance.
	OutgoingChannels []Channel

	// The process id of the process that hosts the service.
	ProcessId int32
}

func NewServiceEx1(instance *cim.WmiInstance) (newInstance *Service, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Service{
		WmiInstance: tmp,
	}
	return
}

func NewServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Service, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Service{
		WmiInstance: tmp,
	}
	return
}

// SetBaseAddresses sets the value of BaseAddresses for the instance
func (instance *Service) SetPropertyBaseAddresses(value []string) (err error) {
	return instance.SetProperty("BaseAddresses", (value))
}

// GetBaseAddresses gets the value of BaseAddresses for the instance
func (instance *Service) GetPropertyBaseAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("BaseAddresses")
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

// SetBehaviors sets the value of Behaviors for the instance
func (instance *Service) SetPropertyBehaviors(value []Behavior) (err error) {
	return instance.SetProperty("Behaviors", (value))
}

// GetBehaviors gets the value of Behaviors for the instance
func (instance *Service) GetPropertyBehaviors() (value []Behavior, err error) {
	retValue, err := instance.GetProperty("Behaviors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Behavior)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Behavior is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Behavior(valuetmp))
	}

	return
}

// SetConfigurationName sets the value of ConfigurationName for the instance
func (instance *Service) SetPropertyConfigurationName(value string) (err error) {
	return instance.SetProperty("ConfigurationName", (value))
}

// GetConfigurationName gets the value of ConfigurationName for the instance
func (instance *Service) GetPropertyConfigurationName() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationName")
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

// SetCounterInstanceName sets the value of CounterInstanceName for the instance
func (instance *Service) SetPropertyCounterInstanceName(value string) (err error) {
	return instance.SetProperty("CounterInstanceName", (value))
}

// GetCounterInstanceName gets the value of CounterInstanceName for the instance
func (instance *Service) GetPropertyCounterInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("CounterInstanceName")
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

// SetDistinguishedName sets the value of DistinguishedName for the instance
func (instance *Service) SetPropertyDistinguishedName(value string) (err error) {
	return instance.SetProperty("DistinguishedName", (value))
}

// GetDistinguishedName gets the value of DistinguishedName for the instance
func (instance *Service) GetPropertyDistinguishedName() (value string, err error) {
	retValue, err := instance.GetProperty("DistinguishedName")
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

// SetExtensions sets the value of Extensions for the instance
func (instance *Service) SetPropertyExtensions(value []string) (err error) {
	return instance.SetProperty("Extensions", (value))
}

// GetExtensions gets the value of Extensions for the instance
func (instance *Service) GetPropertyExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("Extensions")
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

// SetMetadata sets the value of Metadata for the instance
func (instance *Service) SetPropertyMetadata(value []string) (err error) {
	return instance.SetProperty("Metadata", (value))
}

// GetMetadata gets the value of Metadata for the instance
func (instance *Service) GetPropertyMetadata() (value []string, err error) {
	retValue, err := instance.GetProperty("Metadata")
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

// SetName sets the value of Name for the instance
func (instance *Service) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Service) GetPropertyName() (value string, err error) {
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

// SetNamespace sets the value of Namespace for the instance
func (instance *Service) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *Service) GetPropertyNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("Namespace")
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

// SetOpened sets the value of Opened for the instance
func (instance *Service) SetPropertyOpened(value string) (err error) {
	return instance.SetProperty("Opened", (value))
}

// GetOpened gets the value of Opened for the instance
func (instance *Service) GetPropertyOpened() (value string, err error) {
	retValue, err := instance.GetProperty("Opened")
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

// SetOutgoingChannels sets the value of OutgoingChannels for the instance
func (instance *Service) SetPropertyOutgoingChannels(value []Channel) (err error) {
	return instance.SetProperty("OutgoingChannels", (value))
}

// GetOutgoingChannels gets the value of OutgoingChannels for the instance
func (instance *Service) GetPropertyOutgoingChannels() (value []Channel, err error) {
	retValue, err := instance.GetProperty("OutgoingChannels")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Channel)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Channel is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Channel(valuetmp))
	}

	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Service) SetPropertyProcessId(value int32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Service) GetPropertyProcessId() (value int32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
