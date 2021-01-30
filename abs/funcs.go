package lacia

/*func CreateChannel(t reflect.Type, bufferSize int) (v reflect.Value, err error) {
	if t.Kind() != reflect.Chan {
		err = errors.New("input type error")
		return
	}
	chanType := reflect.ChanOf(reflect.BothDir, t)
	v = reflect.MakeChan(chanType, bufferSize)
	return
}
*/