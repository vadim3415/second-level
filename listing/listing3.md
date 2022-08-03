
Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main
import (
	"fmt"
	"os"
)
func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
    fmt.Printf("%v, %T", err, err) // <nil>, *fs.PathError
}
```

Ответ:
```
Программа выведет nil false. 
Интерфейс равен nil только в том случае, когда и тип, и значение равны nil.
```
**Пустой интерфейс**

Структура данных для пустого интерфейса — это `iface` без `itab`. Тому есть две причины:

- Поскольку в пустом интерфейсе нет методов, все, что связано с динамической диспетчеризацией, можно смело выкинуть из структуры данных.
- Когда виртуальная таблица исчезла, тип самого пустого интерфейса, не путать с типом данных, которые он содержит, всегда один и тот же.

`eface` — это корневой тип, представляющий пустой интерфейс в рантайме 
Его определение выглядит так:

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

Где `_type` содержит информацию о типе значения, на которое указывают данные.
`itab` был полностью удален.
```go
type iface struct { // `iface`
    tab *struct { // `itab`
        inter *struct { // `interfacetype`
            typ struct { // `_type`
                size       uintptr
                ptrdata    uintptr
                hash       uint32
                tflag      tflag
                align      uint8
                fieldalign uint8
                kind       uint8
                alg        *typeAlg
                gcdata     *byte
                str        nameOff
                ptrToThis  typeOff
            }
            pkgpath name
            mhdr    []struct { // `imethod`
                name nameOff
                ityp typeOff
            }
        }
        _type *struct { // `_type`
            size       uintptr
            ptrdata    uintptr
            hash       uint32
            tflag      tflag
            align      uint8
            fieldalign uint8
            kind       uint8
            alg        *typeAlg
            gcdata     *byte
            str        nameOff
            ptrToThis  typeOff
        }
        hash uint32
        _    [4]byte
        fun  [1]uintptr
    }
    data unsafe.Pointer
}