package dispatcher

import "context"

type DptCtx struct {
	biz      string
	state    State
	metadata map[string]interface{}
	ctx      context.Context
}

func NewDptCtx(biz string, state State) DptCtx {
	return DptCtx{
		biz:      biz,
		state:    state,
		metadata: make(map[string]interface{}),
		ctx:      context.Background(),
	}
}

func (d *DptCtx) Biz() string {
	return d.biz
}

func (d *DptCtx) State() State {
	return d.state
}

func (d *DptCtx) Get(key string) interface{} {
	if d.metadata == nil {
		return nil
	}
	if val, ok := d.metadata[key]; ok {
		return val
	}
	return nil
}

func (d *DptCtx) Put(key string, val interface{}) {
	if d.metadata == nil {
		panic("data is empty")
	}
	d.metadata[key] = val
}
