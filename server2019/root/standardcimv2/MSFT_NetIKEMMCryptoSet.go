// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKEMMCryptoSet struct
type MSFT_NetIKEMMCryptoSet struct {
	MSFT_NetIKECryptoSet

	//
	ForceDiffieHellman bool

	//
	MaxLifetimeMinutes uint32

	//
	MaxLifetimeSessions uint32
}

// SetForceDiffieHellman sets the value of ForceDiffieHellman for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyForceDiffieHellman(value bool) (err error) {
	return instance.SetProperty("ForceDiffieHellman", value)
}

// GetForceDiffieHellman gets the value of ForceDiffieHellman for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyForceDiffieHellman() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceDiffieHellman")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLifetimeMinutes sets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyMaxLifetimeMinutes(value uint32) (err error) {
	return instance.SetProperty("MaxLifetimeMinutes", value)
}

// GetMaxLifetimeMinutes gets the value of MaxLifetimeMinutes for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyMaxLifetimeMinutes() (value uint32, err error) {
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

// SetMaxLifetimeSessions sets the value of MaxLifetimeSessions for the instance
func (instance *MSFT_NetIKEMMCryptoSet) SetPropertyMaxLifetimeSessions(value uint32) (err error) {
	return instance.SetProperty("MaxLifetimeSessions", value)
}

// GetMaxLifetimeSessions gets the value of MaxLifetimeSessions for the instance
func (instance *MSFT_NetIKEMMCryptoSet) GetPropertyMaxLifetimeSessions() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeSessions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEMMCryptoSet) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewID" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEMMCryptoSet) CloneObject( /* IN */ NewName string,
	/* IN */ NewID string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewID, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
