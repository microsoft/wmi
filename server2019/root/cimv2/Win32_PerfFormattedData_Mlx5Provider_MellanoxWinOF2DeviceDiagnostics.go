// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics struct {
	*Win32_PerfFormattedData

	//
	AvailableDynamicMSIXCount uint64

	//
	CQslicesbusytime uint64

	//
	EQslicesbusytime uint64

	//
	FastpacketshitinLDBL1cache uint64

	//
	FastpacketshitinREQSLL1 uint64

	//
	FastpacketsmissinLDBL1cache uint64

	//
	FastpacketsmissinREQSLL1 uint64

	//
	Fastpathpacketssliceload uint64

	//
	InternalRQoutofbuffer uint64

	//
	L0MPThit uint64

	//
	L0MPThitPerSec uint32

	//
	L0MPTmiss uint64

	//
	L0MPTmissPerSec uint32

	//
	L0MTThit uint64

	//
	L0MTThitPerSec uint32

	//
	L0MTTmiss uint64

	//
	L0MTTmissPerSec uint32

	//
	L1MPThit uint64

	//
	L1MPThitPerSec uint32

	//
	L1MPTmiss uint64

	//
	L1MPTmissPerSec uint32

	//
	L1MTThit uint64

	//
	L1MTThitPerSec uint32

	//
	L1MTTmiss uint64

	//
	L1MTTmissPerSec uint32

	//
	MSIXslicesbusytime uint64

	//
	NictemperatureinCelsiusdegreesunit uint64

	//
	NoPXTcreditstime uint64

	//
	Nosendcreditsforschedulingtime uint64

	//
	Noslowpathsendcreditsforschedulingtime uint64

	//
	PacketshitinLDBL2cache uint64

	//
	PacketshitinREQSLL2 uint64

	//
	PacketsmissinLDBL2cache uint64

	//
	PacketsmissinREQSLL2 uint64

	//
	PacketssentbySXWtoSXP uint64

	//
	QPdoneduetodesched uint64

	//
	QPdoneduetoE2Ecredits uint64

	//
	QPdoneduetolimited uint64

	//
	QPdoneduetoVLlimited uint64

	//
	QPdoneduetoworkdone uint64

	//
	ReceiveWQEcachehit uint64

	//
	ReceiveWQEcachemiss uint64

	//
	RXSnofastpathcredits uint64

	//
	RXSnoslowpathcredits uint64

	//
	RXTnofastpathcredits uint64

	//
	RXTnoslowpathcredits uint64

	//
	SlowpacketshitinLDBL1cache uint64

	//
	SlowpacketshitinREQSLL1 uint64

	//
	SlowpacketsmissinLDBL1cache uint64

	//
	SlowpacketsmissinREQSLL1 uint64

	//
	Slowpathpacketssliceload uint64

	//
	Steeringhit uint64

	//
	Steeringmiss uint64

	//
	Steeringpipe0processingtime uint64

	//
	Steeringpipe1processingtime uint64

	//
	Steeringprocessingtime uint64

	//
	TotalDynamicMSIXCount uint64

	//
	TPTindirectmemorykeyaccess uint64

	//
	WQEaddresstranslationbackpressure uint64
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAvailableDynamicMSIXCount sets the value of AvailableDynamicMSIXCount for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyAvailableDynamicMSIXCount(value uint64) (err error) {
	return instance.SetProperty("AvailableDynamicMSIXCount", (value))
}

// GetAvailableDynamicMSIXCount gets the value of AvailableDynamicMSIXCount for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyAvailableDynamicMSIXCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableDynamicMSIXCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetCQslicesbusytime sets the value of CQslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyCQslicesbusytime(value uint64) (err error) {
	return instance.SetProperty("CQslicesbusytime", (value))
}

// GetCQslicesbusytime gets the value of CQslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyCQslicesbusytime() (value uint64, err error) {
	retValue, err := instance.GetProperty("CQslicesbusytime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetEQslicesbusytime sets the value of EQslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyEQslicesbusytime(value uint64) (err error) {
	return instance.SetProperty("EQslicesbusytime", (value))
}

