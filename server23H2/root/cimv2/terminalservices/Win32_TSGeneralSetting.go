// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSGeneralSetting struct
type Win32_TSGeneralSetting struct {
	*Win32_TerminalSetting

	//
	CertificateName string

	//
	Certificates []uint8

	//
	Comment string

	//
	MinEncryptionLevel uint32

	//
	PolicySourceMinEncryptionLevel uint32

	//
	PolicySourceSecurityLayer uint32

	//
	PolicySourceUserAuthenticationRequired uint32

	//
	SecurityLayer uint32

	//
	SSLCertificateSHA1Hash string

	//
	SSLCertificateSHA1HashType uint32

	//
	TerminalProtocol string

	//
	Transport string

	//
	UserAuthenticationRequired uint32

	//
	WindowsAuthentication uint32
}

func NewWin32_TSGeneralSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSGeneralSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSGeneralSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSGeneralSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSGeneralSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSGeneralSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetCertificateName sets the value of CertificateName for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyCertificateName(value string) (err error) {
	return instance.SetProperty("CertificateName", (value))
}

// GetCertificateName gets the value of CertificateName for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyCertificateName() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCertificates sets the value of Certificates for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyCertificates(value []uint8) (err error) {
	return instance.SetProperty("Certificates", (value))
}

// GetCertificates gets the value of Certificates for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyCertificates() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Certificates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetComment sets the value of Comment for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyComment(value string) (err error) {
	return instance.SetProperty("Comment", (value))
}

// GetComment gets the value of Comment for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyComment() (value string, err error) {
	retValue, err := instance.GetProperty("Comment")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMinEncryptionLevel sets the value of MinEncryptionLevel for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyMinEncryptionLevel(value uint32) (err error) {
	return instance.SetProperty("MinEncryptionLevel", (value))
}

// GetMinEncryptionLevel gets the value of MinEncryptionLevel for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyMinEncryptionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinEncryptionLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceMinEncryptionLevel sets the value of PolicySourceMinEncryptionLevel for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyPolicySourceMinEncryptionLevel(value uint32) (err error) {
	return instance.SetProperty("PolicySourceMinEncryptionLevel", (value))
}

// GetPolicySourceMinEncryptionLevel gets the value of PolicySourceMinEncryptionLevel for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyPolicySourceMinEncryptionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceMinEncryptionLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceSecurityLayer sets the value of PolicySourceSecurityLayer for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyPolicySourceSecurityLayer(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSecurityLayer", (value))
}

// GetPolicySourceSecurityLayer gets the value of PolicySourceSecurityLayer for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyPolicySourceSecurityLayer() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSecurityLayer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceUserAuthenticationRequired sets the value of PolicySourceUserAuthenticationRequired for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyPolicySourceUserAuthenticationRequired(value uint32) (err error) {
	return instance.SetProperty("PolicySourceUserAuthenticationRequired", (value))
}

// GetPolicySourceUserAuthenticationRequired gets the value of PolicySourceUserAuthenticationRequired for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyPolicySourceUserAuthenticationRequired() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceUserAuthenticationRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSecurityLayer sets the value of SecurityLayer for the instance
func (instance *Win32_TSGeneralSetting) SetPropertySecurityLayer(value uint32) (err error) {
	return instance.SetProperty("SecurityLayer", (value))
}

// GetSecurityLayer gets the value of SecurityLayer for the instance
func (instance *Win32_TSGeneralSetting) GetPropertySecurityLayer() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecurityLayer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSSLCertificateSHA1Hash sets the value of SSLCertificateSHA1Hash for the instance
func (instance *Win32_TSGeneralSetting) SetPropertySSLCertificateSHA1Hash(value string) (err error) {
	return instance.SetProperty("SSLCertificateSHA1Hash", (value))
}

// GetSSLCertificateSHA1Hash gets the value of SSLCertificateSHA1Hash for the instance
func (instance *Win32_TSGeneralSetting) GetPropertySSLCertificateSHA1Hash() (value string, err error) {
	retValue, err := instance.GetProperty("SSLCertificateSHA1Hash")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSSLCertificateSHA1HashType sets the value of SSLCertificateSHA1HashType for the instance
func (instance *Win32_TSGeneralSetting) SetPropertySSLCertificateSHA1HashType(value uint32) (err error) {
	return instance.SetProperty("SSLCertificateSHA1HashType", (value))
}

// GetSSLCertificateSHA1HashType gets the value of SSLCertificateSHA1HashType for the instance
func (instance *Win32_TSGeneralSetting) GetPropertySSLCertificateSHA1HashType() (value uint32, err error) {
	retValue, err := instance.GetProperty("SSLCertificateSHA1HashType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTerminalProtocol sets the value of TerminalProtocol for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyTerminalProtocol(value string) (err error) {
	return instance.SetProperty("TerminalProtocol", (value))
}

// GetTerminalProtocol gets the value of TerminalProtocol for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyTerminalProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("TerminalProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTransport sets the value of Transport for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyTransport(value string) (err error) {
	return instance.SetProperty("Transport", (value))
}

// GetTransport gets the value of Transport for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyTransport() (value string, err error) {
	retValue, err := instance.GetProperty("Transport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUserAuthenticationRequired sets the value of UserAuthenticationRequired for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyUserAuthenticationRequired(value uint32) (err error) {
	return instance.SetProperty("UserAuthenticationRequired", (value))
}

// GetUserAuthenticationRequired gets the value of UserAuthenticationRequired for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyUserAuthenticationRequired() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserAuthenticationRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetWindowsAuthentication sets the value of WindowsAuthentication for the instance
func (instance *Win32_TSGeneralSetting) SetPropertyWindowsAuthentication(value uint32) (err error) {
	return instance.SetProperty("WindowsAuthentication", (value))
}

// GetWindowsAuthentication gets the value of WindowsAuthentication for the instance
func (instance *Win32_TSGeneralSetting) GetPropertyWindowsAuthentication() (value uint32, err error) {
	retValue, err := instance.GetProperty("WindowsAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="MinEncryptionLevel" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSGeneralSetting) SetEncryptionLevel( /* IN */ MinEncryptionLevel uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetEncryptionLevel", MinEncryptionLevel)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SecurityLayer" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSGeneralSetting) SetSecurityLayer( /* IN */ SecurityLayer uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSecurityLayer", SecurityLayer)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="UserAuthenticationRequired" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSGeneralSetting) SetUserAuthenticationRequired( /* IN */ UserAuthenticationRequired uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetUserAuthenticationRequired", UserAuthenticationRequired)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
