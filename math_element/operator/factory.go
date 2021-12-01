package operator

func Create(source string) (instance Operator, ok bool) {
	instance = nil
	ok = false
	switch source {
	case "+":
		instance = pluse{}
	case "-":
		instance = minus{}
	case "*":
		instance = multiply{}
	case "/":
		instance = divide{}
	}

	if instance != nil {
		ok = true
	}

	return instance, ok
}
