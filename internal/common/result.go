package common

func Output(data interface{}, err error) (int, interface{}) {
	if err == nil {
		return 200, data
	}
	return Output2(err)

}

func Output2(err error) (int, interface{}) {
	e, t := err.(Error)
	if t {
		return e.StatusCode(), e
	}
	e = SystemError.New(err.Error())
	return e.StatusCode(), e
}
