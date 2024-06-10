// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.TaskScheduler
//
// ////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_ScheduledTask struct
type PS_ScheduledTask struct {
	*cim.WmiInstance
}

func NewPS_ScheduledTaskEx1(instance *cim.WmiInstance) (newInstance *PS_ScheduledTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_ScheduledTask{
		WmiInstance: tmp,
	}
	return
}

func NewPS_ScheduledTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_ScheduledTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_ScheduledTask{
		WmiInstance: tmp,
	}
	return
}

// 22

// <param name="Force" type="bool ">23</param>
// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>
// <param name="Password" type="string ">25</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="User" type="string ">26</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask ">29</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) RegisterByObject( /* IN */ Force bool,
	/* IN */ InputObject MSFT_ScheduledTask,
	/* IN */ Password string,
	/* IN */ User string,
	/* IN */ TaskName string,
	/* IN */ TaskPath string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByObject", Force, InputObject, Password, User, TaskName, TaskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 22

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Description" type="string ">32</param>
// <param name="Force" type="bool ">23</param>
// <param name="Principal" type="MSFT_TaskPrincipal ">30</param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="Trigger" type="MSFT_TaskTrigger []">34</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask ">29</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) RegisterByPrincipal( /* IN */ Force bool,
	/* IN */ Principal MSFT_TaskPrincipal,
	/* IN */ Action []MSFT_TaskAction,
	/* IN */ Description string,
	/* IN */ TaskPath string,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByPrincipal", Force, Principal, Action, Description, TaskPath, Settings, Trigger, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 22

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Description" type="string ">32</param>
// <param name="Force" type="bool ">23</param>
// <param name="Password" type="string ">25</param>
// <param name="RunLevel" type="int32 "></param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="Trigger" type="MSFT_TaskTrigger []">34</param>
// <param name="User" type="string ">26</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask ">29</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) RegisterByUser( /* IN */ Force bool,
	/* IN */ TaskPath string,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Description string,
	/* IN */ User string,
	/* IN */ Password string,
	/* IN */ Action []MSFT_TaskAction,
	/* IN */ RunLevel int32,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByUser", Force, TaskPath, Trigger, Settings, Description, User, Password, Action, RunLevel, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 22

// <param name="Force" type="bool ">23</param>
// <param name="Password" type="string ">25</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="User" type="string ">26</param>
// <param name="Xml" type="string ">35</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask ">29</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) RegisterByXml( /* IN */ Force bool,
	/* IN */ Xml string,
	/* IN */ Password string,
	/* IN */ User string,
	/* IN */ TaskPath string,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByXml", Force, Xml, Password, User, TaskPath, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 36

// <param name="Argument" type="string ">39</param>
// <param name="Execute" type="string ">38</param>
// <param name="Id" type="string ">37</param>
// <param name="WorkingDirectory" type="string ">40</param>

// <param name="cmdletOutput" type="MSFT_TaskAction "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewActionByExec( /* IN */ Id string,
	/* IN */ Execute string,
	/* IN */ Argument string,
	/* IN */ WorkingDirectory string,
	/* OUT */ cmdletOutput MSFT_TaskAction) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewActionByExec", Id, Execute, Argument, WorkingDirectory)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 41

// <param name="GroupId" type="string ">42</param>
// <param name="Id" type="string ">43</param>
// <param name="ProcessTokenSidType" type="int32 ">45</param>
// <param name="RequiredPrivilege" type="string []">46</param>
// <param name="RunLevel" type="int32 ">44</param>

// <param name="cmdletOutput" type="MSFT_TaskPrincipal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewPrincipalByGroup( /* IN */ GroupId string,
	/* IN */ Id string,
	/* IN */ RunLevel int32,
	/* IN */ ProcessTokenSidType int32,
	/* IN */ RequiredPrivilege []string,
	/* OUT */ cmdletOutput MSFT_TaskPrincipal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewPrincipalByGroup", GroupId, Id, RunLevel, ProcessTokenSidType, RequiredPrivilege)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 41

