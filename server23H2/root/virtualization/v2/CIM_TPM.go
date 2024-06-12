// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_TPM struct
type CIM_TPM struct {
	*CIM_LogicalDevice

	// AvailableRequestedTPMStates indicates the possible values for the RequestedTPMState parameter of the method RequestTPMStateChange, used to initiate a state change. The values listed shall be a subset of the values contained in the RequestedTPMStatesSupported property of the associated instance of CIM_TPMCapabilities where the values selected are a function of the current TPM state of the TPM.
	AvailableRequestedTPMStates []TPM_AvailableRequestedTPMStates

	// The TPM manufacturer's major revision.
	TPMManafucturerMajorRevision uint32

	// The TPM manufacturer Identifier as defined by the TCG.
	TPMManufacturerId uint32

	// The additional information defined by the TPM manufacturer.
	TPMManufacturerInfo string

	// The TPM manufacturer's minor revision.
	TPMManufacturerMinorRevision uint32

	// The TPM specification's major version to which the TPM device claims to be conformant.
	TPMSpecMajorVersion uint32

	// The TPM specification's minor version to which the TPM device claims to be conformant.
	TPMSpecMinorVersion uint32

	// Indicates the TPM's operational mode by indicating whether TPM is Enabled, Active and Owned.
	TPMState TPM_TPMState

	// TransitioningToState indicates the TPM's target state to which the TPM is transitioning.
	///A value of 11 "No Change" shall indicate that no transition is in progress.A value of 12 "Not Applicable" shall indicate the implementation does not support representing ongoing transitions.
	///A value other than 11 or 10 shall identify the state to which the element is in the process of transitioning.
	TransitioningToTPMState TPM_TransitioningToTPMState
}

