// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// SPACES_StorageReliabilityCounter struct
type SPACES_StorageReliabilityCounter struct {
	MSFT_StorageReliabilityCounter
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SPACES_StorageReliabilityCounter) Reset() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reset")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
