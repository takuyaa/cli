package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// increment bouryoku
	cmd := exec.Command("curl", strings.Split("-X PUT https://pixe.la/v1/users/t-kyt/graphs/takuya-violence/increment -H X-USER-TOKEN:yu-yu-yu", " ")...)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)
}
