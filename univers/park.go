package univers

/*
Park структура хранит размеры игрового поля

	Width - ширина поля
	Hight - высота поля
*/
type Park struct {
	Width int
	Hight int
}

// NewPark создаёт новое поле
func NewPark(width int, hight int) *Park {
	return &Park{
		Width: width,
		Hight: hight,
	}
}
