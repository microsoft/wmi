// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// WorkflowServiceBehavior struct
type WorkflowServiceBehavior struct {
	*Behavior

	// Specifies how to use address to filter a message, whether to use Exact match, Prefix match or match Any.
	AddressFilterMode string

	// The name of the service configuration.
	ConfigurationName string

	// Specifies whether to send unknown serlialization data onto the wire.
	IgnoreExtensionDataObject bool

	// Specifies whether to include managed exception information in the detail of SOAP faults returned to the clients for debugging purposes.
	IncludeExceptionDetailInFaults bool

	// The maximum number of items allowed in a serialized object.
	MaxItemsInObjectGraph int32

	// The name attribute of the service in WSDL.
	Name string

	// The target namespace of the service in WSDL.
	Namespace string

	// Specifies whether to use the current synchronization context to choose the thread execution.
	UseSynchronizationContext bool

	// Specifies whether the system or the application enforces SOAP MustUnderstand header processing.
	ValidateMustUnderstand bool

	// Specifies the path to XOML defined Workflow.
	WorkflowDefinitionPath string

	// Specifies the path to Rules file for a XOML defined Workflow.
	WorkflowRulesPath string

	// Specifies the Type of Workflow.
	WorkflowType string
}

func NewWorkflowServiceBehaviorEx1(instance *cim.WmiInstance) (newInstance *WorkflowServiceBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WorkflowServiceBehavior{
		Behavior: tmp,
	}
	return
}

func NewWorkflowServiceBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WorkflowServiceBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WorkflowServiceBehavior{
		Behavior: tmp,
	}
	return
}

// SetAddressFilterMode sets the value of AddressFilterMode for the instance
func (instance *WorkflowServiceBehavior) SetPropertyAddressFilterMode(value string) (err error) {
	return instance.SetProperty("AddressFilterMode", (value))
}

// GetAddressFilterMode gets the value of AddressFilterMode for the instance
func (instance *WorkflowServiceBehavior) GetPropertyAddressFilterMode() (value string, err error) {
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

// SetConfigurationName sets the value of ConfigurationName for the instance
func (instance *WorkflowServiceBehavior) SetPropertyConfigurationName(value string) (err error) {
	return instance.SetProperty("ConfigurationName", (value))
}

// GetConfigurationName gets the value of ConfigurationName for the instance
func (instance *WorkflowServiceBehavior) GetPropertyConfigurationName() (value string, err error) {
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

// SetIgnoreExtensionDataObject sets the value of IgnoreExtensionDataObject for the instance
func (instance *WorkflowServiceBehavior) SetPropertyIgnoreExtensionDataObject(value bool) (err error) {
	return instance.SetProperty("IgnoreExtensionDataObject", (value))
}

// GetIgnoreExtensionDataObject gets the value of IgnoreExtensionDataObject for the instance
func (instance *WorkflowServiceBehavior) GetPropertyIgnoreExtensionDataObject() (value bool, err error) {
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
func (instance *WorkflowServiceBehavior) SetPropertyIncludeExceptionDetailInFaults(value bool) (err error) {
	return instance.SetProperty("IncludeExceptionDetailInFaults", (value))
}

// GetIncludeExceptionDetailInFaults gets the value of IncludeExceptionDetailInFaults for the instance
func (instance *WorkflowServiceBehavior) GetPropertyIncludeExceptionDetailInFaults() (value bool, err error) {
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

// SetMaxItemsInObjectGraph sets the value of MaxItemsInObjectGraph for the instance
func (instance *WorkflowServiceBehavior) SetPropertyMaxItemsInObjectGraph(value int32) (err error) {
	return instance.SetProperty("MaxItemsInObjectGraph", (value))
}

// GetMaxItemsInObjectGraph gets the value of MaxItemsInObjectGraph for the instance
func (instance *WorkflowServiceBehavior) GetPropertyMaxItemsInObjectGraph() (value int32, err error) {
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
func (instance *WorkflowServiceBehavior) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *WorkflowServiceBehavior) GetPropertyName() (value string, err error) {
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
func (instance *WorkflowServiceBehavior) SetPropertyNamespace(value string) (err error) {
	return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *WorkflowServiceBehavior) GetPropertyNamespace() (value string, err error) {
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

// SetUseSynchronizationContext sets the value of UseSynchronizationContext for the instance
func (instance *WorkflowServiceBehavior) SetPropertyUseSynchronizationContext(value bool) (err error) {
	return instance.SetProperty("UseSynchronizationContext", (value))
}

// GetUseSynchronizationContext gets the value of UseSynchronizationContext for the instance
func (instance *WorkflowServiceBehavior) GetPropertyUseSynchronizationContext() (value bool, err error) {
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
func (instance *WorkflowServiceBehavior) SetPropertyValidateMustUnderstand(value bool) (err error) {
	return instance.SetProperty("ValidateMustUnderstand", (value))
}

// GetValidateMustUnderstand gets the value of ValidateMustUnderstand for the instance
func (instance *WorkflowServiceBehavior) GetPropertyValidateMustUnderstand() (value bool, err error) {
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

// SetWorkflowDefinitionPath sets the value of WorkflowDefinitionPath for the instance
func (instance *WorkflowServiceBehavior) SetPropertyWorkflowDefinitionPath(value string) (err error) {
	return instance.SetProperty("WorkflowDefinitionPath", (value))
}

// GetWorkflowDefinitionPath gets the value of WorkflowDefinitionPath for the instance
func (instance *WorkflowServiceBehavior) GetPropertyWorkflowDefinitionPath() (value string, err error) {
	retValue, err := instance.GetProperty("WorkflowDefinitionPath")
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

// SetWorkflowRulesPath sets the value of WorkflowRulesPath for the instance
func (instance *WorkflowServiceBehavior) SetPropertyWorkflowRulesPath(value string) (err error) {
	return instance.SetProperty("WorkflowRulesPath", (value))
}

// GetWorkflowRulesPath gets the value of WorkflowRulesPath for the instance
func (instance *WorkflowServiceBehavior) GetPropertyWorkflowRulesPath() (value string, err error) {
	retValue, err := instance.GetProperty("WorkflowRulesPath")
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

// SetWorkflowType sets the value of WorkflowType for the instance
func (instance *WorkflowServiceBehavior) SetPropertyWorkflowType(value string) (err error) {
	return instance.SetProperty("WorkflowType", (value))
}

// GetWorkflowType gets the value of WorkflowType for the instance
func (instance *WorkflowServiceBehavior) GetPropertyWorkflowType() (value string, err error) {
	retValue, err := instance.GetProperty("WorkflowType")
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
