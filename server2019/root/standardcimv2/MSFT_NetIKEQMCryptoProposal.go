// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKEQMCryptoProposal struct
type MSFT_NetIKEQMCryptoProposal struct {
	MSFT_NetIKECryptoProposal

	//
	Encapsulation uint16

	//
	HashAlgorithmAH uint16

	//
	HashAlgorithmESP uint16

	//
	MaxLifetimeMinutes uint32
}

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) SetPropertyEncapsulation(value uint16) (err error) {
	return instance.SetProperty("Encapsulation", value)
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) GetPropertyEncapsulation() (value uint16, err error) {
	retValue, err := instance.GetProperty("Encapsulation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHashAlgorithmAH sets the value of HashAlgorithmAH for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) SetPropertyHashAlgorithmAH(value uint16) (err error) {
	return instance.SetProperty("HashAlgorithmAH", value)
}

// GetHashAlgorithmAH gets the value of HashAlgorithmAH for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) GetPropertyHashAlgorithmAH() (value uint16, err error) {
	retValue, err := instance.GetProperty("HashAlgorithmAH")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHashAlgorithmESP sets the value of HashAlgorithmESP for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) SetPropertyHashAlgorithmESP(value uint16) (err error) {
	return instance.SetProperty("HashAlgorithmESP", value)
}

// GetHashAlgorithmESP gets the value of HashAlgorithmESP for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) GetPropertyHashAlgorithmESP() (value uint16, err error) {
	retValue, err := instance.GetProperty("HashAlgorithmESP")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLifetimeMinutes sets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) SetPropertyMaxLifetimeMinutes(value uint32) (err error) {
	return instance.SetProperty("MaxLifetimeMinutes", value)
}

// GetMaxLifetimeMinutes gets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEQMCryptoProposal) GetPropertyMaxLifetimeMinutes() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeMinutes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
