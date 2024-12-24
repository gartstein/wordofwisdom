package pkg

import (
	"bytes"
	"io"
	"net"
	"time"
)

type MockConn struct {
	readBuffer  *bytes.Buffer
	writeBuffer *bytes.Buffer
}

// NewMockConn create new mock for a connection.
func NewMockConn(input string) *MockConn {
	return &MockConn{
		readBuffer:  bytes.NewBufferString(input),
		writeBuffer: &bytes.Buffer{},
	}
}

// Output returns the data written to the connection.
func (mc *MockConn) Output() string {
	return mc.writeBuffer.String()
}

// Implement Read method
func (mc *MockConn) Read(b []byte) (int, error) {
	if mc.readBuffer.Len() == 0 {
		return 0, io.EOF
	}
	return mc.readBuffer.Read(b)
}

// Implement Write method
func (mc *MockConn) Write(b []byte) (int, error) {
	return mc.writeBuffer.Write(b)
}

// Implement Close method (no-op for mock)
func (mc *MockConn) Close() error {
	return nil
}

// Implement LocalAddr method (returning nil for simplicity)
func (mc *MockConn) LocalAddr() net.Addr {
	return nil
}

// Implement RemoteAddr method (returning nil for simplicity)
func (mc *MockConn) RemoteAddr() net.Addr {
	return nil
}

// Implement SetDeadline method (no-op for mock)
func (mc *MockConn) SetDeadline(t time.Time) error {
	return nil
}

// Implement SetReadDeadline method (no-op for mock)
func (mc *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

// Implement SetWriteDeadline method (no-op for mock)
func (mc *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
