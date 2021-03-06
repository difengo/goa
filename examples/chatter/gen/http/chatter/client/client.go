// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter client HTTP transport
//
// Command:
// $ goa gen goa.design/goa/examples/chatter/design -o
// $(GOPATH)/src/goa.design/goa/examples/chatter

package client

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	goa "goa.design/goa"
	chattersvc "goa.design/goa/examples/chatter/gen/chatter"
	chattersvcviews "goa.design/goa/examples/chatter/gen/chatter/views"
	goahttp "goa.design/goa/http"
)

// Client lists the chatter service endpoint HTTP clients.
type Client struct {
	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// Echoer Doer is the HTTP client used to make requests to the echoer endpoint.
	EchoerDoer goahttp.Doer

	// Listener Doer is the HTTP client used to make requests to the listener
	// endpoint.
	ListenerDoer goahttp.Doer

	// Summary Doer is the HTTP client used to make requests to the summary
	// endpoint.
	SummaryDoer goahttp.Doer

	// History Doer is the HTTP client used to make requests to the history
	// endpoint.
	HistoryDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme       string
	host         string
	encoder      func(*http.Request) goahttp.Encoder
	decoder      func(*http.Response) goahttp.Decoder
	dialer       goahttp.Dialer
	connConfigFn goahttp.ConnConfigureFunc
}

// echoerClientStream implements the chattersvc.EchoerClientStream interface.
type echoerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// listenerClientStream implements the chattersvc.ListenerClientStream
// interface.
type listenerClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// summaryClientStream implements the chattersvc.SummaryClientStream interface.
type summaryClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// historyClientStream implements the chattersvc.HistoryClientStream interface.
type historyClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render  result type before sending to the websocket
	// connection.
	view string
}

// NewClient instantiates HTTP clients for all the chatter service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
	dialer goahttp.Dialer,
	connConfigFn goahttp.ConnConfigureFunc,
) *Client {
	return &Client{
		LoginDoer:           doer,
		EchoerDoer:          doer,
		ListenerDoer:        doer,
		SummaryDoer:         doer,
		HistoryDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
		dialer:              dialer,
		connConfigFn:        connConfigFn,
	}
}

// Login returns an endpoint that makes HTTP requests to the chatter service
// login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("chatter", "login", err)
		}
		return decodeResponse(resp)
	}
}

// Echoer returns an endpoint that makes HTTP requests to the chatter service
// echoer server.
func (c *Client) Echoer() goa.Endpoint {
	var (
		encodeRequest  = EncodeEchoerRequest(c.encoder)
		decodeResponse = DecodeEchoerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildEchoerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		conn, resp, err := c.dialer.Dial(req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "echoer", err)
		}
		if c.connConfigFn != nil {
			conn = c.connConfigFn(conn)
		}
		stream := &echoerClientStream{conn: conn}
		return stream, nil
	}
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *echoerClientStream) Recv() (string, error) {
	var (
		rv   string
		body string
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	return body, nil
}

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *echoerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "echoer" endpoint websocket connection.
func (s *echoerClientStream) Close() error {
	defer s.conn.Close()
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return nil
}

// Listener returns an endpoint that makes HTTP requests to the chatter service
// listener server.
func (c *Client) Listener() goa.Endpoint {
	var (
		encodeRequest  = EncodeListenerRequest(c.encoder)
		decodeResponse = DecodeListenerResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListenerRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		conn, resp, err := c.dialer.Dial(req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "listener", err)
		}
		if c.connConfigFn != nil {
			conn = c.connConfigFn(conn)
		}
		stream := &listenerClientStream{conn: conn}
		return stream, nil
	}
}

// Send streams instances of "string" to the "listener" endpoint websocket
// connection.
func (s *listenerClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// Close closes the "listener" endpoint websocket connection.
func (s *listenerClientStream) Close() error {
	defer s.conn.Close()
	var err error
	// Send a nil payload to the server implying client closing connection.
	if err = s.conn.WriteJSON(nil); err != nil {
		return err
	}
	return nil
}

// Summary returns an endpoint that makes HTTP requests to the chatter service
// summary server.
func (c *Client) Summary() goa.Endpoint {
	var (
		encodeRequest  = EncodeSummaryRequest(c.encoder)
		decodeResponse = DecodeSummaryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSummaryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		conn, resp, err := c.dialer.Dial(req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "summary", err)
		}
		if c.connConfigFn != nil {
			conn = c.connConfigFn(conn)
		}
		stream := &summaryClientStream{conn: conn}
		return stream, nil
	}
}

// CloseAndRecv stops sending messages to the "summary" endpoint websocket
// connection and reads instances of "chattersvc.ChatSummaryCollection" from
// the connection.
func (s *summaryClientStream) CloseAndRecv() (chattersvc.ChatSummaryCollection, error) {
	var (
		rv   chattersvc.ChatSummaryCollection
		body SummaryResponseBody
		err  error
	)
	defer s.conn.Close()
	// Send a nil payload to the server implying end of message
	if err = s.conn.WriteJSON(nil); err != nil {
		return rv, err
	}
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewSummaryChatSummaryCollectionOK(body)
	vres := chattersvcviews.ChatSummaryCollection{res, "default"}
	if err := vres.Validate(); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "summary", err)
	}
	return chattersvc.NewChatSummaryCollection(vres), nil
}

// Send streams instances of "string" to the "summary" endpoint websocket
// connection.
func (s *summaryClientStream) Send(v string) error {
	return s.conn.WriteJSON(v)
}

// History returns an endpoint that makes HTTP requests to the chatter service
// history server.
func (c *Client) History() goa.Endpoint {
	var (
		encodeRequest  = EncodeHistoryRequest(c.encoder)
		decodeResponse = DecodeHistoryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHistoryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		conn, resp, err := c.dialer.Dial(req.URL.String(), req.Header)
		if err != nil {
			if resp != nil {
				return decodeResponse(resp)
			}
			return nil, goahttp.ErrRequestError("chatter", "history", err)
		}
		if c.connConfigFn != nil {
			conn = c.connConfigFn(conn)
		}
		stream := &historyClientStream{conn: conn}
		view := resp.Header.Get("goa-view")
		stream.SetView(view)
		return stream, nil
	}
}

// Recv reads instances of "chattersvc.ChatSummary" from the "history" endpoint
// websocket connection.
func (s *historyClientStream) Recv() (*chattersvc.ChatSummary, error) {
	var (
		rv   *chattersvc.ChatSummary
		body HistoryResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewHistoryChatSummaryOK(&body)
	vres := &chattersvcviews.ChatSummary{res, s.view}
	if err := vres.Validate(); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "history", err)
	}
	return chattersvc.NewChatSummary(vres), nil
}

// SetView sets the view to render the  type before sending to the "history"
// endpoint websocket connection.
func (s *historyClientStream) SetView(view string) {
	s.view = view
}
