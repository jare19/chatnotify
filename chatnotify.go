package chatnotify

// Payloader is an interface for sending chatops messages
// useable packages are currently:
// "github.com/jare19/chatnotify/chatops/hipchat"
// "github.com/jare19/chatnotify/chatops/slack"
type Payloader interface {
	SendMessage() error
}

// Send takes a Payloader and Sends the chatops message.
func Send(p Payloader) error {
	err := p.SendMessage()
	if err != nil {
		return err
	}
	return nil
}
