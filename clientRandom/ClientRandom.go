package clientRandom

import (
	"encoding/binary"
	"math/rand"
	"time"
)

type ClientRandom struct {
	gmt_unix_time uint32
	random_bytes  []byte // 28 bytes long
}

func NewClientRandom() ClientRandom {
	var random = ClientRandom{
		gmt_unix_time: uint32(time.Now().Unix()),
		random_bytes:  make([]byte, 28)}

	rand.Read(random.random_bytes)

	return random
}

/*
	Returns the total size in bytes of this struct
*/
func (random ClientRandom) GetSize() int {
	return 32
}

/*
	Serializes this struct into a given buffer, which is assumed to be 32 bytes.
*/
func (random ClientRandom) SerializeInto(buf []byte) {
	binary.BigEndian.PutUint32(buf[0:4], random.gmt_unix_time)
	copy(buf[4:31], random.random_bytes)
}
