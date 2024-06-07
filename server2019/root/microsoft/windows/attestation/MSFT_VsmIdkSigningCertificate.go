// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Attestation
//////////////////////////////////////////////
package attestation
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// MSFT_VsmIdkSigningCertificate struct
type MSFT_VsmIdkSigningCertificate struct { 
	*cim.WmiInstance
}

	func NewMSFT_VsmIdkSigningCertificateEx1(instance *cim.WmiInstance) (newInstance *MSFT_VsmIdkSigningCertificate, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_VsmIdkSigningCertificate {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_VsmIdkSigningCertificateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_VsmIdkSigningCertificate, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_VsmIdkSigningCertificate {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="AttestationServerUrl" type="string "></param>
// <param name="Force" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_AttestationResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_VsmIdkSigningCertificate) Get( /* IN */ AttestationServerUrl string,
 /* IN */ Force bool,
 /* OUT */ cmdletOutput MSFT_AttestationResult) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , AttestationServerUrl, Force)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


