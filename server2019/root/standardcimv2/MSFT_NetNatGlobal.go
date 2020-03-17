// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetNatGlobal struct
type MSFT_NetNatGlobal struct {
	MSFT_NetSettingData

	//
	InterRoutingDomainHairpinningMode uint32
}

// SetInterRoutingDomainHairpinningMode sets the value of InterRoutingDomainHairpinningMode for the instance
func (instance *MSFT_NetNatGlobal) SetPropertyInterRoutingDomainHairpinningMode(value uint32) (err error) {
	return instance.SetProperty("InterRoutingDomainHairpinningMode", value)
}

// GetInterRoutingDomainHairpinningMode gets the value of InterRoutingDomainHairpinningMode for the instance
func (instance *MSFT_NetNatGlobal) GetPropertyInterRoutingDomainHairpinningMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterRoutingDomainHairpinningMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
