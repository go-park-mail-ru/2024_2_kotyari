package notifications

func (n *NotificationsApp) Run() {
	n.log.Info("[NotificationsApp.Run] Started worker")
	n.orderStatusesChanger.Run()
}