// GetEQslicesbusytime gets the value of EQslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyEQslicesbusytime() (value uint64, err error) {
	retValue, err := instance.GetProperty("EQslicesbusytime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastpacketshitinLDBL1cache sets the value of FastpacketshitinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyFastpacketshitinLDBL1cache(value uint64) (err error) {
	return instance.SetProperty("FastpacketshitinLDBL1cache", (value))
}

// GetFastpacketshitinLDBL1cache gets the value of FastpacketshitinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyFastpacketshitinLDBL1cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastpacketshitinLDBL1cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastpacketshitinREQSLL1 sets the value of FastpacketshitinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyFastpacketshitinREQSLL1(value uint64) (err error) {
	return instance.SetProperty("FastpacketshitinREQSLL1", (value))
}

// GetFastpacketshitinREQSLL1 gets the value of FastpacketshitinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyFastpacketshitinREQSLL1() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastpacketshitinREQSLL1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastpacketsmissinLDBL1cache sets the value of FastpacketsmissinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyFastpacketsmissinLDBL1cache(value uint64) (err error) {
	return instance.SetProperty("FastpacketsmissinLDBL1cache", (value))
}

// GetFastpacketsmissinLDBL1cache gets the value of FastpacketsmissinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyFastpacketsmissinLDBL1cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastpacketsmissinLDBL1cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastpacketsmissinREQSLL1 sets the value of FastpacketsmissinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyFastpacketsmissinREQSLL1(value uint64) (err error) {
	return instance.SetProperty("FastpacketsmissinREQSLL1", (value))
}

// GetFastpacketsmissinREQSLL1 gets the value of FastpacketsmissinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyFastpacketsmissinREQSLL1() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastpacketsmissinREQSLL1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFastpathpacketssliceload sets the value of Fastpathpacketssliceload for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyFastpathpacketssliceload(value uint64) (err error) {
	return instance.SetProperty("Fastpathpacketssliceload", (value))
}

// GetFastpathpacketssliceload gets the value of Fastpathpacketssliceload for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyFastpathpacketssliceload() (value uint64, err error) {
	retValue, err := instance.GetProperty("Fastpathpacketssliceload")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInternalRQoutofbuffer sets the value of InternalRQoutofbuffer for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyInternalRQoutofbuffer(value uint64) (err error) {
	return instance.SetProperty("InternalRQoutofbuffer", (value))
}

// GetInternalRQoutofbuffer gets the value of InternalRQoutofbuffer for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyInternalRQoutofbuffer() (value uint64, err error) {
	retValue, err := instance.GetProperty("InternalRQoutofbuffer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL0MPThit sets the value of L0MPThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MPThit(value uint64) (err error) {
	return instance.SetProperty("L0MPThit", (value))
}

// GetL0MPThit gets the value of L0MPThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MPThit() (value uint64, err error) {
	retValue, err := instance.GetProperty("L0MPThit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL0MPThitPerSec sets the value of L0MPThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MPThitPerSec(value uint32) (err error) {
	return instance.SetProperty("L0MPThitPerSec", (value))
}

// GetL0MPThitPerSec gets the value of L0MPThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MPThitPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L0MPThitPerSec")
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

// SetL0MPTmiss sets the value of L0MPTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MPTmiss(value uint64) (err error) {
	return instance.SetProperty("L0MPTmiss", (value))
}

// GetL0MPTmiss gets the value of L0MPTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MPTmiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("L0MPTmiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL0MPTmissPerSec sets the value of L0MPTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MPTmissPerSec(value uint32) (err error) {
	return instance.SetProperty("L0MPTmissPerSec", (value))
}

// GetL0MPTmissPerSec gets the value of L0MPTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MPTmissPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L0MPTmissPerSec")
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

// SetL0MTThit sets the value of L0MTThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MTThit(value uint64) (err error) {
	return instance.SetProperty("L0MTThit", (value))
}

// GetL0MTThit gets the value of L0MTThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MTThit() (value uint64, err error) {
	retValue, err := instance.GetProperty("L0MTThit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL0MTThitPerSec sets the value of L0MTThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MTThitPerSec(value uint32) (err error) {
	return instance.SetProperty("L0MTThitPerSec", (value))
}

