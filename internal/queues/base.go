package queues

func Run() {
	ScheduledDeviceAlarmLogRun()
	ScheduledSysOperLogRun()
	TaskDeviceDataTsdSaveRun()
	DeviceInfoUpdateRun()
}