// <param name="Id" type="string ">43</param>
// <param name="LogonType" type="int32 "></param>
// <param name="ProcessTokenSidType" type="int32 ">45</param>
// <param name="RequiredPrivilege" type="string []">46</param>
// <param name="RunLevel" type="int32 ">44</param>
// <param name="UserId" type="string ">47</param>

// <param name="cmdletOutput" type="MSFT_TaskPrincipal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewPrincipalByUser( /* IN */ UserId string,
	/* IN */ LogonType int32,
	/* IN */ Id string,
	/* IN */ ProcessTokenSidType int32,
	/* IN */ RequiredPrivilege []string,
	/* IN */ RunLevel int32,
	/* OUT */ cmdletOutput MSFT_TaskPrincipal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewPrincipalByUser", UserId, LogonType, Id, ProcessTokenSidType, RequiredPrivilege, RunLevel)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 48

// <param name="AllowStartIfOnBatteries" type="bool ">53</param>
// <param name="Compatibility" type="int32 ">51</param>
// <param name="DeleteExpiredTaskAfter" type="string ">52</param>
// <param name="Disable" type="bool ">54</param>
// <param name="DisallowDemandStart" type="bool ">49</param>
// <param name="DisallowHardTerminate" type="bool ">50</param>
// <param name="DisallowStartOnRemoteAppSession" type="bool ">61</param>
// <param name="DontStopIfGoingOnBatteries" type="bool ">65</param>
// <param name="DontStopOnIdleEnd" type="bool ">69</param>
// <param name="ExecutionTimeLimit" type="string ">70</param>
// <param name="Hidden" type="bool ">56</param>
// <param name="IdleDuration" type="string ">67</param>
// <param name="IdleWaitTimeout" type="string ">58</param>
// <param name="MaintenanceDeadline" type="string ">63</param>
// <param name="MaintenanceExclusive" type="bool ">55</param>
// <param name="MaintenancePeriod" type="string ">62</param>
// <param name="MultipleInstances" type="int32 ">71</param>
// <param name="NetworkId" type="string ">59</param>
// <param name="NetworkName" type="string ">60</param>
// <param name="Priority" type="int32 ">72</param>
// <param name="RestartCount" type="int32 ">73</param>
// <param name="RestartInterval" type="string ">74</param>
// <param name="RestartOnIdle" type="bool ">68</param>
// <param name="RunOnlyIfIdle" type="bool ">57</param>
// <param name="RunOnlyIfNetworkAvailable" type="bool ">75</param>
// <param name="StartWhenAvailable" type="bool ">64</param>
// <param name="WakeToRun" type="bool ">66</param>

// <param name="cmdletOutput" type="MSFT_TaskSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewSettings( /* IN */ DisallowDemandStart bool,
	/* IN */ DisallowHardTerminate bool,
	/* IN */ Compatibility int32,
	/* IN */ DeleteExpiredTaskAfter string,
	/* IN */ AllowStartIfOnBatteries bool,
	/* IN */ Disable bool,
	/* IN */ MaintenanceExclusive bool,
	/* IN */ Hidden bool,
	/* IN */ RunOnlyIfIdle bool,
	/* IN */ IdleWaitTimeout string,
	/* IN */ NetworkId string,
	/* IN */ NetworkName string,
	/* IN */ DisallowStartOnRemoteAppSession bool,
	/* IN */ MaintenancePeriod string,
	/* IN */ MaintenanceDeadline string,
	/* IN */ StartWhenAvailable bool,
	/* IN */ DontStopIfGoingOnBatteries bool,
	/* IN */ WakeToRun bool,
	/* IN */ IdleDuration string,
	/* IN */ RestartOnIdle bool,
	/* IN */ DontStopOnIdleEnd bool,
	/* IN */ ExecutionTimeLimit string,
	/* IN */ MultipleInstances int32,
	/* IN */ Priority int32,
	/* IN */ RestartCount int32,
	/* IN */ RestartInterval string,
	/* IN */ RunOnlyIfNetworkAvailable bool,
	/* OUT */ cmdletOutput MSFT_TaskSettings) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewSettings", DisallowDemandStart, DisallowHardTerminate, Compatibility, DeleteExpiredTaskAfter, AllowStartIfOnBatteries, Disable, MaintenanceExclusive, Hidden, RunOnlyIfIdle, IdleWaitTimeout, NetworkId, NetworkName, DisallowStartOnRemoteAppSession, MaintenancePeriod, MaintenanceDeadline, StartWhenAvailable, DontStopIfGoingOnBatteries, WakeToRun, IdleDuration, RestartOnIdle, DontStopOnIdleEnd, ExecutionTimeLimit, MultipleInstances, Priority, RestartCount, RestartInterval, RunOnlyIfNetworkAvailable)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 76

