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

// SecurityBindingElement struct
type SecurityBindingElement struct {
	*BindingElement

	// A boolean value that specifies if security should be enabled with insecure transports.
	AllowInsecureTransport bool

	// Specifies the algorithms to use with the binding.
	DefaultAlgorithmSuite string

	// A boolean value that specifies if the response can be unsecured.
	EnableUnsecuredResponse bool

	// A boolean value that specifies if each message contains a timestamp.
	IncludeTimestamp bool

	// The source of entropy used to create keys.
	KeyEntropyMode string

	// The binding specific security properties for the local service.
	LocalServiceSecuritySettings LocalServiceSecuritySettings

	// The version used for message security.
	MessageSecurityVersion string

	// The order of elements in the security header for this binding.
	SecurityHeaderLayout string
}

func NewSecurityBindingElementEx1(instance *cim.WmiInstance) (newInstance *SecurityBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SecurityBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewSecurityBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SecurityBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SecurityBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetAllowInsecureTransport sets the value of AllowInsecureTransport for the instance
func (instance *SecurityBindingElement) SetPropertyAllowInsecureTransport(value bool) (err error) {
	return instance.SetProperty("AllowInsecureTransport", (value))
}

// GetAllowInsecureTransport gets the value of AllowInsecureTransport for the instance
func (instance *SecurityBindingElement) GetPropertyAllowInsecureTransport() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowInsecureTransport")
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

// SetDefaultAlgorithmSuite sets the value of DefaultAlgorithmSuite for the instance
func (instance *SecurityBindingElement) SetPropertyDefaultAlgorithmSuite(value string) (err error) {
	return instance.SetProperty("DefaultAlgorithmSuite", (value))
}

// GetDefaultAlgorithmSuite gets the value of DefaultAlgorithmSuite for the instance
func (instance *SecurityBindingElement) GetPropertyDefaultAlgorithmSuite() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultAlgorithmSuite")
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

// SetEnableUnsecuredResponse sets the value of EnableUnsecuredResponse for the instance
func (instance *SecurityBindingElement) SetPropertyEnableUnsecuredResponse(value bool) (err error) {
	return instance.SetProperty("EnableUnsecuredResponse", (value))
}

// GetEnableUnsecuredResponse gets the value of EnableUnsecuredResponse for the instance
func (instance *SecurityBindingElement) GetPropertyEnableUnsecuredResponse() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableUnsecuredResponse")
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

// SetIncludeTimestamp sets the value of IncludeTimestamp for the instance
func (instance *SecurityBindingElement) SetPropertyIncludeTimestamp(value bool) (err error) {
	return instance.SetProperty("IncludeTimestamp", (value))
}

// GetIncludeTimestamp gets the value of IncludeTimestamp for the instance
func (instance *SecurityBindingElement) GetPropertyIncludeTimestamp() (value bool, err error) {
	retValue, err := instance.GetProperty("IncludeTimestamp")
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

// SetKeyEntropyMode sets the value of KeyEntropyMode for the instance
func (instance *SecurityBindingElement) SetPropertyKeyEntropyMode(value string) (err error) {
	return instance.SetProperty("KeyEntropyMode", (value))
}

// GetKeyEntropyMode gets the value of KeyEntropyMode for the instance
func (instance *SecurityBindingElement) GetPropertyKeyEntropyMode() (value string, err error) {
	retValue, err := instance.GetProperty("KeyEntropyMode")
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

// SetLocalServiceSecuritySettings sets the value of LocalServiceSecuritySettings for the instance
func (instance *SecurityBindingElement) SetPropertyLocalServiceSecuritySettings(value LocalServiceSecuritySettings) (err error) {
	return instance.SetProperty("LocalServiceSecuritySettings", (value))
}

// GetLocalServiceSecuritySettings gets the value of LocalServiceSecuritySettings for the instance
func (instance *SecurityBindingElement) GetPropertyLocalServiceSecuritySettings() (value LocalServiceSecuritySettings, err error) {
	retValue, err := instance.GetProperty("LocalServiceSecuritySettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(LocalServiceSecuritySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " LocalServiceSecuritySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = LocalServiceSecuritySettings(valuetmp)

	return
}

// SetMessageSecurityVersion sets the value of MessageSecurityVersion for the instance
func (instance *SecurityBindingElement) SetPropertyMessageSecurityVersion(value string) (err error) {
	return instance.SetProperty("MessageSecurityVersion", (value))
}

// GetMessageSecurityVersion gets the value of MessageSecurityVersion for the instance
func (instance *SecurityBindingElement) GetPropertyMessageSecurityVersion() (value string, err error) {
	retValue, err := instance.GetProperty("MessageSecurityVersion")
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

// SetSecurityHeaderLayout sets the value of SecurityHeaderLayout for the instance
func (instance *SecurityBindingElement) SetPropertySecurityHeaderLayout(value string) (err error) {
	return instance.SetProperty("SecurityHeaderLayout", (value))
}

// GetSecurityHeaderLayout gets the value of SecurityHeaderLayout for the instance
func (instance *SecurityBindingElement) GetPropertySecurityHeaderLayout() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityHeaderLayout")
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
