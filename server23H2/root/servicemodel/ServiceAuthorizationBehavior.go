// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceAuthorizationBehavior struct
type ServiceAuthorizationBehavior struct {
	*Behavior

	// A value that controls whether the service attempts to impersonate using the credentials provided by the incoming message.
	ImpersonateCallerForAllOperations bool

	// A value that controls whether the service attempts to impersonate using the credentials provided by the incoming message while serializing the body of the response message.
	ImpersonateOnSerializingReply bool

	// The principal used to carry out operations on the server.
	PrincipalPermissionMode string

	// The name of the ASP .Net role provider.
	RoleProvider string

	// The authorization manager used for custom authorization.
	ServiceAuthorizationManager string
}

func NewServiceAuthorizationBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceAuthorizationBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceAuthorizationBehavior{
		Behavior: tmp,
	}
	return
}

func NewServiceAuthorizationBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceAuthorizationBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceAuthorizationBehavior{
		Behavior: tmp,
	}
	return
}

// SetImpersonateCallerForAllOperations sets the value of ImpersonateCallerForAllOperations for the instance
func (instance *ServiceAuthorizationBehavior) SetPropertyImpersonateCallerForAllOperations(value bool) (err error) {
	return instance.SetProperty("ImpersonateCallerForAllOperations", (value))
}

// GetImpersonateCallerForAllOperations gets the value of ImpersonateCallerForAllOperations for the instance
func (instance *ServiceAuthorizationBehavior) GetPropertyImpersonateCallerForAllOperations() (value bool, err error) {
	retValue, err := instance.GetProperty("ImpersonateCallerForAllOperations")
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

// SetImpersonateOnSerializingReply sets the value of ImpersonateOnSerializingReply for the instance
func (instance *ServiceAuthorizationBehavior) SetPropertyImpersonateOnSerializingReply(value bool) (err error) {
	return instance.SetProperty("ImpersonateOnSerializingReply", (value))
}

// GetImpersonateOnSerializingReply gets the value of ImpersonateOnSerializingReply for the instance
func (instance *ServiceAuthorizationBehavior) GetPropertyImpersonateOnSerializingReply() (value bool, err error) {
	retValue, err := instance.GetProperty("ImpersonateOnSerializingReply")
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

// SetPrincipalPermissionMode sets the value of PrincipalPermissionMode for the instance
func (instance *ServiceAuthorizationBehavior) SetPropertyPrincipalPermissionMode(value string) (err error) {
	return instance.SetProperty("PrincipalPermissionMode", (value))
}

// GetPrincipalPermissionMode gets the value of PrincipalPermissionMode for the instance
func (instance *ServiceAuthorizationBehavior) GetPropertyPrincipalPermissionMode() (value string, err error) {
	retValue, err := instance.GetProperty("PrincipalPermissionMode")
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

// SetRoleProvider sets the value of RoleProvider for the instance
func (instance *ServiceAuthorizationBehavior) SetPropertyRoleProvider(value string) (err error) {
	return instance.SetProperty("RoleProvider", (value))
}

// GetRoleProvider gets the value of RoleProvider for the instance
func (instance *ServiceAuthorizationBehavior) GetPropertyRoleProvider() (value string, err error) {
	retValue, err := instance.GetProperty("RoleProvider")
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

// SetServiceAuthorizationManager sets the value of ServiceAuthorizationManager for the instance
func (instance *ServiceAuthorizationBehavior) SetPropertyServiceAuthorizationManager(value string) (err error) {
	return instance.SetProperty("ServiceAuthorizationManager", (value))
}

// GetServiceAuthorizationManager gets the value of ServiceAuthorizationManager for the instance
func (instance *ServiceAuthorizationBehavior) GetPropertyServiceAuthorizationManager() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceAuthorizationManager")
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
