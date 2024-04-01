package builder

import "errors"

type Message struct {
	Headers map[string]string
	Body    []byte
}

type Response struct {
	ResultCode    int
	ResultMessage *Message
}

type builder struct {
	msgType  string
	handlers []Handler
}

type Handler func(m *Message) *Response

func New() *builder {
	return &builder{}
}

func (b *builder) WithType(t string) *builder {
	b.msgType = t
	return b
}

func (b *builder) WithHandler(h Handler) *builder {
	b.handlers = append(b.handlers, h)
	return b
}

type Builder interface {
	WithType(t string) Builder
	WithHandler(h Handler) Builder
	Build() HandlerService
}

func (b *builder) toFinishedService() (*finishedService, error) {
	if b.msgType == "" {
		return nil, errors.New("Services need a type!")
	}

	return &finishedService{
		msgType:  b.msgType,
		handlers: b.handlers,
	}, nil
}

func (b *builder) Build() (HandlerService, error) {
	return b.toFinishedService()
}

type finishedService struct {
	msgType  string
	handlers []Handler
}

type HandlerService interface {
	Handle(m *Message) *Response
}

func (fs *finishedService) Handle(m *Message) *Response {
	for _, h := range fs.handlers {
		if m.Headers["type"] == fs.msgType {
			return h(m)
		}
	}
	return &Response{
		ResultCode:    500,
		ResultMessage: nil,
	}
}
