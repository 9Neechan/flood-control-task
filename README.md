### Рассуждения в ходе решения

Загвоздка решения состояла в том, как хранить данные. У меня было несколько идей:
- хранить пары, где ключом является UserID, а занчением - Time Stamp запроса. Это решение не подходит при использовании Redis, тк он не моддеживает храниение записей с одинаковыми ключами (в отличие от MongoDB. Но решила не брать Монгу тк она намного медленнее Редиса)
- хранить пары, где ключом является UserID, а занчением - список Time Stamp запросов (этот вариант я выбрала для реализации)

Решение реализовано с использованием RedisDB. Была выбрана эта БД, тк она быстрая и решение предполагало хранение данных в виде пар ключ-значение.

В ходе проделанной работы был реализован метод Check(). Суть метода: 
- добавляем в список Time Stamp запроса по UserID
- по UserID достаем весь список запросов
- пробегаемся по запросам и ищем первое вхождение Time Stamp, который входит в проверяемый промежуток времени
- удаляем из массива все элементы, стоящие до него и перезаписываем в БД
- длину оставшегося списка проверяем на условие флуд-контроля

В итоге метод работает за O(n), где n - кол-во вызова метода Check().


### Запуск

```go run main.go```


