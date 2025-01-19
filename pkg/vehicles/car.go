package vehicles

type Car struct {
	Brand string
	Model string
}

func NewCar(brand, model string) *Car {
	return &Car{
		Brand: brand,
		Model: model,
	}
}
