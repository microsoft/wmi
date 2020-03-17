// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterLsoEncapsulationTypes struct
type MSFT_NetAdapterLsoEncapsulationTypes struct {
	cim.WmiInstance

	//
	NdisEncapsulationIeee802_3 bool

	//
	NdisEncapsulationIeee802_3pAndq bool

	//
	NdisEncapsulationIeee802_3PAndQInOob bool

	//
	NdisEncapsulationIeeLlcSnapRouted bool

	//
	NdisEncapsulationNotNull bool

	//
	NdisEncapsulationNotSupported bool
}

// SetNdisEncapsulationIeee802_3 sets the value of NdisEncapsulationIeee802_3 for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationIeee802_3", value)
}

// GetNdisEncapsulationIeee802_3 gets the value of NdisEncapsulationIeee802_3 for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisEncapsulationIeee802_3pAndq sets the value of NdisEncapsulationIeee802_3pAndq for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3pAndq(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationIeee802_3pAndq", value)
}

// GetNdisEncapsulationIeee802_3pAndq gets the value of NdisEncapsulationIeee802_3pAndq for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3pAndq() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3pAndq")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisEncapsulationIeee802_3PAndQInOob sets the value of NdisEncapsulationIeee802_3PAndQInOob for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationIeee802_3PAndQInOob(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationIeee802_3PAndQInOob", value)
}

// GetNdisEncapsulationIeee802_3PAndQInOob gets the value of NdisEncapsulationIeee802_3PAndQInOob for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationIeee802_3PAndQInOob() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationIeee802_3PAndQInOob")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisEncapsulationIeeLlcSnapRouted sets the value of NdisEncapsulationIeeLlcSnapRouted for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationIeeLlcSnapRouted(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationIeeLlcSnapRouted", value)
}

// GetNdisEncapsulationIeeLlcSnapRouted gets the value of NdisEncapsulationIeeLlcSnapRouted for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationIeeLlcSnapRouted() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationIeeLlcSnapRouted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisEncapsulationNotNull sets the value of NdisEncapsulationNotNull for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationNotNull(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationNotNull", value)
}

// GetNdisEncapsulationNotNull gets the value of NdisEncapsulationNotNull for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationNotNull() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationNotNull")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisEncapsulationNotSupported sets the value of NdisEncapsulationNotSupported for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) SetPropertyNdisEncapsulationNotSupported(value bool) (err error) {
	return instance.SetProperty("NdisEncapsulationNotSupported", value)
}

// GetNdisEncapsulationNotSupported gets the value of NdisEncapsulationNotSupported for the instance
func (instance *MSFT_NetAdapterLsoEncapsulationTypes) GetPropertyNdisEncapsulationNotSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("NdisEncapsulationNotSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
