// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_VirtualSystemMigrationService struct
type CIM_VirtualSystemMigrationService struct {
	CIM_Service
}

// Method to move, migrate or relocate a virtual system to a target host specified by a network name or IP address.
///Return code description:
///0: Success: Virtual system was migrated.
///1: Error: Method not supported by implementation.
///2: Error: Virtual system migration failed for unspecified reasons.
///3: Error: Virtual system migration time out; the virtual system state is unknown.
///4: Error: One or more parameters are formally invalid. For example, the value of the DestinationHost parameter could have been specified in an unsupported format.
///5: Error: The source virtual system, the source host system or the target host system are in a state that does allow initiation of the requested virtual system migration; this may be a temporary condition.
///6: Error: One or more input parameters are incompatible as a set, or with respect to the target host. For example the value of the MigrationNewSettingData parameter contains properties that are not supported by the target host system identified by the value of the DestinationHost parameter. Note: The MigrateVirtualSystemToHost( ) methods is intended as a transitional solution only until modelling of cluster support is available.

// <param name="ComputerSystem" type="CIM_ComputerSystem ">Source virtual computer system to be migrated.</param>
// <param name="DestinationHost" type="string ">Target host system for the migration.
///Acceptable formats for this parameter are conveyed through values of elements of the DestinationHostFormatsSupported[ ] array property in the instance of the CIM_VirtualSystemMigrationCapabilities that is associated through the CIM_ElementCapabilities assocation.</param>
// <param name="MigrationSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemMigrationSettingData class representing migration settings applicable to the migration operation.</param>
// <param name="NewResourceSettingData" type="string []">Array of strings each containing an embedded instance of the CIM_ResourceAllocationSettingData class representing new properties applicable to virtual resources in the scope of the virtual system after it is migrated.</param>
// <param name="NewSystemSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemSettingData class representing new properties applicable to the virtual system after it is migrated.</param>

