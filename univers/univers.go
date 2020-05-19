package univers

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

/*
Univers структура мира

	Park хранит размеры игрового поля
	Colony текущее состояние мира
	ColonyComput состояние колонии на новом шаге
*/
type Univers struct {
	Park         *Park
	Colony       Colony
	ColonyComput Colony
}

// NewUnivers создаёт игровое поле разсером 120x15
func NewUnivers() *Univers {
	park := NewPark(120, 15)
	u := &Univers{
		Park:         park,
		Colony:       NewColony(*park),
		ColonyComput: NewColony(*park),
	}
	u.seed()
	return u
}

// Show вывод колонии
func (u *Univers) Show() {
	u.clear()
	for i := range u.Colony {
		for j := range u.Colony[i] {
			switch u.Colony[i][j] {
			case true:
				fmt.Print("*")
			case false:
				fmt.Print(" ")
			}
		}
	}
}

// Next состояние клетки в новом поколении
func (u *Univers) Next(x int, y int) bool {
	n := u.Colony.Neighbors(x, y, u.Park)
	return n == 3 || n == 2 && u.Colony.Alive(x, y, u.Park)
}

// Computing вычисляет состояние колонии в новом поколение
func (u *Univers) Computing() {
	for i := range u.Colony {
		for j := range u.Colony[i] {
			u.ColonyComput[i][j] = u.Next(j, i)
		}
	}

	u.swop()
}

// seed генерация поколений 25% от общего числа клеток
func (u *Univers) seed() {
	limit := int(float64(u.Park.Hight) * float64(u.Park.Width) * 0.25)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < limit; i++ {
		w := rand.Intn(u.Park.Width)
		h := rand.Intn(u.Park.Hight)
		u.Colony[h][w] = true
	}
}

func (u *Univers) swop() {
	u.Colony, u.ColonyComput = u.ColonyComput, u.Colony
}

// clear очищает консоль
func (u *Univers) clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
