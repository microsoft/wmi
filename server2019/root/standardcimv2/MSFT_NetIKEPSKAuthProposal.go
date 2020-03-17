// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKEPSKAuthProposal struct
type MSFT_NetIKEPSKAuthProposal struct {
	MSFT_NetIKEAuthProposal

	//
	PreSharedKey string
}

// SetPreSharedKey sets the value of PreSharedKey for the instance
func (instance *MSFT_NetIKEPSKAuthProposal) SetPropertyPreSharedKey(value string) (err error) {
	return instance.SetProperty("PreSharedKey", value)
}

// GetPreSharedKey gets the value of PreSharedKey for the instance
func (instance *MSFT_NetIKEPSKAuthProposal) GetPropertyPreSharedKey() (value string, err error) {
	retValue, err := instance.GetProperty("PreSharedKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
