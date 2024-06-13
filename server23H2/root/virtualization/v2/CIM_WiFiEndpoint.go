// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_WiFiEndpoint struct
type CIM_WiFiEndpoint struct {
	*CIM_LANEndpoint

	// AccessPointAddress shall contain the MAC address of the access point with which the WiFiEndpoint is currently associated. If the WiFiEndpoint is not currently associated, then AccessPointAddress shall be NULL.The MAC address shall be formatted as twelve hexadecimal digits (for example, "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order. (Therefore, the Group address bit is found in the low order bit of the first character of the string.)
	AccessPointAddress string

	// Associated shall indicate whether or not the WiFiEndpoint is currently associated to an access point or client station.
	Associated bool

	// AuthenticationMethod shall specify the method used to authenticate the WiFiEndpoint and the network to one another.
	///	* Unknown (0): shall indicate that the authentication method is unknown to the server.	* Other (1): shall indicate that the authentication method is known to the server but not specified in the list below. If AuthenticationMethod contains 1, OtherAuthenticationMethod shall not be NULL and shall not be empty.
	///	* Open System (2): shall indicate that the authentication method is Open System. AuthenticationMethod shall contain 2 only if EncryptionMethod contains 2 ("WEP").
	///	* Shared Key (3): shall indicate that the authentication method is Shared Key. AuthenticationMethod shall contain 3 only if EncryptionMethod contains 2 ("WEP").
	///	* WPA PSK (4): shall indicate that the authentication method is WPA (Wi-Fi Protected Access) PSK (Pre-Shared Key). AuthenticationMethod shall contain 4 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	///	* WPA IEEE 802.1x (5): shall indicate that the authentication method is WPA (Wi-Fi Protected Access) IEEE 802.1x. AuthenticationMethod shall contain 5 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	///	* WPA2 PSK (6): shall indicate that the authentication method is WPA2 (Wi-Fi Protected Access Version 2) PSK (Pre-Shared Key). AuthenticationMethod shall contain 6 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	///	* WPA2 IEEE 802.1x (7): shall indicate that the authentication method is WPA2 (Wi-Fi Protected Access Version 2) IEEE 802.1x. AuthenticationMethod shall contain 6 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	///	* CCKM IEEE 802.1x (8): CCKM (Cisco Centralized Key Management with LEAP or EAP-FAST)
	AuthenticationMethod WiFiEndpoint_AuthenticationMethod

	// BSSType shall indicate the Basic Service Set (BSS) Type of the network that corresponds to the instance. A Basic Service Set is a set of stations controlled by a single coordination function.
	///	* Independent: the WiFiEndpoint is associated directly to another client station.
	///	* Infrastructure: the WiFiEndpoint is associated to a network via an access point.
	BSSType WiFiEndpoint_BSSType

	// EncryptionMethod shall specify the encryption method in use to protect the confidentiality of data sent and received by the WiFiEndpoint.
	///	* Unknown (0): shall indicate that the encryption method is unknown to the server.
	///	* Other (1): shall indicate that the encryption method is known to the server but not defined in the list below. If EncryptionMethod contains 1, OtherEncryptionMethod shall not be NULL and shall not be empty.
	///	* WEP (2): shall indicate that the encryption method is Wired Equivalency Privacy (WEP). The value of EncryptionMethod shall be 2 only if the value of AuthenticationMethod is 2 ("Open System") or 3 ("Shared Key").
	///	* TKIP (3): shall indicate that the encryption method is Temporal Key Integrity Protocol (TKIP). The value of EncryptionMethod shall be 3 only if the value of AuthenticationMethod is 4 ("WPA PSK"), 5 ("WPA IEEE 802.1x"), 6 ("WPA2 PSK"), or 7 ("WPA2 IEEE 802.1x").
	///	* CCMP (4): shall indicate that the encryption method is Counter Mode with Cipher Block Chaining Message Authentication Code Protocol (CCMP). The value of EncryptionMethod shall be 4 only if the value of AuthenticationMethod is 4 ("WPA PSK"), 5 ("WPA IEEE 802.1x"), 6 ("WPA2 PSK"), or 7 ("WPA2 IEEE 802.1x").
	///	* None (5): shall indicate that no encryption method is in use. The value of EncryptionMethod shall be 5 only if the value of AuthenticationMethod is 2 ("Open System") or 3 ("Shared Key").
	EncryptionMethod WiFiEndpoint_EncryptionMethod

	// IEEE8021xAuthenticationProtocol shall contain the EAP (Extensible Authentication Protocol) type if and only if AuthenticationMethod contains "WPA IEEE 802.1x" or "WPA2 IEEE 802.1x" or "CCKM IEEE 802.1x"
	///.	* EAP-TLS (0): shall indicate the Transport Layer Security EAP type specified in RFC 2716.
	///	* EAP-TTLS/MSCHAPv2 (1): shall indicate the Tunneled TLS Authentication Protocol EAP type (specified in draft-ietf-pppext-eap-ttls) with Microsoft PPP CHAP Extensions, Version 2 (MSCHAPv2) as the inner authentication method.
	///	* PEAPv0/EAP-MSCHAPv2 (2): shall indicate the Protected Extensible Authentication Protocol (PEAP) Version 0 EAP type (specified in draft-kamath-pppext-peapv0), with Microsoft PPP CHAP Extensions, Version 2 (MSCHAPv2) as the inner authentication method.
	///	* PEAPv1/EAP-GTC (3): shall indicate the Protected Extensible Authentication Protocol (PEAP) Version 1 EAP type (specified in draft-josefsson-pppext-eap-tls-eap), with Generic Token Card (GTC) as the inner authentication method.
	///	* EAP-FAST/MSCHAPv2 (4): shall indicate the Flexible Authentication Extensible Authentication Protocol EAP type specified in IETF RFC 4851, with Microsoft PPP CHAP Extensions, Version 2 (MSCHAPv2) as the inner authentication method.
	///	* EAP-FAST/GTC (5): shall indicate the Flexible Authentication Extensible Authentication Protocol EAP type specified in IETF RFC 4851, with Generic Token Card (GTC) as the inner authentication method.
	///	* EAP-MD5 (6): shall indicate the EAP MD5 authentication method, specified in RFC 3748.
	///	* EAP-PSK (7): shall indicate the EAP-PSK (Pre-shared Key) Protocol specified in RFC 4764.
	///	* EAP-SIM (8): shall indicate the Extensible Authentication Protocol Method for Global System for Mobile Communications (GSM) Subscriber Identity Modules (EAP-SIM), specified in RFC 4186.
	///	* EAP-AKA (9): shall indicate the Extensible Authentication Protocol Method for 3rd Generation Authentication and Key Agreement (EAP-AKA) authentication method, specified in RFC 4187.
	///	* EAP-FAST/TLS (10): shall indicate the Flexible Authentication Extensible Authentication Protocol EAP type specified in IETF RFC 4851, with TLS as the inner authentication method.
	IEEE8021xAuthenticationProtocol WiFiEndpoint_IEEE8021xAuthenticationProtocol

	// OtherAuthenticationMethod shall specify the 802.11 authentication method if and only if AuthenticationMethod contains "Other". The format of this string shall be vendor-specific.
	OtherAuthenticationMethod string

	// OtherEncryptionMethod shall specify the 802.11 encryption method if and only if EncryptionMethod contains "Other". The format of this string shall be vendor-specific.
	OtherEncryptionMethod string
}

