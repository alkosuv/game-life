package univers

// Colony тип представляющей состояния поля
type Colony [][]bool

// NewColony инициализация новой колонии
func NewColony(park Park) Colony {
	colony := make(Colony, park.Hight)
	for i := range colony {
		colony[i] = make([]bool, park.Width)
	}

	return colony
}

// Alive получения статуса клетки
// Функция возвращает состояние клетки
func (c Colony) Alive(x int, y int, park *Park) bool {
	w := (x + park.Width) % park.Width
	h := (y + park.Hight) % park.Hight
	return c[h][w]
}

// Neighbors подсчет количества соседей
// если значение соседений клетки true, то count увеличивается на единицу
func (c *Colony) Neighbors(x int, y int, park *Park) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if c.Alive(x+j, y+i, park) == true {
				count++
			}
		}
	}

	return count
}
