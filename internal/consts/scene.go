package consts

const (
	SceneTypeManual = "manual"
	SceneTypeTimer  = "timer"
	SceneTypeDevice = "device"

	SceneDefinition = "definition"
	SceneActionConf = "action"

	// 场景触发类型
	SceneTriggerOnLine          = "onLine"
	SceneTriggerOffLine         = "offLine"
	SceneTriggerReportAttribute = "reportAttribute"
	SceneTriggerReportEvent     = "reportEvent"

	// 场景动作类型
	SceneActionTypeDeviceOutput       = "deviceOutput"
	SceneActionTypeSendNotice         = "sendNotice"
	SceneActionTypeCallWebService     = "callWebService"
	SceneActionTypeTriggerAlarm       = "triggerAlarm"
	SceneActionTypeDelayExecution     = "delayExecution"
	SceneActionTypeTriggerCustomEvent = "triggerCustomEvent"
)