func NewCIM_WiFiEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_WiFiEndpoint, err error) {
	tmp, err := NewCIM_LANEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_WiFiEndpoint{
		CIM_LANEndpoint: tmp,
	}
	return
}

func NewCIM_WiFiEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_WiFiEndpoint, err error) {
	tmp, err := NewCIM_LANEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_WiFiEndpoint{
		CIM_LANEndpoint: tmp,
	}
	return
}

// SetAccessPointAddress sets the value of AccessPointAddress for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyAccessPointAddress(value string) (err error) {
	return instance.SetProperty("AccessPointAddress", (value))
}

// GetAccessPointAddress gets the value of AccessPointAddress for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyAccessPointAddress() (value string, err error) {
	retValue, err := instance.GetProperty("AccessPointAddress")
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

// SetAssociated sets the value of Associated for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyAssociated(value bool) (err error) {
	return instance.SetProperty("Associated", (value))
}

// GetAssociated gets the value of Associated for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyAssociated() (value bool, err error) {
	retValue, err := instance.GetProperty("Associated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyAuthenticationMethod(value WiFiEndpoint_AuthenticationMethod) (err error) {
	return instance.SetProperty("AuthenticationMethod", (value))
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyAuthenticationMethod() (value WiFiEndpoint_AuthenticationMethod, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WiFiEndpoint_AuthenticationMethod(valuetmp)

	return
}

// SetBSSType sets the value of BSSType for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyBSSType(value WiFiEndpoint_BSSType) (err error) {
	return instance.SetProperty("BSSType", (value))
}

// GetBSSType gets the value of BSSType for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyBSSType() (value WiFiEndpoint_BSSType, err error) {
	retValue, err := instance.GetProperty("BSSType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WiFiEndpoint_BSSType(valuetmp)

	return
}

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyEncryptionMethod(value WiFiEndpoint_EncryptionMethod) (err error) {
	return instance.SetProperty("EncryptionMethod", (value))
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyEncryptionMethod() (value WiFiEndpoint_EncryptionMethod, err error) {
	retValue, err := instance.GetProperty("EncryptionMethod")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WiFiEndpoint_EncryptionMethod(valuetmp)

	return
}

// SetIEEE8021xAuthenticationProtocol sets the value of IEEE8021xAuthenticationProtocol for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyIEEE8021xAuthenticationProtocol(value WiFiEndpoint_IEEE8021xAuthenticationProtocol) (err error) {
	return instance.SetProperty("IEEE8021xAuthenticationProtocol", (value))
}

// GetIEEE8021xAuthenticationProtocol gets the value of IEEE8021xAuthenticationProtocol for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyIEEE8021xAuthenticationProtocol() (value WiFiEndpoint_IEEE8021xAuthenticationProtocol, err error) {
	retValue, err := instance.GetProperty("IEEE8021xAuthenticationProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = WiFiEndpoint_IEEE8021xAuthenticationProtocol(valuetmp)

	return
}

// SetOtherAuthenticationMethod sets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyOtherAuthenticationMethod(value string) (err error) {
	return instance.SetProperty("OtherAuthenticationMethod", (value))
}

// GetOtherAuthenticationMethod gets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyOtherAuthenticationMethod() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAuthenticationMethod")
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

// SetOtherEncryptionMethod sets the value of OtherEncryptionMethod for the instance
func (instance *CIM_WiFiEndpoint) SetPropertyOtherEncryptionMethod(value string) (err error) {
	return instance.SetProperty("OtherEncryptionMethod", (value))
}

// GetOtherEncryptionMethod gets the value of OtherEncryptionMethod for the instance
func (instance *CIM_WiFiEndpoint) GetPropertyOtherEncryptionMethod() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEncryptionMethod")
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