// GetL0MTThitPerSec gets the value of L0MTThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MTThitPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L0MTThitPerSec")
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

// SetL0MTTmiss sets the value of L0MTTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MTTmiss(value uint64) (err error) {
	return instance.SetProperty("L0MTTmiss", (value))
}

// GetL0MTTmiss gets the value of L0MTTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MTTmiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("L0MTTmiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL0MTTmissPerSec sets the value of L0MTTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL0MTTmissPerSec(value uint32) (err error) {
	return instance.SetProperty("L0MTTmissPerSec", (value))
}

// GetL0MTTmissPerSec gets the value of L0MTTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL0MTTmissPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L0MTTmissPerSec")
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

// SetL1MPThit sets the value of L1MPThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MPThit(value uint64) (err error) {
	return instance.SetProperty("L1MPThit", (value))
}

// GetL1MPThit gets the value of L1MPThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MPThit() (value uint64, err error) {
	retValue, err := instance.GetProperty("L1MPThit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL1MPThitPerSec sets the value of L1MPThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MPThitPerSec(value uint32) (err error) {
	return instance.SetProperty("L1MPThitPerSec", (value))
}

// GetL1MPThitPerSec gets the value of L1MPThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MPThitPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L1MPThitPerSec")
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

// SetL1MPTmiss sets the value of L1MPTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MPTmiss(value uint64) (err error) {
	return instance.SetProperty("L1MPTmiss", (value))
}

// GetL1MPTmiss gets the value of L1MPTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MPTmiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("L1MPTmiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL1MPTmissPerSec sets the value of L1MPTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MPTmissPerSec(value uint32) (err error) {
	return instance.SetProperty("L1MPTmissPerSec", (value))
}

// GetL1MPTmissPerSec gets the value of L1MPTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MPTmissPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L1MPTmissPerSec")
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

// SetL1MTThit sets the value of L1MTThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MTThit(value uint64) (err error) {
	return instance.SetProperty("L1MTThit", (value))
}

// GetL1MTThit gets the value of L1MTThit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MTThit() (value uint64, err error) {
	retValue, err := instance.GetProperty("L1MTThit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL1MTThitPerSec sets the value of L1MTThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MTThitPerSec(value uint32) (err error) {
	return instance.SetProperty("L1MTThitPerSec", (value))
}

// GetL1MTThitPerSec gets the value of L1MTThitPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MTThitPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L1MTThitPerSec")
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

// SetL1MTTmiss sets the value of L1MTTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MTTmiss(value uint64) (err error) {
	return instance.SetProperty("L1MTTmiss", (value))
}

// GetL1MTTmiss gets the value of L1MTTmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MTTmiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("L1MTTmiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetL1MTTmissPerSec sets the value of L1MTTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyL1MTTmissPerSec(value uint32) (err error) {
	return instance.SetProperty("L1MTTmissPerSec", (value))
}

// GetL1MTTmissPerSec gets the value of L1MTTmissPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyL1MTTmissPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("L1MTTmissPerSec")
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

// SetMSIXslicesbusytime sets the value of MSIXslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyMSIXslicesbusytime(value uint64) (err error) {
	return instance.SetProperty("MSIXslicesbusytime", (value))
}

// GetMSIXslicesbusytime gets the value of MSIXslicesbusytime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyMSIXslicesbusytime() (value uint64, err error) {
	retValue, err := instance.GetProperty("MSIXslicesbusytime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNictemperatureinCelsiusdegreesunit sets the value of NictemperatureinCelsiusdegreesunit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyNictemperatureinCelsiusdegreesunit(value uint64) (err error) {
	return instance.SetProperty("NictemperatureinCelsiusdegreesunit", (value))
}

// GetNictemperatureinCelsiusdegreesunit gets the value of NictemperatureinCelsiusdegreesunit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyNictemperatureinCelsiusdegreesunit() (value uint64, err error) {
	retValue, err := instance.GetProperty("NictemperatureinCelsiusdegreesunit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNoPXTcreditstime sets the value of NoPXTcreditstime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyNoPXTcreditstime(value uint64) (err error) {
	return instance.SetProperty("NoPXTcreditstime", (value))
}

