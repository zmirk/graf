package graf

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zmirk/zmlib"
)

type Level interface {
	GetMark() int
	GetAge() int
	GetName() string
	SetMark(mark int)
	SetAge(age int)
	SetName(name string)
	String() string
}

type Student struct {
	Age  int
	Name string
	Mark int
}

func (s *Student) GetMark() int {
	return s.Mark
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetMark(mark int) {
	s.Mark = mark
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %s\t|\tAge: %d\t|\tMark: %d", s.Name, s.Age, s.Mark)
}

func init() {
	println("*********** graf.go **********")
}

func Test() {
	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = zmlib.RandInt(0, 100)
	}

	fmt.Println(a)
}

func Test2() {
	a := zmlib.RandInt(0, 5)

	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	case 4:
		fmt.Println("a = 4")
	default:
		fmt.Println("a >= 5")
	}
}

func Test3() [10][10]int {
	var test [10][10]int

	for i := 0; i < 10; i++ {
		// test[i] = make(map[int]int)
		fmt.Printf("[")
		for j := 0; j < 10; j++ {
			test[i][j] = zmlib.RandInt(0, 99)

			if test[i][j] > 9 {
				fmt.Printf("%d ", test[i][j])
			} else {
				fmt.Printf(" %d ", test[i][j])
			}
		}
		fmt.Printf("]\n")
	}

	return test
}

func Test4() {
	matrix := Test3()

	for i, j := range matrix {
		fmt.Println(i, j)
	}

	fmt.Println("Test 4 end")
}

func Test5() [10]Student {
	students := [10]Student{}

	for i := 0; i < 10; i++ {
		students[i].SetMark(zmlib.RandInt(1, 5))
		students[i].SetName("Name" + fmt.Sprintf("%d", i))
		students[i].SetAge(zmlib.RandInt(18, 99))
	}

	return students
}

func Test6() {
	serialize, err := json.Marshal(Test5())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(serialize))
	}
}

func Test7() {
	fmt.Println(strings.Contains("Hello, World!", "World"))
}

func Test8() {
	go func() {
		zmlib.RandSleep(3000, 6000)
		fmt.Println(zmlib.RandInt(0, 100))
	}()
}
