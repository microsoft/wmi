// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Hgs
//////////////////////////////////////////////
package hgs

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_HgsKeyProtectorOperations struct
type MSFT_HgsKeyProtectorOperations struct {
	cim.WmiInstance
}

//

// <param name="IngressKeyProtector" type="uint8 []"></param>

// <param name="EgressKeyProtector" type="uint8 []"></param>
// <param name="EncryptedKeys" type="uint8 []"></param>
// <param name="EncryptedTransferKey" type="uint8 []"></param>
// <param name="EncryptedWrappingKey" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtectorOperations) UnwrapKeyProtector( /* IN */ IngressKeyProtector []uint8,
	/* OUT */ EncryptedTransferKey []uint8,
	/* OUT */ EncryptedWrappingKey []uint8,
	/* OUT */ EncryptedKeys []uint8,
	/* OUT */ EgressKeyProtector []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnwrapKeyProtector", IngressKeyProtector)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EgressKeyProtector" type="uint8 []"></param>
// <param name="EncryptedKeys" type="uint8 []"></param>
// <param name="EncryptedTransferKey" type="uint8 []"></param>
// <param name="EncryptedWrappingKey" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtectorOperations) CreateKeyProtector( /* OUT */ EncryptedTransferKey []uint8,
	/* OUT */ EncryptedWrappingKey []uint8,
	/* OUT */ EncryptedKeys []uint8,
	/* OUT */ EgressKeyProtector []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateKeyProtector")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BaseKeyProtector" type="uint8 []"></param>
// <param name="PlaintextData" type="uint8 []"></param>
// <param name="RollKeyProtector" type="bool "></param>
// <param name="UniqueEncryptionIdentifier" type="uint32 "></param>

// <param name="EgressKeyProtector" type="uint8 []"></param>
// <param name="EncryptedData" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtectorOperations) EncryptDataWithKeyProtector( /* IN */ BaseKeyProtector []uint8,
	/* IN */ UniqueEncryptionIdentifier uint32,
	/* IN */ PlaintextData []uint8,
	/* IN */ RollKeyProtector bool,
	/* OUT */ EgressKeyProtector []uint8,
	/* OUT */ EncryptedData []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EncryptDataWithKeyProtector", BaseKeyProtector, UniqueEncryptionIdentifier, PlaintextData, RollKeyProtector)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BaseKeyProtector" type="uint8 []"></param>
// <param name="EncryptedData" type="uint8 []"></param>

// <param name="PlaintextData" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsKeyProtectorOperations) DecryptDataWithKeyProtector( /* IN */ BaseKeyProtector []uint8,
	/* IN */ EncryptedData []uint8,
	/* OUT */ PlaintextData []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DecryptDataWithKeyProtector", BaseKeyProtector, EncryptedData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