// GetNoPXTcreditstime gets the value of NoPXTcreditstime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyNoPXTcreditstime() (value uint64, err error) {
	retValue, err := instance.GetProperty("NoPXTcreditstime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNosendcreditsforschedulingtime sets the value of Nosendcreditsforschedulingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyNosendcreditsforschedulingtime(value uint64) (err error) {
	return instance.SetProperty("Nosendcreditsforschedulingtime", (value))
}

// GetNosendcreditsforschedulingtime gets the value of Nosendcreditsforschedulingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyNosendcreditsforschedulingtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Nosendcreditsforschedulingtime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetNoslowpathsendcreditsforschedulingtime sets the value of Noslowpathsendcreditsforschedulingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyNoslowpathsendcreditsforschedulingtime(value uint64) (err error) {
	return instance.SetProperty("Noslowpathsendcreditsforschedulingtime", (value))
}

// GetNoslowpathsendcreditsforschedulingtime gets the value of Noslowpathsendcreditsforschedulingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyNoslowpathsendcreditsforschedulingtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Noslowpathsendcreditsforschedulingtime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketshitinLDBL2cache sets the value of PacketshitinLDBL2cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyPacketshitinLDBL2cache(value uint64) (err error) {
	return instance.SetProperty("PacketshitinLDBL2cache", (value))
}

// GetPacketshitinLDBL2cache gets the value of PacketshitinLDBL2cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyPacketshitinLDBL2cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketshitinLDBL2cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketshitinREQSLL2 sets the value of PacketshitinREQSLL2 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyPacketshitinREQSLL2(value uint64) (err error) {
	return instance.SetProperty("PacketshitinREQSLL2", (value))
}

// GetPacketshitinREQSLL2 gets the value of PacketshitinREQSLL2 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyPacketshitinREQSLL2() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketshitinREQSLL2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsmissinLDBL2cache sets the value of PacketsmissinLDBL2cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyPacketsmissinLDBL2cache(value uint64) (err error) {
	return instance.SetProperty("PacketsmissinLDBL2cache", (value))
}

// GetPacketsmissinLDBL2cache gets the value of PacketsmissinLDBL2cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyPacketsmissinLDBL2cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsmissinLDBL2cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketsmissinREQSLL2 sets the value of PacketsmissinREQSLL2 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyPacketsmissinREQSLL2(value uint64) (err error) {
	return instance.SetProperty("PacketsmissinREQSLL2", (value))
}

// GetPacketsmissinREQSLL2 gets the value of PacketsmissinREQSLL2 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyPacketsmissinREQSLL2() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsmissinREQSLL2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetPacketssentbySXWtoSXP sets the value of PacketssentbySXWtoSXP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyPacketssentbySXWtoSXP(value uint64) (err error) {
	return instance.SetProperty("PacketssentbySXWtoSXP", (value))
}

// GetPacketssentbySXWtoSXP gets the value of PacketssentbySXWtoSXP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyPacketssentbySXWtoSXP() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketssentbySXWtoSXP")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQPdoneduetodesched sets the value of QPdoneduetodesched for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyQPdoneduetodesched(value uint64) (err error) {
	return instance.SetProperty("QPdoneduetodesched", (value))
}

// GetQPdoneduetodesched gets the value of QPdoneduetodesched for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyQPdoneduetodesched() (value uint64, err error) {
	retValue, err := instance.GetProperty("QPdoneduetodesched")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQPdoneduetoE2Ecredits sets the value of QPdoneduetoE2Ecredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyQPdoneduetoE2Ecredits(value uint64) (err error) {
	return instance.SetProperty("QPdoneduetoE2Ecredits", (value))
}

// GetQPdoneduetoE2Ecredits gets the value of QPdoneduetoE2Ecredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyQPdoneduetoE2Ecredits() (value uint64, err error) {
	retValue, err := instance.GetProperty("QPdoneduetoE2Ecredits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQPdoneduetolimited sets the value of QPdoneduetolimited for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyQPdoneduetolimited(value uint64) (err error) {
	return instance.SetProperty("QPdoneduetolimited", (value))
}

