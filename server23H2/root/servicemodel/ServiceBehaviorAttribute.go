// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceBehaviorAttribute struct
type ServiceBehaviorAttribute struct {
	*Behavior

	// Specifies how to use address to filter a message, whether to use Exact match, Prefix match or match Any.
	AddressFilterMode string

	// Indicates whether to automatically close a session when a client closes an output session.
	AutomaticSessionShutdown bool

	// Indicates whether a service supports one thread, multiple threads, or reentrant calls.
	ConcurrencyMode string

	// The name of the service configuration.
	ConfigurationName string

	// Specifies whether to process multiple messages concurrently at the dispatcher layer.
	EnsureOrderedDispatch bool

	// Specifies whether to send unknown serlialization data onto the wire.
	IgnoreExtensionDataObject bool

	// Specifies whether to include managed exception information in the detail of SOAP faults returned to the clients for debugging purposes.
	IncludeExceptionDetailInFaults bool

	// Specifies when a new service object is created.
	InstanceContextMode string

	// The maximum number of items allowed in a serialized object.
	MaxItemsInObjectGraph int32

	// The name attribute of the service in WSDL.
	Name string

	// The target namespace of the service in WSDL.
	Namespace string

	// Specifies whether the service object is recycled when the current transaction completes.
	ReleaseServiceInstanceOnTransactionComplete bool

	// Specifies whether pending transactions are completed when the current session closes.
	TransactionAutoCompleteOnSessionClose bool

	// Specifies the transaction isolation level.
	TransactionIsolationLevel string

	// The period within which a transaction must complete.
	TransactionTimeout string

	// Specifies whether to use the current synchronization context to choose the thread execution.
	UseSynchronizationContext bool

	// Specifies whether the system or the application enforces SOAP MustUnderstand header processing.
	ValidateMustUnderstand bool
}

func NewServiceBehaviorAttributeEx1(instance *cim.WmiInstance) (newInstance *ServiceBehaviorAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceBehaviorAttribute{
		Behavior: tmp,
	}
	return
}

func NewServiceBehaviorAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceBehaviorAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceBehaviorAttribute{
		Behavior: tmp,
	}
	return
}

// SetAddressFilterMode sets the value of AddressFilterMode for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyAddressFilterMode(value string) (err error) {
	return instance.SetProperty("AddressFilterMode", (value))
}

// GetAddressFilterMode gets the value of AddressFilterMode for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyAddressFilterMode() (value string, err error) {
	retValue, err := instance.GetProperty("AddressFilterMode")
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

// SetAutomaticSessionShutdown sets the value of AutomaticSessionShutdown for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyAutomaticSessionShutdown(value bool) (err error) {
	return instance.SetProperty("AutomaticSessionShutdown", (value))
}

// GetAutomaticSessionShutdown gets the value of AutomaticSessionShutdown for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyAutomaticSessionShutdown() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticSessionShutdown")
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

// SetConcurrencyMode sets the value of ConcurrencyMode for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyConcurrencyMode(value string) (err error) {
	return instance.SetProperty("ConcurrencyMode", (value))
}

// GetConcurrencyMode gets the value of ConcurrencyMode for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyConcurrencyMode() (value string, err error) {
	retValue, err := instance.GetProperty("ConcurrencyMode")
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

// SetConfigurationName sets the value of ConfigurationName for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyConfigurationName(value string) (err error) {
	return instance.SetProperty("ConfigurationName", (value))
}

// GetConfigurationName gets the value of ConfigurationName for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyConfigurationName() (value string, err error) {
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

// SetEnsureOrderedDispatch sets the value of EnsureOrderedDispatch for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyEnsureOrderedDispatch(value bool) (err error) {
	return instance.SetProperty("EnsureOrderedDispatch", (value))
}

// GetEnsureOrderedDispatch gets the value of EnsureOrderedDispatch for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyEnsureOrderedDispatch() (value bool, err error) {
	retValue, err := instance.GetProperty("EnsureOrderedDispatch")
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

// SetIgnoreExtensionDataObject sets the value of IgnoreExtensionDataObject for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyIgnoreExtensionDataObject(value bool) (err error) {
	return instance.SetProperty("IgnoreExtensionDataObject", (value))
}

