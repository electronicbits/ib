package ib

import (
	"bytes"
)

// Request .
type Request interface {
	writable
	code() OutgoingMessageID
	version() int64
}

// Reply .
type Reply interface {
	readable
	code() IncomingMessageID
}

// MatchedRequest .
type MatchedRequest interface {
	Request
	SetId(id int64)
	ID() int64
}

// MatchedReply .
type MatchedReply interface {
	Reply
	ID() int64
}

type clientHandshake struct {
	version int64
	id      int64
}

func (c *clientHandshake) write(b *bytes.Buffer) error {
	if err := writeInt(b, c.version); err != nil {
		return err
	}
	if err := writeInt(b, mStartAPI); err != nil {
		return err
	}
	if err := writeInt(b, 1); err != nil {
		return err
	}
	return writeInt(b, c.id)
}
