package algorithm

//import "fmt"
import "math"

type Pid_t struct {
	Kp float32
	Ki float32
	Kd float32
	TargetTemperature float32

	err_history float32
	err_pre float32
	u float32

	ClampMin float32
	ClampMax float32

	DBKey string
}

func (o* Pid_t)Compute(currentTemperature float32) float32 {
	//Pid
	err := currentTemperature - o.TargetTemperature
	P := o.Kp * err
	//pId
	o.err_history = (o.err_history + err)/2
	I := o.Ki * o.err_history
	//piD
	D := o.Kd * (err - o.err_pre)

	val := o.u + P + I + D
	val = float32(math.Floor(float64(val + 0.5)))

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

