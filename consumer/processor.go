package consumer

import "fmt"

type IMessageProcessor interface {
	ProcessMessage(*NotificationMessage) error
	ClearLocalCache()
}

type processor struct {
	msgMap map[string]*NotificationMessage
}

func NewProcessor() IMessageProcessor {
	return &processor{
		msgMap: make(map[string]*NotificationMessage),
	}
}

func (p *processor) ProcessMessage(msg *NotificationMessage) error {
	if _, ok := p.msgMap[msg.Key]; ok {
		fmt.Println(p.msgMap[msg.Key].Topic)
		return nil
	}
	p.msgMap[msg.Key] = msg
	return nil
}

func (p *processor) ClearLocalCache() {
	// TODO: implement me
	p.msgMap = make(map[string]*NotificationMessage)
}
