// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Service struct
type CIM_Service struct {
	*CIM_EnabledLogicalElement

	// CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	// A string that provides information on how the primary owner of the Service can be reached (for example, phone number, e-mail address, and so on).
	PrimaryOwnerContact string

	// The name of the primary owner for the service, if one is defined. The primary owner is the initial support contact for the Service.
	PrimaryOwnerName string

	// Started is a Boolean that indicates whether the Service has been started (TRUE), or stopped (FALSE).
	Started bool

	// Note: The use of this element is deprecated in lieu of the EnabledDefault property that is inherited from EnabledLogicalElement. The EnabledLogicalElement addresses the same semantics. The change to a uint16 data type was discussed when CIM V2.0 was defined. However, existing V1.0 implementations used the string property. To remain compatible with those implementations, StartMode was grandfathered into the schema. Use of the deprecated qualifier allows the maintenance of the existing property but also permits an improved, clarified definition using EnabledDefault.
	///Deprecated description: StartMode is a string value that indicates whether the Service is automatically started by a System, an Operating System, and so on, or is started only upon request.
	StartMode string

	// The CreationClassName of the scoping System.
	SystemCreationClassName string

	// The Name of the scoping System.
	SystemName string
}

func NewCIM_ServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_Service, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Service{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_ServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Service, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Service{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Service) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Service) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetPrimaryOwnerContact sets the value of PrimaryOwnerContact for the instance
func (instance *CIM_Service) SetPropertyPrimaryOwnerContact(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerContact", (value))
}

// GetPrimaryOwnerContact gets the value of PrimaryOwnerContact for the instance
func (instance *CIM_Service) GetPropertyPrimaryOwnerContact() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerContact")
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

// SetPrimaryOwnerName sets the value of PrimaryOwnerName for the instance
func (instance *CIM_Service) SetPropertyPrimaryOwnerName(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerName", (value))
}

// GetPrimaryOwnerName gets the value of PrimaryOwnerName for the instance
func (instance *CIM_Service) GetPropertyPrimaryOwnerName() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerName")
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

// SetStarted sets the value of Started for the instance
func (instance *CIM_Service) SetPropertyStarted(value bool) (err error) {
	return instance.SetProperty("Started", (value))
}

// GetStarted gets the value of Started for the instance
func (instance *CIM_Service) GetPropertyStarted() (value bool, err error) {
	retValue, err := instance.GetProperty("Started")
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

// SetStartMode sets the value of StartMode for the instance
func (instance *CIM_Service) SetPropertyStartMode(value string) (err error) {
	return instance.SetProperty("StartMode", (value))
}

// GetStartMode gets the value of StartMode for the instance
func (instance *CIM_Service) GetPropertyStartMode() (value string, err error) {
	retValue, err := instance.GetProperty("StartMode")
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

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_Service) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", (value))
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_Service) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
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

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_Service) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", (value))
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_Service) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
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

// The StartService method places the Service in the started state. Note that the function of this method overlaps with the RequestedState property. RequestedState was added to the model to maintain a record (such as a persisted value) of the last state request. Invoking the StartService method should set the RequestedState property appropriately. The method returns an integer value of 0 if the Service was successfully started, 1 if the request is not supported, and any other number to indicate an error. In a subclass, the set of possible return codes could be specified using a ValueMap qualifier on the method. The strings to which the ValueMap contents are translated can also be specified in the subclass as a Values array qualifier.
///
///Note: The semantics of this method overlap with the RequestStateChange method that is inherited from EnabledLogicalElement. This method is maintained because it has been widely implemented, and its simple "start" semantics are convenient to use.

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Service) StartService() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartService")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// The StopService method places the Service in the stopped state. Note that the function of this method overlaps with the RequestedState property. RequestedState was added to the model to maintain a record (such as a persisted value) of the last state request. Invoking the StopService method should set the RequestedState property appropriately. The method returns an integer value of 0 if the Service was successfully stopped, 1 if the request is not supported, and any other number to indicate an error. In a subclass, the set of possible return codes could be specified using a ValueMap qualifier on the method. The strings to which the ValueMap contents are translated can also be specified in the subclass as a Values array qualifier.
///
///Note: The semantics of this method overlap with the RequestStateChange method that is inherited from EnabledLogicalElement. This method is maintained because it has been widely implemented, and its simple "stop" semantics are convenient to use.

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Service) StopService() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopService")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
