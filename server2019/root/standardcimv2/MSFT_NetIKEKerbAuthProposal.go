// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKEKerbAuthProposal struct
type MSFT_NetIKEKerbAuthProposal struct {
	MSFT_NetIKEAuthProposal

	//
	KerbProxy string
}

// SetKerbProxy sets the value of KerbProxy for the instance
func (instance *MSFT_NetIKEKerbAuthProposal) SetPropertyKerbProxy(value string) (err error) {
	return instance.SetProperty("KerbProxy", value)
}

// GetKerbProxy gets the value of KerbProxy for the instance
func (instance *MSFT_NetIKEKerbAuthProposal) GetPropertyKerbProxy() (value string, err error) {
	retValue, err := instance.GetProperty("KerbProxy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
