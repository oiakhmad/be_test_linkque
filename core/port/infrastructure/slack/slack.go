package slack

type Slack interface {
	SendAlertSlack(message string, mention string) error
}
