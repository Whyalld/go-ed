package main

func main() {
	whiteList := []string{"HSsrx788", "TSlkj112"}
	processor := GetProcessorInstance(whiteList)

	types := []string{"humidity", "temperature", "pressure"}

	for _, t := range types {
		sensor, err := CreateSensor(t)
		if err != nil {
			continue
		}

		data := &SensorData{
			ID:    sensor.GetName(),
			Value: sensor.GetValue(),
		}

		processor.Process(data)
	}
}
