package algorithm

//import "fmt"
import "math"

type Pid_t struct {
	Kp float64
	Ki float64
	Kd float64
	TargetTemperature float64

	err_history float64
	err_pre float64
	u float64

	ClampMin float64
	ClampMax float64

	DBKey string
}

func (o* Pid_t)Compute(currentTemperature float64) float64 {
	//Pid
	err := currentTemperature - o.TargetTemperature
	P := o.Kp * err
	//pId
	o.err_history = (o.err_history + err)/2
	I := o.Ki * o.err_history
	//piD
	D := o.Kd * (err - o.err_pre)

	val := o.u + P + I + D
	val = math.Floor(val + 0.5)

	if val <= o.ClampMin {
		val = o.ClampMin
	}
	if val >= o.ClampMax {
		val = o.ClampMax
	}

	o.err_pre = err
	o.u = val

	//fmt.Printf("PWM(%f)=%f + %f*%f + %f*%f + %f*(%f-%f)\n", val, o.u, o.Kp, err, o.Ki, o.err_history, o.Kd, err, o.err_pre)

	return val
}

