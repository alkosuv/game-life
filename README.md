# Game life


Проект содержит cli реализацию игры "Game Life" на языке golang. 

Проект создавался в учебных целях и для ознакомления с правилами игра.

Ознакомится с правилами [Game Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

___

## Информация о реализации


*Ширина поля* - 120

*Высота поля* - 15

*Количество поколений* - 600

Если при генерации пиксел выходит за пределы поля, 
то он появляется на противоположной стороне поля.
___

## Сборка проекта

### build
Собранный проект будет находиться в директории *./bin/*

*windows os*

    go build -o .\bin\gamelife.exe .\cmd\gamelife\main.go

*mac os*

    go build -o .\bin\gamelife .\cmd\gamelife\main.go

### run
Приложение запустится в консоли

    go run .\cmd\gamelife\main.go