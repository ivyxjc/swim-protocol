package swim

type ping struct {
	SeqNo uint32

	Node string
}

type indirectPing struct {
}

type ackResp struct {
	SeqNo uint32
}

type alive struct {
}

type suspect struct {
}

type dead struct {
}
