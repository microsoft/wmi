// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceDebugBehavior struct
type ServiceDebugBehavior struct {
	*Behavior

	// Controls the binding for metadata retrieval using HTTP.
	HttpHelpPageBinding Binding

	// Controls whether the service publishes its WSDL at the address controlled by the HttpGetUrl attribute.
	HttpHelpPageEnabled bool

	// Sets the location at which the service WSDL is published for retrieval using HTTP.
	HttpHelpPageUrl string

	// Controls the binding for metadata retrieval using HTTPS.
	HttpsHelpPageBinding Binding

	// Controls whether the service publishes its WSDL over HTTPS at the address controlled by the HttpsGetUrl attribute.
	HttpsHelpPageEnabled bool

	// Sets the location at which the service WSDL is published for retrieval using HTTPS.
	HttpsHelpPageUrl string

	// Specifies whether to include managed exception information in the detail of SOAP faults returned to the clients for debugging purposes.
	IncludeExceptionDetailInFaults bool
}

func NewServiceDebugBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceDebugBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceDebugBehavior{
		Behavior: tmp,
	}
	return
}

func NewServiceDebugBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceDebugBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceDebugBehavior{
		Behavior: tmp,
	}
	return
}

// SetHttpHelpPageBinding sets the value of HttpHelpPageBinding for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpHelpPageBinding(value Binding) (err error) {
	return instance.SetProperty("HttpHelpPageBinding", (value))
}

// GetHttpHelpPageBinding gets the value of HttpHelpPageBinding for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpHelpPageBinding() (value Binding, err error) {
	retValue, err := instance.GetProperty("HttpHelpPageBinding")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Binding)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Binding is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Binding(valuetmp)

	return
}

// SetHttpHelpPageEnabled sets the value of HttpHelpPageEnabled for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpHelpPageEnabled(value bool) (err error) {
	return instance.SetProperty("HttpHelpPageEnabled", (value))
}

// GetHttpHelpPageEnabled gets the value of HttpHelpPageEnabled for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpHelpPageEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HttpHelpPageEnabled")
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

// SetHttpHelpPageUrl sets the value of HttpHelpPageUrl for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpHelpPageUrl(value string) (err error) {
	return instance.SetProperty("HttpHelpPageUrl", (value))
}

// GetHttpHelpPageUrl gets the value of HttpHelpPageUrl for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpHelpPageUrl() (value string, err error) {
	retValue, err := instance.GetProperty("HttpHelpPageUrl")
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

// SetHttpsHelpPageBinding sets the value of HttpsHelpPageBinding for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpsHelpPageBinding(value Binding) (err error) {
	return instance.SetProperty("HttpsHelpPageBinding", (value))
}

// GetHttpsHelpPageBinding gets the value of HttpsHelpPageBinding for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpsHelpPageBinding() (value Binding, err error) {
	retValue, err := instance.GetProperty("HttpsHelpPageBinding")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Binding)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Binding is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Binding(valuetmp)

	return
}

// SetHttpsHelpPageEnabled sets the value of HttpsHelpPageEnabled for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpsHelpPageEnabled(value bool) (err error) {
	return instance.SetProperty("HttpsHelpPageEnabled", (value))
}

// GetHttpsHelpPageEnabled gets the value of HttpsHelpPageEnabled for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpsHelpPageEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HttpsHelpPageEnabled")
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

// SetHttpsHelpPageUrl sets the value of HttpsHelpPageUrl for the instance
func (instance *ServiceDebugBehavior) SetPropertyHttpsHelpPageUrl(value string) (err error) {
	return instance.SetProperty("HttpsHelpPageUrl", (value))
}

// GetHttpsHelpPageUrl gets the value of HttpsHelpPageUrl for the instance
func (instance *ServiceDebugBehavior) GetPropertyHttpsHelpPageUrl() (value string, err error) {
	retValue, err := instance.GetProperty("HttpsHelpPageUrl")
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

// SetIncludeExceptionDetailInFaults sets the value of IncludeExceptionDetailInFaults for the instance
func (instance *ServiceDebugBehavior) SetPropertyIncludeExceptionDetailInFaults(value bool) (err error) {
	return instance.SetProperty("IncludeExceptionDetailInFaults", (value))
}

// GetIncludeExceptionDetailInFaults gets the value of IncludeExceptionDetailInFaults for the instance
func (instance *ServiceDebugBehavior) GetPropertyIncludeExceptionDetailInFaults() (value bool, err error) {
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
