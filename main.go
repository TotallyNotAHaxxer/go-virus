package main

import (
    "fmt"
	//"os"
    "log"
    "net/http"
	"os/exec"
	"runtime"
	"io/ioutil"
	"flag"
)


func clear() {
	out, err := exec.Command("clear").Output()
		if err != nil {
			fmt.Println("%s", err)
		}
		fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)
}


func viurs() {
	urlvirus := "https://bit.ly/3ild93L"

	fmt.Println(urlvirus)
}


func Command() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[+] Current Dir", output)
}


func check() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	if runtime.GOOS == "linux" {
		fmt.Println("[+] Looks Good.....")
		fmt.Println("[+] Testing.....")
	}
	} else {
		fmt.Println("[+] Looks Good.....")
		fmt.Println("[+] Testing.....")
		out, err := exec.Command("clear").Output()
		if err != nil {
			fmt.Println("%s", err)
		}
		fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)
	}
}

// use of color go get github.com/fatih/color

func urltest() {
	url1 := "https://youtube.com"
	url2 := "https://google.com"
	url3 := "https://hulu.com"
	url4 := "https://netflix.com"
	url5 := "https://www.twitter.com/"
	url6 := "https://www.facebook.com"
	url7 := "https://instagram.com"
	url8 := "https://keybase.io"
	url9 := "https://doxbin.org"

	fmt.Println("-------------- Testing URL's ------------------")
	fmt.Println("[+] -> ", url1)
	fmt.Println("[+] -> ", url2)
	fmt.Println("[+] -> ", url3)
	fmt.Println("[+] -> ", url4)
	fmt.Println("[+] -> ", url5)
	fmt.Println("[+] -> ", url6)
	fmt.Println("[+] -> ", url7)
	fmt.Println("[+] -> ", url8)
	fmt.Println("[+] -> ", url9)
	fmt.Println("------------------------------------------------")
	resp0, err := http.Get(url1)
	resp1, err := http.Get(url2)
	resp2, err := http.Get(url3)
	resp3, err := http.Get(url4)
	resp4, err := http.Get(url5)
	resp5, err := http.Get(url6)
	resp9, err := http.Get(url7)
	resp6, err := http.Get(url8)
	resp7, err := http.Get(url9)

	if err != nil {
		log.Fatal(err)
	}
	if resp7.StatusCode >= 200 && resp7.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp6.StatusCode >= 200 && resp6.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp5.StatusCode >= 200 && resp5.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp4.StatusCode >= 200 && resp4.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp3.StatusCode >= 200 && resp3.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp2.StatusCode >= 200 && resp2.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp1.StatusCode >= 200 && resp1.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp9.StatusCode >= 200 && resp9.StatusCode <= 299 {
		fmt.Println("OK")
	}
	if resp0.StatusCode >= 200 && resp0.StatusCode <= 299 {
		fmt.Println("OK")
	} else {
		fmt.Println("OK")
	}
}

func secondtest() {
	resp, err := http.Get("http://golangcode.com")
    if err != nil {
        log.Fatal(err)
    }

    // Print the HTTP Status Code and Status Name
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("Returned Error => ", resp.StatusCode)
	}
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Returned in range of 200 ")
    } else {
        fmt.Println("Argh! Broken")
    }
}

// why not make a nicer bannr than the other one had  
func banner() string {
    b, err := ioutil.ReadFile("ascii.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
	return (string(b))
}


func userin() {
	var name string;
 	fmt.Printf("Write Device >>> ")
  	fmt.Scanf("%s", &name)
//	name :=

}


func cmdarg() {
	output := flag.Bool("output", false, "Android")
	input := flag.Bool("input", true, "Iphone") 
	flag.Parse()
	fmt.Println("Android => ",*output)
	fmt.Println("Iphone  => ",*input)


}

func payloads() {

	fmt.Println("-------------- IPHONE SEND PAYLOADS --------------------")
	fmt.Println("  effective. ")
	fmt.Println("      Power  ")
	fmt.Println("	         لُلُصّبُلُلصّبُررً ॣ ॣh ॣ ॣ")
	fmt.Println("冗")
	fmt.Println("------------------ ANDROID MANUAL PAYLOADS ------------")
	fmt.Println("https://bit.ly/3ild93L")
	fmt.Println("https://bit.ly/3fX8ljZ")
	fmt.Println("-----------------------")
	fmt.Println(" sorry there is not more (◕‿◕)♡ have a great day! ")

}

func dir() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		out, err := exec.Command("pwd").Output()
		if err != nil {
			fmt.Println("%s", err)
		}
		//fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println("                    Working Dir : ",output)
	}
}


func main() {
	clear()
	urltest()
	check()
	Command()
	secondtest()
	clear()
	banner()
	dir()
	//cmdarg()
	//userin()
	payloads()

}
