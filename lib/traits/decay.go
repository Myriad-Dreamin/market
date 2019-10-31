package traits

type DecayResult struct {
	First interface{}
	Err   error
}

func Decay(d interface{}, e error) DecayResult {
	return DecayResult{d, e}
}

type DecayResult2 struct {
	First interface{}
	Second interface{}
	Err   error
}

func Decay2(d, d2 interface{}, e error) DecayResult2 {
	return DecayResult2{d, d2, e}
}


type DecayResult3 struct {
	First interface{}
	Second interface{}
	Third interface{}
	Err   error
}

func Decay3(d, d2, d3 interface{}, e error) DecayResult3 {
	return DecayResult3{d, d2, d3, e}
}

