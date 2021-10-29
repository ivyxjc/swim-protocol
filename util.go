package swim

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"time"
)

func decode(buf []byte, out interface{}) error {
	r := bytes.NewBuffer(buf)
	dec := gob.NewDecoder(r)
	return dec.Decode(out)
}

func encode(msgType messageType, in interface{}) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.BigEndian, msgType); err != nil {
		return nil, err
	}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(in)
	return buf, err
}

func triggerFunc(C <-chan time.Time, stop <-chan struct{}, f func()) {
	// todo  use a random stagger to avoid sync
	for {
		select {
		case <-C:
			f()
		case <-stop:
			return
		}

	}
}
