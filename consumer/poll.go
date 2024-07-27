package consumer

import "context"

type Poll interface {
	Start()
	Stop()
}

type poll struct {
	l      int
	r      int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPoll() Poll {
	ctx, cancel := context.WithCancel(context.TODO())
	return &poll{
		l:      0,
		r:      1,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (p *poll) Start() {
	processor := NewProcessor()

	for _, msg := range GetMessages() {
		select {
		case <-p.ctx.Done():
			return
		default:
			go processor.ProcessMessage(&msg)
		}
	}

	processor.ClearLocalCache()
	p.Stop()
}

func (p *poll) Stop() {
	p.cancel()
}
