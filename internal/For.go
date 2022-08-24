package internal

import (
	"fmt"
	"os"
	"syscall"
)

func GetCurrentThreadId() int {
	var user32 *syscall.DLL
	var GetCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		fmt.Printf("syscall.LoadDLL fail: %v\n", err.Error())
		return 0
	}
	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		fmt.Printf("user32.FindProc fail: %v\n", err.Error())
		return 0
	}

	var pid uintptr
	pid, _, err = GetCurrentThreadId.Call()

	return int(pid)
}

func ForTest() {
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	fmt.Printf("pid:%d    tid:%d \n", os.Getpid(), GetCurrentThreadId())

	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	//使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}
}
