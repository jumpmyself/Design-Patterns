package main

import "fmt"

// 观察者接口
type Observer interface {
	Update(temperature float64, humidity float64, pressure float64)
}

// 被观察者接口
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

// 具体被观察者：气象站
type WeatherStation struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (w *WeatherStation) Attach(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) Detach(observer Observer) {
	for i, obs := range w.observers {
		if obs == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherStation) Notify() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherStation) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.Notify() // 通知所有观察者
}

// 具体观察者：手机显示
type PhoneDisplay struct{}

func (p *PhoneDisplay) Update(temperature float64, humidity float64, pressure float64) {
	fmt.Printf("手机显示：温度=%.2f°C, 湿度=%.2f%%, 气压=%.2fhPa\n", temperature, humidity, pressure)
}

// 具体观察者：电视显示
type TVDisplay struct{}

func (t *TVDisplay) Update(temperature float64, humidity float64, pressure float64) {
	fmt.Printf("电视显示：温度=%.2f°C, 湿度=%.2f%%, 气压=%.2fhPa\n", temperature, humidity, pressure)
}

// 客户端代码
func main() {
	// 创建被观察者（气象站）
	weatherStation := &WeatherStation{}

	// 创建观察者（显示设备）
	phoneDisplay := &PhoneDisplay{}
	tvDisplay := &TVDisplay{}

	// 注册观察者
	weatherStation.Attach(phoneDisplay)
	weatherStation.Attach(tvDisplay)

	// 更新气象数据，观察者会自动收到通知
	weatherStation.SetMeasurements(25.5, 65, 1013.25)
	fmt.Println("---")
	weatherStation.SetMeasurements(26.0, 70, 1012.50)
}
