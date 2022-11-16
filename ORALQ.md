#### Запуск сервиса
```shell
Какой самый эффективный способ конкатенации строк?
```
```shell
Что такое интерфейсы, как они применяются в Go?

```
```shell
Чем отличаются RWMutex от Mutex?
```
```shell
Чем отличаются буферизированные и не буферизированные каналы?
```
```shell
Какой размер у структуры struct{}{}?
```
```shell
Есть ли в Go перегрузка методов или операторов?
```
```shell
В какой последовательности будут выведены элементы map[int]int?

Пример:
m[0]=1
m[1]=124
m[2]=281

```
```shell
В чем разница make и new?
```
```shell
Сколько существует способов задать переменную типа slice или map?
```
```shell
Что выведет данная программа и почему?


func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}

```
```shell
Что выведет данная программа и почему?

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}

```
```shell
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}

```
```shell
Что выведет данная программа и почему?

func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}

```
```shell
Что выведет данная программа и почему?


func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

```
