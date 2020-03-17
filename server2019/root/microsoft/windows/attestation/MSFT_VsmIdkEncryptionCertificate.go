// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Attestation
//////////////////////////////////////////////
package attestation

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_VsmIdkEncryptionCertificate struct
type MSFT_VsmIdkEncryptionCertificate struct {
	cim.WmiInstance
}

//

// <param name="AttestationServerUrl" type="string "></param>
// <param name="Force" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_AttestationResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VsmIdkEncryptionCertificate) Get( /* IN */ AttestationServerUrl string,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput MSFT_AttestationResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", AttestationServerUrl, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
