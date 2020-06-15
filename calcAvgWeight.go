package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter 1st number: ")
    str1, _ := reader.ReadString('\n')

    // remove newline
    str1 = strings.Replace(str1,"\n","",-1)

    fmt.Print("Enter 2nd number: ")
    str2, _ := reader.ReadString('\n')

    // remove newline
    str2 = strings.Replace(str2,"\n","",-1)

    fmt.Print("Enter 3rd number: ")
    str3, _ := reader.ReadString('\n')

    // remove newline
    str3 = strings.Replace(str3,"\n","",-1)

    fmt.Print("Enter 4th number: ")
    str4, _ := reader.ReadString('\n')

    // remove newline
    str4 = strings.Replace(str4,"\n","",-1)

    fmt.Print("Enter 5th number: ")
    str5, _ := reader.ReadString('\n')

    // remove newline
    str5 = strings.Replace(str5,"\n","",-1)


    // convert string to int
    num1, e := strconv.Atoi(str1)
    if e != nil {
	fmt.Println("conversion error:", str1)
    }
    num2, e := strconv.Atoi(str2)
    if e != nil {
	fmt.Println("conversion error:", str2)
    }
    num3, e := strconv.Atoi(str3)
    if e != nil {
	fmt.Println("conversion error:", str3)
    }
    num4, e := strconv.Atoi(str4)
    if e != nil {
	fmt.Println("conversion error:", str4)
    }
    num5, e := strconv.Atoi(str5)
    if e != nil {
	fmt.Println("conversion error:", str5)
    }
        avg := (num1+num2+num3+num4+num5) / 5

        fmt.Println("Average:", avg)

    } else {
        fmt.Println("num(s) not in range")
    }
}
