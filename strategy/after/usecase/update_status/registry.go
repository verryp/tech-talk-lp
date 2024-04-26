package update_status

type (
	Strategy map[string]IOrderStatus
)

func Registry(ords ...IOrderStatus) Strategy {
	o := Strategy{}

	for _, ord := range ords {
		o.Set(ord)
	}

	return o
}

func (s Strategy) Set(o IOrderStatus) bool {
	_, ok := s[o.Type()]

	if ok {
		return false
	}

	s[o.Type()] = o

	return true
}

func (s Strategy) Get(orderStatus string) IOrderStatus {
	o, ok := s[orderStatus]
	if !ok {
		return nil
	}

	return o
}
