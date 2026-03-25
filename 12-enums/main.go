package main

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w Weekday) String() string {
	return [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[w-1]
}

type TrafficLight int

const (
	Red TrafficLight = iota
	Yellow
	Green
)

func (t TrafficLight) String() string {
	return [...]string{"Red", "Yellow", "Green"}[t]
}

func (t TrafficLight) Next() TrafficLight {
	//the cycle is Red(0) → Green(2) → Yellow(1) → Red(0)
	switch t {
	case Red:
		return Green
	case Green:
		return Yellow
	case Yellow:
		return Red
	default:
		return Red // default case, should not happen
	}
}

func main() {
	for i := Monday; i <= Sunday; i++ {
		println(i.String())
	}

	light := Red
	for i := 0; i < 6; i++ {
		println(light.String())
		light = light.Next()
	}
}
