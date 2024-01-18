package weatherdata

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	notifyObservers()
}

type WeatherData struct {
	observers map[string]Observer
	data      map[string]string
}

func NewWeatherData() *WeatherData {
	return &WeatherData{observers: map[string]Observer{}}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers[o.Name()] = o
}

func (w *WeatherData) RemoveObserver(o Observer) {
	delete(w.observers, o.Name())
}

func (w *WeatherData) notifyObservers() {
	for _, o := range w.observers {
		o.Update(w.data)
	}
}

func (w *WeatherData) SetWeatherData(data map[string]string) {
	w.data = data
	w.notifyObservers()
}
