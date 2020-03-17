// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Attestation
//////////////////////////////////////////////
package attestation

// __Trustee struct
type __Trustee struct {
	__SecurityRelatedClass

	//
	Domain string

	//
	Name string

	//
	SID []uint8

	//
	SidLength uint32

	//
	SIDString string

	//
	TIME_CREATED uint64
}

// SetDomain sets the value of Domain for the instance
func (instance *__Trustee) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *__Trustee) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *__Trustee) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *__Trustee) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSID sets the value of SID for the instance
func (instance *__Trustee) SetPropertySID(value []uint8) (err error) {
	return instance.SetProperty("SID", value)
}

// GetSID gets the value of SID for the instance
func (instance *__Trustee) GetPropertySID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SID")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSidLength sets the value of SidLength for the instance
func (instance *__Trustee) SetPropertySidLength(value uint32) (err error) {
	return instance.SetProperty("SidLength", value)
}

// GetSidLength gets the value of SidLength for the instance
func (instance *__Trustee) GetPropertySidLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("SidLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSIDString sets the value of SIDString for the instance
func (instance *__Trustee) SetPropertySIDString(value string) (err error) {
	return instance.SetProperty("SIDString", value)
}

// GetSIDString gets the value of SIDString for the instance
func (instance *__Trustee) GetPropertySIDString() (value string, err error) {
	retValue, err := instance.GetProperty("SIDString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTIME_CREATED sets the value of TIME_CREATED for the instance
func (instance *__Trustee) SetPropertyTIME_CREATED(value uint64) (err error) {
	return instance.SetProperty("TIME_CREATED", value)
}

// GetTIME_CREATED gets the value of TIME_CREATED for the instance
func (instance *__Trustee) GetPropertyTIME_CREATED() (value uint64, err error) {
	retValue, err := instance.GetProperty("TIME_CREATED")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
