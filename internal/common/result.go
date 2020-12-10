package common

func Output(data interface{}, err error) (int, interface{}) {
	if err != nil {
		return 200, data
	}
	e, t := err.(Error)
	if t {
		return e.StatusCode(), e
	}
	se := SystemError(err.Error())
	return se.StatusCode(), err

}