// <param name="Job" type="CIM_ConcreteJob ">If operation is long running then optionally a job may be returned.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemMigrationService) MigrateVirtualSystemToHost( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ DestinationHost string,
	/* IN */ MigrationSettingData string,
	/* IN */ NewSystemSettingData string,
	/* IN */ NewResourceSettingData []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("MigrateVirtualSystemToHost", Action, PercentComplete, Timeout, ComputerSystem, DestinationHost, MigrationSettingData, NewSystemSettingData, NewResourceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Method to move, migrate or relocate a virtual system to a target system.
///Return code description:
///0: Success: Virtual system was migrated
///1: Error: Method not supported by implementation
///2: Error: Virtual system migration failed for unspecified reasons
///3: Error: Virtual system migration time out; the virtual system state is unknown
///4: Error: One or more parameters are formally invalid For example, the value of the Destination System parameter does not contain a valid object path
///5: Error: The source virtual system, the source host system or the target host system are in a state that does allow initiation of the requested virtual system migration; this may be a temporary condition.
///6: Error: One or more input parameters are incompatible as a set, or with respect to the target host. For example the value of the MigrationNewSettingData parameter contains properties that are not supported by the target host system identified by the value of the DestinationSystem parameter.

// <param name="ComputerSystem" type="CIM_ComputerSystem ">Source virtual computer system to be migrated.</param>
// <param name="DestinationSystem" type="CIM_System ">Destination host system whereto migrate the virtual system.</param>
// <param name="MigrationSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemMigrationSettingData class representing migration settings applicable to the migration operation.</param>
// <param name="NewResourceSettingData" type="string []">Array of strings each containing an embedded instance of the CIM_ResourceAllocationSettingData class representing new properties applicable to virtual resources in the scope of the virtual system after it is migrated.</param>
// <param name="NewSystemSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemSettingData class representing new properties applicable to the virtual system after it is migrated.</param>

// <param name="Job" type="CIM_ConcreteJob ">If operation is long running then optionally a job may be returned.</param>
// <param name="NewComputerSystem" type="CIM_ComputerSystem ">Reference to an instance of the CIM_ComputerSystem class representing the virtual computer system after it has been migrated.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemMigrationService) MigrateVirtualSystemToSystem( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ DestinationSystem CIM_System,
	/* IN */ MigrationSettingData string,
	/* IN */ NewSystemSettingData string,
	/* IN */ NewResourceSettingData []string,
	/* OUT */ NewComputerSystem CIM_ComputerSystem,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("MigrateVirtualSystemToSystem", Action, PercentComplete, Timeout, ComputerSystem, DestinationSystem, MigrationSettingData, NewSystemSettingData, NewResourceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Method to perform a pre-check to determine whether a virtual system is likely to be successfully migrated to a target host specified by a network name or IP address. This method does not guarantee that a subsequent migration will always succeed, due to dynamic resource availability.
///Return code description:
///0: Success: Check performed; result reported through the value of the [Out] IsMigratable parameter.
///1: Error: Method not supported by implementation. No result reported through the value of the [Out] IsMigratable parameter.
///2: Error: Check failed for unspecified reasons. No result reported through the value of the [Out] IsMigratable parameter.
///3: Error: Check timed out. No result reported through the value of the [Out] IsMigratable parameter.
///4: Error: One or more parameters are formally invalid. For example, the value of the DestinationHost parameter could have been specified in an unsupported format.
///No result reported through the value of the [Out] IsMigratable parameter.
///5: Error: The source virtual system, the source host system or the target host system are in a state that does allow initiation of the requested virtual system migration; this may be a temporary condition. No result reported reported through the value of the [Out] IsMigratable parameter.
///6: Error: One or more input parameters are incompatible as a set, or with respect to the target host. For example the value of the MigrationNewSettingData parameter contains properties that are not supported by the target host system identified by the value of the DestinationHost parameter. No result reported through the value of the [Out] IsMigratable parameter.
///Note: The CheckVirtualSystemIsMigratableToHost( ) method is intended as a transitional solution only until modelling of cluster support is available.

// <param name="ComputerSystem" type="CIM_ComputerSystem ">Source virtual computer system to be migrated.</param>
// <param name="DestinationHost" type="string ">Target host system for the migration.
///Acceptable formats for this parameter are conveyed through values of elements of the DestinationHostFormatsSupported[ ] array property in the instance of the CIM_VirtualSystemMigrationCapabilities that is associated through the CIM_ElementCapabilities assocation.</param>
// <param name="MigrationSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemMigrationSettingData class representing migration settings applicable to the migration operation.</param>
// <param name="NewResourceSettingData" type="string []">Array of strings each containing an embedded instance of the CIM_ResourceAllocationSettingData class representing new properties applicable to virtual resources in the scope of the virtual system after it is migrated.</param>
// <param name="NewSystemSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemSettingData class representing new properties applicable to the virtual system after it is migrated.</param>

// <param name="IsMigratable" type="bool ">The migration check result indicating whether or not the virtual system can be successfully migrated.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemMigrationService) CheckVirtualSystemIsMigratableToHost( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ DestinationHost string,
	/* IN */ MigrationSettingData string,
	/* IN */ NewSystemSettingData string,
	/* IN */ NewResourceSettingData []string,
	/* OUT */ IsMigratable bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckVirtualSystemIsMigratableToHost", ComputerSystem, DestinationHost, MigrationSettingData, NewSystemSettingData, NewResourceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Method to perform a pre-check to determine whether a virtual system is likely to be successfully migrated to a target system. This method does not guarantee that a subsequent migration will always succeed, due to dynamic resource availability. Return code description:
///0: Success: Check performed; result reported through the value of the [Out] IsMigratable parameter.
///1: Error: Method not supported by implementation. No result reported through the value of the [Out] IsMigratable parameter.
///2: Error: Check failed for unspecified reasons. No result reported through the value of the [Out] IsMigratable parameter.
///3: Error: Check timed out. No result reported through the value of the [Out] IsMigratable parameter.
///4: Error: One or more parameters are formally invalid. For example, the value of the DestinationSystem parameter does not comprise a valid object path. No result reported through the value of the [Out] IsMigratable parameter.
///5: Error: The source virtual system, the source host system or the target host system are in a state that does allow initiation of the requested virtual system migration; this may be a temporary condition. No result reported reported through the value of the [Out] IsMigratable parameter.
///6: Error: One or more input parameters are incompatible as a set, or with respect to the target host. For example the value of the NewSettingData parameter contains properties that are not supported by the target host system identified by the value of the DestinationSystem parameter. No result reported through the value of the [Out] IsMigratable parameter.

// <param name="ComputerSystem" type="CIM_ComputerSystem ">Source virtual computer system to be migrated.</param>
// <param name="DestinationSystem" type="CIM_System ">Destination system onto which to migrate the virtual system.</param>
// <param name="MigrationSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemMigrationSettingData class representing migration settings applicable to the migration operation.</param>
// <param name="NewResourceSettingData" type="string []">Array of strings each containing an embedded instance of the CIM_ResourceAllocationSettingData class representing new properties applicable to virtual resources in the scope of the virtual system after it is migrated.</param>
// <param name="NewSystemSettingData" type="string ">String containing an embedded instance of the CIM_VirtualSystemSettingData class representing new properties applicable to the virtual system after it is migrated.</param>

// <param name="IsMigratable" type="bool ">The migration check result indicating whether or not the virtual system can be successfully migrated.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_VirtualSystemMigrationService) CheckVirtualSystemIsMigratableToSystem( /* IN */ ComputerSystem CIM_ComputerSystem,
	/* IN */ DestinationSystem CIM_System,
	/* IN */ MigrationSettingData string,
	/* IN */ NewSystemSettingData string,
	/* IN */ NewResourceSettingData []string,
	/* OUT */ IsMigratable bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckVirtualSystemIsMigratableToSystem", ComputerSystem, DestinationSystem, MigrationSettingData, NewSystemSettingData, NewResourceSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
