Реализуйте функцию, которая принимает на вход слайс каналов и параллельно ожидает сообщений из всех этих каналов. Кол-во каналов в слайсе каждый раз отличается. Известно только, что их всегда 1 и более.

Чтобы это не казалось абстрактной задачей, представьте, что вы пишете websocket сервер, где каждый канал представляет собой TCP соединение с клиентом (конечно же TCP соединение реализовано где-то в другом месте и вы работаете уже только с данными в канале) и нужна функция, которая при получении сообщения от любого из клиентов в заданном слайсе может выполнить некоторую обработку.  А с каждым клиентом как раз и связан слайс каналов (т.е. список активных соединений этого клиента).

Реализовывать websocket и работу с TCP, конечно же не нужно.

Заготовка задачи

```go
package homework

func selectMany(channels []chan int64) chan int64 {
	// ваш код
}
```