// <param name="InputObject" type="MSFT_ScheduledTask ">77</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) StartByObject( /* IN */ InputObject MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 76

// <param name="TaskName" type="string ">79</param>
// <param name="TaskPath" type="string ">78</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) StartByPath( /* IN */ TaskPath string,
	/* IN */ TaskName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartByPath", TaskPath, TaskName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 80

// <param name="InputObject" type="MSFT_ScheduledTask ">81</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) StopByObject( /* IN */ InputObject MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 80

// <param name="TaskName" type="string ">83</param>
// <param name="TaskPath" type="string ">82</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) StopByPath( /* IN */ TaskPath string,
	/* IN */ TaskName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopByPath", TaskPath, TaskName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 84

// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>
// <param name="Password" type="string ">25</param>
// <param name="User" type="string ">26</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) SetByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* IN */ Password string,
	/* IN */ User string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByObject", InputObject, Password, User)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 84

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Principal" type="MSFT_TaskPrincipal ">85</param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="Trigger" type="MSFT_TaskTrigger []">86</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) SetByPrincipal( /* IN */ Principal MSFT_TaskPrincipal,
	/* IN */ Action []MSFT_TaskAction,
	/* IN */ TaskPath string,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByPrincipal", Principal, Action, TaskPath, Settings, Trigger, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 84

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Password" type="string ">25</param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string ">27</param>
// <param name="TaskPath" type="string ">28</param>
// <param name="Trigger" type="MSFT_TaskTrigger []">86</param>
// <param name="User" type="string ">26</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) SetByUser( /* IN */ Action []MSFT_TaskAction,
	/* IN */ TaskPath string,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ Password string,
	/* IN */ User string,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByUser", Action, TaskPath, Settings, Trigger, Password, User, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 87

// <param name="TaskName" type="string ">88</param>
// <param name="TaskPath" type="string ">89</param>

// <param name="cmdletOutput" type="MSFT_TaskDynamicInfo "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) GetInfoByName( /* IN */ TaskName string,
	/* IN */ TaskPath string,
	/* OUT */ cmdletOutput MSFT_TaskDynamicInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetInfoByName", TaskName, TaskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 87

// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>

// <param name="cmdletOutput" type="MSFT_TaskDynamicInfo "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) GetInfoByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* OUT */ cmdletOutput MSFT_TaskDynamicInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetInfoByObject", InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 90

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Description" type="string ">32</param>
// <param name="Principal" type="MSFT_TaskPrincipal ">85</param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="Trigger" type="MSFT_TaskTrigger []">86</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) New( /* IN */ Action []MSFT_TaskAction,
	/* IN */ Description string,
	/* IN */ Principal MSFT_TaskPrincipal,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("New", Action, Description, Principal, Settings, Trigger)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 91