// GetIgnoreExtensionDataObject gets the value of IgnoreExtensionDataObject for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyIgnoreExtensionDataObject() (value bool, err error) {
	retValue, err := instance.GetProperty("IgnoreExtensionDataObject")
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

// SetIncludeExceptionDetailInFaults sets the value of IncludeExceptionDetailInFaults for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyIncludeExceptionDetailInFaults(value bool) (err error) {
	return instance.SetProperty("IncludeExceptionDetailInFaults", (value))
}

// GetIncludeExceptionDetailInFaults gets the value of IncludeExceptionDetailInFaults for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyIncludeExceptionDetailInFaults() (value bool, err error) {
	retValue, err := instance.GetProperty("IncludeExceptionDetailInFaults")
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

// SetInstanceContextMode sets the value of InstanceContextMode for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyInstanceContextMode(value string) (err error) {
	return instance.SetProperty("InstanceContextMode", (value))
}

// GetInstanceContextMode gets the value of InstanceContextMode for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyInstanceContextMode() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceContextMode")
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

// SetMaxItemsInObjectGraph sets the value of MaxItemsInObjectGraph for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyMaxItemsInObjectGraph(value int32) (err error) {
	return instance.SetProperty("MaxItemsInObjectGraph", (value))
}

// GetMaxItemsInObjectGraph gets the value of MaxItemsInObjectGraph for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyMaxItemsInObjectGraph() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxItemsInObjectGraph")
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

// SetName sets the value of Name for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyName() (value string, err error) {
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
func (instance *ServiceBehaviorAttribute) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyNamespace() (value string, err error) {
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

// SetReleaseServiceInstanceOnTransactionComplete sets the value of ReleaseServiceInstanceOnTransactionComplete for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyReleaseServiceInstanceOnTransactionComplete(value bool) (err error) {
	return instance.SetProperty("ReleaseServiceInstanceOnTransactionComplete", (value))
}

// GetReleaseServiceInstanceOnTransactionComplete gets the value of ReleaseServiceInstanceOnTransactionComplete for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyReleaseServiceInstanceOnTransactionComplete() (value bool, err error) {
	retValue, err := instance.GetProperty("ReleaseServiceInstanceOnTransactionComplete")
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

// SetTransactionAutoCompleteOnSessionClose sets the value of TransactionAutoCompleteOnSessionClose for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyTransactionAutoCompleteOnSessionClose(value bool) (err error) {
	return instance.SetProperty("TransactionAutoCompleteOnSessionClose", (value))
}

// GetTransactionAutoCompleteOnSessionClose gets the value of TransactionAutoCompleteOnSessionClose for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyTransactionAutoCompleteOnSessionClose() (value bool, err error) {
	retValue, err := instance.GetProperty("TransactionAutoCompleteOnSessionClose")
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

// SetTransactionIsolationLevel sets the value of TransactionIsolationLevel for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyTransactionIsolationLevel(value string) (err error) {
	return instance.SetProperty("TransactionIsolationLevel", (value))
}

// GetTransactionIsolationLevel gets the value of TransactionIsolationLevel for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyTransactionIsolationLevel() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionIsolationLevel")
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

// SetTransactionTimeout sets the value of TransactionTimeout for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyTransactionTimeout(value string) (err error) {
	return instance.SetProperty("TransactionTimeout", (value))
}

// GetTransactionTimeout gets the value of TransactionTimeout for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyTransactionTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("TransactionTimeout")
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

// SetUseSynchronizationContext sets the value of UseSynchronizationContext for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyUseSynchronizationContext(value bool) (err error) {
	return instance.SetProperty("UseSynchronizationContext", (value))
}

// GetUseSynchronizationContext gets the value of UseSynchronizationContext for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyUseSynchronizationContext() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSynchronizationContext")
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

// SetValidateMustUnderstand sets the value of ValidateMustUnderstand for the instance
func (instance *ServiceBehaviorAttribute) SetPropertyValidateMustUnderstand(value bool) (err error) {
	return instance.SetProperty("ValidateMustUnderstand", (value))
}

// GetValidateMustUnderstand gets the value of ValidateMustUnderstand for the instance
func (instance *ServiceBehaviorAttribute) GetPropertyValidateMustUnderstand() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateMustUnderstand")
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
