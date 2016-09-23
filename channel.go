package logrusChannel

import "github.com/Sirupsen/logrus"

// ChannelHook is a hook for logrus
type ChannelHook struct {
	Channel chan *logrus.Entry
}

// NewChannelHook creates a hook to log to a channel
func NewChannelHook(channel chan *logrus.Entry) *ChannelHook {
	return &ChannelHook{
		Channel: channel,
	}
}

// Fire is called whan a log event is fired
func (hook *ChannelHook) Fire(entry *logrus.Entry) error {
	hook.Channel <- entry
	return nil
}