// GetQPdoneduetolimited gets the value of QPdoneduetolimited for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyQPdoneduetolimited() (value uint64, err error) {
	retValue, err := instance.GetProperty("QPdoneduetolimited")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQPdoneduetoVLlimited sets the value of QPdoneduetoVLlimited for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyQPdoneduetoVLlimited(value uint64) (err error) {
	return instance.SetProperty("QPdoneduetoVLlimited", (value))
}

// GetQPdoneduetoVLlimited gets the value of QPdoneduetoVLlimited for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyQPdoneduetoVLlimited() (value uint64, err error) {
	retValue, err := instance.GetProperty("QPdoneduetoVLlimited")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetQPdoneduetoworkdone sets the value of QPdoneduetoworkdone for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyQPdoneduetoworkdone(value uint64) (err error) {
	return instance.SetProperty("QPdoneduetoworkdone", (value))
}

// GetQPdoneduetoworkdone gets the value of QPdoneduetoworkdone for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyQPdoneduetoworkdone() (value uint64, err error) {
	retValue, err := instance.GetProperty("QPdoneduetoworkdone")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReceiveWQEcachehit sets the value of ReceiveWQEcachehit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyReceiveWQEcachehit(value uint64) (err error) {
	return instance.SetProperty("ReceiveWQEcachehit", (value))
}

// GetReceiveWQEcachehit gets the value of ReceiveWQEcachehit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyReceiveWQEcachehit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveWQEcachehit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetReceiveWQEcachemiss sets the value of ReceiveWQEcachemiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyReceiveWQEcachemiss(value uint64) (err error) {
	return instance.SetProperty("ReceiveWQEcachemiss", (value))
}

// GetReceiveWQEcachemiss gets the value of ReceiveWQEcachemiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyReceiveWQEcachemiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveWQEcachemiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRXSnofastpathcredits sets the value of RXSnofastpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyRXSnofastpathcredits(value uint64) (err error) {
	return instance.SetProperty("RXSnofastpathcredits", (value))
}

// GetRXSnofastpathcredits gets the value of RXSnofastpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyRXSnofastpathcredits() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXSnofastpathcredits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRXSnoslowpathcredits sets the value of RXSnoslowpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyRXSnoslowpathcredits(value uint64) (err error) {
	return instance.SetProperty("RXSnoslowpathcredits", (value))
}

// GetRXSnoslowpathcredits gets the value of RXSnoslowpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyRXSnoslowpathcredits() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXSnoslowpathcredits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRXTnofastpathcredits sets the value of RXTnofastpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyRXTnofastpathcredits(value uint64) (err error) {
	return instance.SetProperty("RXTnofastpathcredits", (value))
}

// GetRXTnofastpathcredits gets the value of RXTnofastpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyRXTnofastpathcredits() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXTnofastpathcredits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRXTnoslowpathcredits sets the value of RXTnoslowpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyRXTnoslowpathcredits(value uint64) (err error) {
	return instance.SetProperty("RXTnoslowpathcredits", (value))
}

// GetRXTnoslowpathcredits gets the value of RXTnoslowpathcredits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyRXTnoslowpathcredits() (value uint64, err error) {
	retValue, err := instance.GetProperty("RXTnoslowpathcredits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowpacketshitinLDBL1cache sets the value of SlowpacketshitinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySlowpacketshitinLDBL1cache(value uint64) (err error) {
	return instance.SetProperty("SlowpacketshitinLDBL1cache", (value))
}

// GetSlowpacketshitinLDBL1cache gets the value of SlowpacketshitinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySlowpacketshitinLDBL1cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowpacketshitinLDBL1cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowpacketshitinREQSLL1 sets the value of SlowpacketshitinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySlowpacketshitinREQSLL1(value uint64) (err error) {
	return instance.SetProperty("SlowpacketshitinREQSLL1", (value))
}

// GetSlowpacketshitinREQSLL1 gets the value of SlowpacketshitinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySlowpacketshitinREQSLL1() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowpacketshitinREQSLL1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowpacketsmissinLDBL1cache sets the value of SlowpacketsmissinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySlowpacketsmissinLDBL1cache(value uint64) (err error) {
	return instance.SetProperty("SlowpacketsmissinLDBL1cache", (value))
}

