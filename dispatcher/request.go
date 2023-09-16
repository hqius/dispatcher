package dispatcher

type Request struct {
	data map[string]interface{}
}

func NewRequest() *Request {
	return &Request{
		data: make(map[string]interface{}),
	}
}

func NewMetaRequest(meta map[string]interface{}) *Request {
	return &Request{
		data: meta,
	}
}

func (r *Request) Data() map[string]interface{} {
	return r.data
}

func (r *Request) Get(key string) interface{} {
	if r.data == nil {
		return nil
	}
	if val, ok := r.data[key]; ok {
		return val
	}
	return nil
}

func (r *Request) Put(key string, val interface{}) {
	if r.data == nil {
		panic("data is empty")
	}
	r.data[key] = val
}
