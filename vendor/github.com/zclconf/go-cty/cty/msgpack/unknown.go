package msgpack

type myuserType struct{}

var myuserVal = myuserType{}

// myuserValBytes is the raw bytes of the msgpack fixext1 value we
// write to represent an myuser value. It's an extension value of
// type zero whose value is irrelevant. Since it's irrelevant, we
// set it to a single byte whose value is also zero, since that's
// the most compact possible representation.
var myuserValBytes = []byte{0xd4, 0, 0}

func (uv myuserType) MarshalMsgpack() ([]byte, error) {
	return myuserValBytes, nil
}
