package sofarpc

import (
	"gitlab.alipay-inc.com/afe/mosn/pkg/types"
	"gitlab.alipay-inc.com/afe/mosn/pkg/protocol/sofarpc"
	str "gitlab.alipay-inc.com/afe/mosn/pkg/stream"
	"gitlab.alipay-inc.com/afe/mosn/pkg/protocol"
)

func init() {
	str.Register(protocol.SofaRpc, &streamConnFactory{})
}

type streamConnFactory struct{}

func (f *streamConnFactory) CreateClientStream(connection types.ClientConnection,
	streamConnCallbacks types.StreamConnectionCallbacks, connCallbacks types.ConnectionCallbacks) types.ClientStreamConnection {
	return newClientStreamConnection(connection, streamConnCallbacks)
}

func (f *streamConnFactory) CreateServerStream(connection types.Connection,
	callbacks types.ServerStreamConnectionCallbacks) types.ServerStreamConnection {
	return newServerStreamConnection(connection, callbacks)
}

// types.DecodeFilter
// types.StreamConnection
type streamConnection struct {
	protocol      types.Protocol
	connection    types.Connection
	activeStreams map[uint32]*stream
	protocols     types.Protocols
}

// types.StreamConnection
func (conn *streamConnection) Dispatch(buffer types.IoBuffer) {
	conn.protocols.Decode(buffer, conn)

}

func (conn *streamConnection) Protocol() types.Protocol {
	return conn.protocol
}

// types.DecodeFilter
func (conn *streamConnection) OnDecodeHeader(streamId uint32, headers map[string]string) types.FilterStatus {
	if stream, ok := conn.activeStreams[streamId]; ok {
		stream.decoder.OnDecodeHeaders(headers, false)   //回调PROXY层的OnDecodeHeaders
	}

	return types.Continue
}

func (conn *streamConnection) OnDecodeData(streamId uint32, data types.IoBuffer) types.FilterStatus {
	if stream, ok := conn.activeStreams[streamId]; ok {
		stream.decoder.OnDecodeData(data, true)
	}

	return types.Continue
}

func (conn *streamConnection) OnDecodeTrailer(streamId uint32, trailers map[string]string) types.FilterStatus {
	if stream, ok := conn.activeStreams[streamId]; ok {
		stream.decoder.OnDecodeTrailers(trailers)
	}

	return types.Continue
}

func (conn *streamConnection) OnDecodeComplete(streamId uint32, buf types.IoBuffer) {
	if stream, ok := conn.activeStreams[streamId]; ok {
		stream.decoder.OnDecodeComplete(buf)
	}
}

// types.ClientStreamConnection
type clientStreamConnection struct {
	streamConnection
	streamConnCallbacks types.StreamConnectionCallbacks
}

func newClientStreamConnection(connection types.Connection,
	callbacks types.StreamConnectionCallbacks) types.ClientStreamConnection {

	return &clientStreamConnection{
		streamConnection: streamConnection{
			connection:    connection,
			protocols:     sofarpc.DefaultProtocols(),
			activeStreams: make(map[uint32]*stream),
		},
		streamConnCallbacks: callbacks,
	}
}

func (c *clientStreamConnection) NewStream(streamId uint32, responseDecoder types.StreamDecoder) types.StreamEncoder {
	stream := &stream{
		streamId:   streamId,
		connection: &c.streamConnection,
		decoder:    responseDecoder,
	}

	c.activeStreams[streamId] = stream

	return stream
}

// types.ServerStreamConnection
type serverStreamConnection struct {
	streamConnection
	serverStreamConnCallbacks types.ServerStreamConnectionCallbacks
}

func newServerStreamConnection(connection types.Connection,
	callbacks types.ServerStreamConnectionCallbacks) types.ServerStreamConnection {
	return &serverStreamConnection{
		streamConnection: streamConnection{
			connection:    connection,
			protocols:     sofarpc.DefaultProtocols(),
			activeStreams: make(map[uint32]*stream),
		},
		serverStreamConnCallbacks: callbacks,
	}
}

func (sc *serverStreamConnection) Dispatch(buffer types.IoBuffer) {
	sc.protocols.Decode(buffer, sc)  //调用协议的decode，在decode中调用PROXY层的NEW STREAM作为一种通告机制，将返回STREAM DECODER



}

// types.DecodeFilter
func (sc *serverStreamConnection) OnDecodeHeader(streamId uint32, headers map[string]string) types.FilterStatus {
	if streamId == 0 {
		return types.Continue
	}

	//把 map[string]string 传到这个位置
	sc.onNewStreamDetected(streamId)   //创建NEW STREAM

	sc.streamConnection.OnDecodeHeader(streamId, headers)   //间接回调 PROXY层的 接口

	return types.StopIteration
}

func (sc *serverStreamConnection) OnDecodeData(streamId uint32, data types.IoBuffer) types.FilterStatus {
	if streamId == 0 {
		return types.Continue
	}    //

	sc.onNewStreamDetected(streamId)
	sc.streamConnection.OnDecodeData(streamId, data)

	return types.StopIteration
}

func (sc *serverStreamConnection) onNewStreamDetected(streamId uint32) {
	if _, ok := sc.activeStreams[streamId]; ok {
		return
	}

	stream := &stream{
		streamId:   streamId,
		connection: &sc.streamConnection,
	}

	//调用PROXY中定义的NEWSTREA，同时将 NEW出来的 STREAM作为 encoder 传进去
	stream.decoder = sc.serverStreamConnCallbacks.NewStream(streamId, stream)
	sc.activeStreams[streamId] = stream
}

// types.Stream
// types.StreamEncoder
type stream struct {
	streamId         uint32
	readDisableCount int
	connection       *streamConnection
	decoder          types.StreamDecoder
	streamCbs        []types.StreamCallbacks
}

// ~~ types.Stream
func (s *stream) AddCallbacks(cb types.StreamCallbacks) {
	s.streamCbs = append(s.streamCbs, cb)
}

func (s *stream) RemoveCallbacks(cb types.StreamCallbacks) {
	cbIdx := -1

	for i, streamCb := range s.streamCbs {
		if streamCb == cb {
			cbIdx = i
			break
		}
	}

	if cbIdx > -1 {
		s.streamCbs = append(s.streamCbs[:cbIdx], s.streamCbs[cbIdx+1:]...)
	}
}

func (s *stream) ResetStream(reason types.StreamResetReason) {
	for _, cb := range s.streamCbs {
		cb.OnResetStream(reason)
	}
}

func (s *stream) ReadDisable(disable bool) {
	s.connection.connection.SetReadDisable(disable)
}

func (s *stream) BufferLimit() uint32 {
	return s.connection.connection.BufferLimit()
}

// types.StreamEncoder
func (s *stream) EncodeHeaders(headers map[string]string, endStream bool) {
	// todo: encode headers  由stream层调用，将HEADERS 按照BOLT 的格式进行分装

	if endStream {
		s.endStream()
	}
}

func (s *stream) EncodeData(data types.IoBuffer, endStream bool) {
	// todo: encode data

	if endStream {
		s.endStream()
	}
}

func (s *stream) EncodeTrailers(trailers map[string]string) {
	// todo: encode trailer

	s.endStream()
}

func (s *stream) endStream() {
	// todo: flush stream data

	delete(s.connection.activeStreams, s.streamId)
}

func (s *stream) GetStream() types.Stream {
	return s
}