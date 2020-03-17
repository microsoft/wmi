// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_PassportForWork_PINComplexity03 struct
type MDM_PassportForWork_PINComplexity03 struct {
	cim.WmiInstance

	//
	Digits int32

	//
	Expiration int32

	//
	History int32

	//
	InstanceID string

	//
	LowercaseLetters int32

	//
	MaximumPINLength int32

	//
	MinimumPINLength int32

	//
	ParentID string

	//
	SpecialCharacters int32

	//
	UppercaseLetters int32
}

// SetDigits sets the value of Digits for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyDigits(value int32) (err error) {
	return instance.SetProperty("Digits", value)
}

// GetDigits gets the value of Digits for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyDigits() (value int32, err error) {
	retValue, err := instance.GetProperty("Digits")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpiration sets the value of Expiration for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyExpiration(value int32) (err error) {
	return instance.SetProperty("Expiration", value)
}

// GetExpiration gets the value of Expiration for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyExpiration() (value int32, err error) {
	retValue, err := instance.GetProperty("Expiration")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHistory sets the value of History for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyHistory(value int32) (err error) {
	return instance.SetProperty("History", value)
}

// GetHistory gets the value of History for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("History")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowercaseLetters sets the value of LowercaseLetters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyLowercaseLetters(value int32) (err error) {
	return instance.SetProperty("LowercaseLetters", value)
}

// GetLowercaseLetters gets the value of LowercaseLetters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyLowercaseLetters() (value int32, err error) {
	retValue, err := instance.GetProperty("LowercaseLetters")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumPINLength sets the value of MaximumPINLength for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyMaximumPINLength(value int32) (err error) {
	return instance.SetProperty("MaximumPINLength", value)
}

// GetMaximumPINLength gets the value of MaximumPINLength for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyMaximumPINLength() (value int32, err error) {
	retValue, err := instance.GetProperty("MaximumPINLength")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumPINLength sets the value of MinimumPINLength for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyMinimumPINLength(value int32) (err error) {
	return instance.SetProperty("MinimumPINLength", value)
}

// GetMinimumPINLength gets the value of MinimumPINLength for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyMinimumPINLength() (value int32, err error) {
	retValue, err := instance.GetProperty("MinimumPINLength")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecialCharacters sets the value of SpecialCharacters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertySpecialCharacters(value int32) (err error) {
	return instance.SetProperty("SpecialCharacters", value)
}

// GetSpecialCharacters gets the value of SpecialCharacters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertySpecialCharacters() (value int32, err error) {
	retValue, err := instance.GetProperty("SpecialCharacters")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUppercaseLetters sets the value of UppercaseLetters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) SetPropertyUppercaseLetters(value int32) (err error) {
	return instance.SetProperty("UppercaseLetters", value)
}

// GetUppercaseLetters gets the value of UppercaseLetters for the instance
func (instance *MDM_PassportForWork_PINComplexity03) GetPropertyUppercaseLetters() (value int32, err error) {
	retValue, err := instance.GetProperty("UppercaseLetters")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
