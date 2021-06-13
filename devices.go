package remo

import "time"

type Device struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	TemperatureOffset float32   `json:"temperature_offset"`
	HumidityOffset    float32   `json:"humidity_offset"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	FirmwareVersion   string    `json:"firmware_version"`
	MacAddress        string    `json:"mac_address"`
	SerialNumber      string    `json:"serial_number"`
	NewestEvents      Events    `json:"newest_events"`
}

type Events struct {
	Temperature  SensorValue `json:"te"`
	Humidity     SensorValue `json:"hu"`
	Illumination SensorValue `json:"il"`
	Movement     SensorValue `json:"mo"`
}

type SensorValue struct {
	Value     float32   `json:"val"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Client) GetDevices() ([]Device, error) {
	result := []Device{}
	if err := c.getApi("/1/devices", &result); err != nil {
		return []Device{}, err
	}

	return result, nil
}