// <param name="TaskName" type="string ">92</param>
// <param name="TaskPath" type="string ">93</param>

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) ExportByName( /* IN */ TaskName string,
	/* IN */ TaskPath string,
	/* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ExportByName", TaskName, TaskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 91

// <param name="InputObject" type="MSFT_ScheduledTask ">94</param>

// <param name="cmdletOutput" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) ExportByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* OUT */ cmdletOutput string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ExportByObject", InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 95

// <param name="TaskName" type="string ">92</param>
// <param name="TaskPath" type="string ">89</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) EnableByName( /* IN */ TaskName string,
	/* IN */ TaskPath string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByName", TaskName, TaskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 95

// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) EnableByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnableByObject", InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 96

// <param name="TaskName" type="string ">92</param>
// <param name="TaskPath" type="string ">89</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) DisableByName( /* IN */ TaskName string,
	/* IN */ TaskPath string,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByName", TaskName, TaskPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 96

// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>

// <param name="cmdletOutput" type="MSFT_ScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) DisableByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* OUT */ cmdletOutput MSFT_ScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DisableByObject", InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 99

// <param name="At" type="string ">101</param>
// <param name="Daily" type="bool ">99</param>
// <param name="DaysInterval" type="uint32 ">100</param>
// <param name="RandomDelay" type="string ">98</param>

// <param name="cmdletOutput" type="MSFT_TaskTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewTriggerByDaily( /* IN */ Daily bool,
	/* IN */ DaysInterval uint32,
	/* IN */ RandomDelay string,
	/* IN */ At string,
	/* OUT */ cmdletOutput MSFT_TaskTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewTriggerByDaily", Daily, DaysInterval, RandomDelay, At)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 105

// <param name="At" type="string ">101</param>
// <param name="Once" type="bool ">105</param>
// <param name="RandomDelay" type="string ">98</param>
// <param name="RepetitionDuration" type="string ">106</param>
// <param name="RepetitionInterval" type="string ">107</param>

// <param name="cmdletOutput" type="MSFT_TaskTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewTriggerByOnce( /* IN */ Once bool,
	/* IN */ RandomDelay string,
	/* IN */ At string,
	/* IN */ RepetitionDuration string,
	/* IN */ RepetitionInterval string,
	/* OUT */ cmdletOutput MSFT_TaskTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewTriggerByOnce", Once, RandomDelay, At, RepetitionDuration, RepetitionInterval)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 103

// <param name="AtLogOn" type="bool ">103</param>
// <param name="RandomDelay" type="string ">98</param>
// <param name="User" type="string ">104</param>

// <param name="cmdletOutput" type="MSFT_TaskTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewTriggerByLogon( /* IN */ RandomDelay string,
	/* IN */ AtLogOn bool,
	/* IN */ User string,
	/* OUT */ cmdletOutput MSFT_TaskTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewTriggerByLogon", RandomDelay, AtLogOn, User)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 108

// <param name="AtStartup" type="bool ">108</param>
// <param name="RandomDelay" type="string ">98</param>
// <param name="User" type="string ">104</param>

// <param name="cmdletOutput" type="MSFT_TaskTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewTriggerByStartup( /* IN */ RandomDelay string,
	/* IN */ AtStartup bool,
	/* IN */ User string,
	/* OUT */ cmdletOutput MSFT_TaskTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewTriggerByStartup", RandomDelay, AtStartup, User)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 110

// <param name="At" type="string ">101</param>
// <param name="DaysOfWeek" type="int32 []">109</param>
// <param name="RandomDelay" type="string ">98</param>
// <param name="Weekly" type="bool ">110</param>
// <param name="WeeksInterval" type="uint32 ">100</param>

// <param name="cmdletOutput" type="MSFT_TaskTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ScheduledTask) NewTriggerByWeekly( /* IN */ RandomDelay string,
	/* IN */ DaysOfWeek []int32,
	/* IN */ Weekly bool,
	/* IN */ WeeksInterval uint32,
	/* IN */ At string,
	/* OUT */ cmdletOutput MSFT_TaskTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewTriggerByWeekly", RandomDelay, DaysOfWeek, Weekly, WeeksInterval, At)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