func NewCIM_TPMEx1(instance *cim.WmiInstance) (newInstance *CIM_TPM, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_TPM{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_TPMEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_TPM, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_TPM{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetAvailableRequestedTPMStates sets the value of AvailableRequestedTPMStates for the instance
func (instance *CIM_TPM) SetPropertyAvailableRequestedTPMStates(value []TPM_AvailableRequestedTPMStates) (err error) {
	return instance.SetProperty("AvailableRequestedTPMStates", (value))
}

// GetAvailableRequestedTPMStates gets the value of AvailableRequestedTPMStates for the instance
func (instance *CIM_TPM) GetPropertyAvailableRequestedTPMStates() (value []TPM_AvailableRequestedTPMStates, err error) {
	retValue, err := instance.GetProperty("AvailableRequestedTPMStates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, TPM_AvailableRequestedTPMStates(valuetmp))
	}

	return
}

// SetTPMManafucturerMajorRevision sets the value of TPMManafucturerMajorRevision for the instance
func (instance *CIM_TPM) SetPropertyTPMManafucturerMajorRevision(value uint32) (err error) {
	return instance.SetProperty("TPMManafucturerMajorRevision", (value))
}

// GetTPMManafucturerMajorRevision gets the value of TPMManafucturerMajorRevision for the instance
func (instance *CIM_TPM) GetPropertyTPMManafucturerMajorRevision() (value uint32, err error) {
	retValue, err := instance.GetProperty("TPMManafucturerMajorRevision")
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

// SetTPMManufacturerId sets the value of TPMManufacturerId for the instance
func (instance *CIM_TPM) SetPropertyTPMManufacturerId(value uint32) (err error) {
	return instance.SetProperty("TPMManufacturerId", (value))
}

// GetTPMManufacturerId gets the value of TPMManufacturerId for the instance
func (instance *CIM_TPM) GetPropertyTPMManufacturerId() (value uint32, err error) {
	retValue, err := instance.GetProperty("TPMManufacturerId")
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

// SetTPMManufacturerInfo sets the value of TPMManufacturerInfo for the instance
func (instance *CIM_TPM) SetPropertyTPMManufacturerInfo(value string) (err error) {
	return instance.SetProperty("TPMManufacturerInfo", (value))
}

// GetTPMManufacturerInfo gets the value of TPMManufacturerInfo for the instance
func (instance *CIM_TPM) GetPropertyTPMManufacturerInfo() (value string, err error) {
	retValue, err := instance.GetProperty("TPMManufacturerInfo")
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

// SetTPMManufacturerMinorRevision sets the value of TPMManufacturerMinorRevision for the instance
func (instance *CIM_TPM) SetPropertyTPMManufacturerMinorRevision(value uint32) (err error) {
	return instance.SetProperty("TPMManufacturerMinorRevision", (value))
}

// GetTPMManufacturerMinorRevision gets the value of TPMManufacturerMinorRevision for the instance
func (instance *CIM_TPM) GetPropertyTPMManufacturerMinorRevision() (value uint32, err error) {
	retValue, err := instance.GetProperty("TPMManufacturerMinorRevision")
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

// SetTPMSpecMajorVersion sets the value of TPMSpecMajorVersion for the instance
func (instance *CIM_TPM) SetPropertyTPMSpecMajorVersion(value uint32) (err error) {
	return instance.SetProperty("TPMSpecMajorVersion", (value))
}

// GetTPMSpecMajorVersion gets the value of TPMSpecMajorVersion for the instance
func (instance *CIM_TPM) GetPropertyTPMSpecMajorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("TPMSpecMajorVersion")
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

// SetTPMSpecMinorVersion sets the value of TPMSpecMinorVersion for the instance
func (instance *CIM_TPM) SetPropertyTPMSpecMinorVersion(value uint32) (err error) {
	return instance.SetProperty("TPMSpecMinorVersion", (value))
}

// GetTPMSpecMinorVersion gets the value of TPMSpecMinorVersion for the instance
func (instance *CIM_TPM) GetPropertyTPMSpecMinorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("TPMSpecMinorVersion")
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

// SetTPMState sets the value of TPMState for the instance
func (instance *CIM_TPM) SetPropertyTPMState(value TPM_TPMState) (err error) {
	return instance.SetProperty("TPMState", (value))
}

// GetTPMState gets the value of TPMState for the instance
func (instance *CIM_TPM) GetPropertyTPMState() (value TPM_TPMState, err error) {
	retValue, err := instance.GetProperty("TPMState")
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

	value = TPM_TPMState(valuetmp)

	return
}

// SetTransitioningToTPMState sets the value of TransitioningToTPMState for the instance
func (instance *CIM_TPM) SetPropertyTransitioningToTPMState(value TPM_TransitioningToTPMState) (err error) {
	return instance.SetProperty("TransitioningToTPMState", (value))
}

// GetTransitioningToTPMState gets the value of TransitioningToTPMState for the instance
func (instance *CIM_TPM) GetPropertyTransitioningToTPMState() (value TPM_TransitioningToTPMState, err error) {
	retValue, err := instance.GetProperty("TransitioningToTPMState")
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

	value = TPM_TransitioningToTPMState(valuetmp)

	return
}

// Requests that the state of the TPM be changed to the value specified in the RequestedTPMState parameter. If the method invokation completes successfuly, the TPMState property shall be equal to the RequestedTPMState parameter. Invoking the RequestTPMStateChange method multiple times could result in earlier requests being overwritten or lost.
///A return code of 0 shall indicate the state change was successfully initiated.
///A return code of 3 shall indicate that the state transition cannot complete within the interval specified by the TimeoutPeriod parameter.
///A return code of 4096 (0x1000) shall indicate the state change was successfully initiated, a ConcreteJob has been created, and its reference returned in the output parameter Job. Any other return code indicates an error condition.

// <param name="AuthorizationToken" type="string ">Authorization token that may be required for the action to take effect. The AuthorizationToken parameter may be required to establish Physical Presence, or to pass the OwnerAuth, the TCG defined owner authorization password. In the case of OwnerAuth, the CIM_SharedCredential with non-null value of the CIM_SharedCredential.Secret may be required. The CIM_SharedCredential.Algorithm property may also be specified based on the property CIM_TPMCapabilities.SupportedPasswordAlgorithms.</param>
// <param name="RequestedTPMState" type="TPM_RequestedTPMState ">The requested TPM states.</param>
// <param name="TimeoutPeriod" type="string ">A timeout period that specifies the maximum amount of time that the client expects the transition to the new state to take. The interval format must be used to specify the TimeoutPeriod. A value of 0 or a null parameter indicates that the client has no time requirements for the transition.</param>

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the state transition initiated by the method invocation.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_TPM) RequestTPMStateChange( /* IN */ RequestedTPMState TPM_RequestedTPMState,
	/* IN */ AuthorizationToken string,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestTPMStateChange", Action, PercentComplete, Timeout, RequestedTPMState, AuthorizationToken, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method changes the owner authorization credential of the TPM device. The old and new owner authorization passwords are required.Reference: See Section 17 (Changing AuthData) of Spec (#3).

// <param name="NewOwnerAuth" type="string ">NewOwnerAuth represents new owner authorization credential required to take ownership of the TPM device.The CIM_SharedCredential subclass may be required with non-null value of the CIM_SharedCredential.Secret property for the parameter.</param>
// <param name="OldOwnerAuth" type="string ">OldOwnerAuth represents old owner authorization credential required to take ownership of the TPM device.The CIM_SharedCredential subclass may be required with non-null value of the CIM_SharedCredential.Secret property for the parameter.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_TPM) ChangeOwnerAuth( /* IN */ OldOwnerAuth string,
	/* IN */ NewOwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ChangeOwnerAuth", OldOwnerAuth, NewOwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