// GetSlowpacketsmissinLDBL1cache gets the value of SlowpacketsmissinLDBL1cache for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySlowpacketsmissinLDBL1cache() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowpacketsmissinLDBL1cache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowpacketsmissinREQSLL1 sets the value of SlowpacketsmissinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySlowpacketsmissinREQSLL1(value uint64) (err error) {
	return instance.SetProperty("SlowpacketsmissinREQSLL1", (value))
}

// GetSlowpacketsmissinREQSLL1 gets the value of SlowpacketsmissinREQSLL1 for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySlowpacketsmissinREQSLL1() (value uint64, err error) {
	retValue, err := instance.GetProperty("SlowpacketsmissinREQSLL1")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSlowpathpacketssliceload sets the value of Slowpathpacketssliceload for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySlowpathpacketssliceload(value uint64) (err error) {
	return instance.SetProperty("Slowpathpacketssliceload", (value))
}

// GetSlowpathpacketssliceload gets the value of Slowpathpacketssliceload for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySlowpathpacketssliceload() (value uint64, err error) {
	retValue, err := instance.GetProperty("Slowpathpacketssliceload")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSteeringhit sets the value of Steeringhit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySteeringhit(value uint64) (err error) {
	return instance.SetProperty("Steeringhit", (value))
}

// GetSteeringhit gets the value of Steeringhit for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySteeringhit() (value uint64, err error) {
	retValue, err := instance.GetProperty("Steeringhit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSteeringmiss sets the value of Steeringmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySteeringmiss(value uint64) (err error) {
	return instance.SetProperty("Steeringmiss", (value))
}

// GetSteeringmiss gets the value of Steeringmiss for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySteeringmiss() (value uint64, err error) {
	retValue, err := instance.GetProperty("Steeringmiss")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSteeringpipe0processingtime sets the value of Steeringpipe0processingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySteeringpipe0processingtime(value uint64) (err error) {
	return instance.SetProperty("Steeringpipe0processingtime", (value))
}

// GetSteeringpipe0processingtime gets the value of Steeringpipe0processingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySteeringpipe0processingtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Steeringpipe0processingtime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSteeringpipe1processingtime sets the value of Steeringpipe1processingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySteeringpipe1processingtime(value uint64) (err error) {
	return instance.SetProperty("Steeringpipe1processingtime", (value))
}

// GetSteeringpipe1processingtime gets the value of Steeringpipe1processingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySteeringpipe1processingtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Steeringpipe1processingtime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSteeringprocessingtime sets the value of Steeringprocessingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertySteeringprocessingtime(value uint64) (err error) {
	return instance.SetProperty("Steeringprocessingtime", (value))
}

// GetSteeringprocessingtime gets the value of Steeringprocessingtime for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertySteeringprocessingtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Steeringprocessingtime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTotalDynamicMSIXCount sets the value of TotalDynamicMSIXCount for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyTotalDynamicMSIXCount(value uint64) (err error) {
	return instance.SetProperty("TotalDynamicMSIXCount", (value))
}

// GetTotalDynamicMSIXCount gets the value of TotalDynamicMSIXCount for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyTotalDynamicMSIXCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalDynamicMSIXCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTPTindirectmemorykeyaccess sets the value of TPTindirectmemorykeyaccess for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyTPTindirectmemorykeyaccess(value uint64) (err error) {
	return instance.SetProperty("TPTindirectmemorykeyaccess", (value))
}

// GetTPTindirectmemorykeyaccess gets the value of TPTindirectmemorykeyaccess for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyTPTindirectmemorykeyaccess() (value uint64, err error) {
	retValue, err := instance.GetProperty("TPTindirectmemorykeyaccess")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetWQEaddresstranslationbackpressure sets the value of WQEaddresstranslationbackpressure for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) SetPropertyWQEaddresstranslationbackpressure(value uint64) (err error) {
	return instance.SetProperty("WQEaddresstranslationbackpressure", (value))
}

// GetWQEaddresstranslationbackpressure gets the value of WQEaddresstranslationbackpressure for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DeviceDiagnostics) GetPropertyWQEaddresstranslationbackpressure() (value uint64, err error) {
	retValue, err := instance.GetProperty("WQEaddresstranslationbackpressure")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
