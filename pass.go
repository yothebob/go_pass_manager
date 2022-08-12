package main

import (
	"fmt"
	b64 "encoding/base64"
	"os"
	"log"
	"bufio"
	"strings"
)


func read_file() []string{
	var accounts []string
	file, err := os.Open(".data.txt")
	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		accounts = append(accounts,scanner.Text())
	}
	return accounts
}


func write_account(account_data string){
	file, err := os.OpenFile(".data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		log.Fatal(err)
	}
	datawriter := bufio.NewWriter(file)
	_,_ = datawriter.WriteString(account_data)
	_,_ = datawriter.WriteString("\n")
	datawriter.Flush()
	file.Close()
}


func supersecure_hashing(app_name string, account_name string, passwd string) string {
	hashed_str := app_name+ "." + b64.URLEncoding.EncodeToString([]byte(account_name)) + "." + b64.URLEncoding.EncodeToString([]byte(passwd))
	return hashed_str
}


func de_hashify(hashed_str string) []string {
	var res_str []string
	split_str := strings.Split(hashed_str,".")
	fmt.Println(split_str)
	for idx , val := range split_str {
		switch idx {
		case 1:
			decoded, _ := b64.URLEncoding.DecodeString(val)
			res_str = append(res_str, string(decoded))
		case 2:
			decoded, _ := b64.URLEncoding.DecodeString(val)
			res_str = append(res_str, string(decoded))
		}
	}
	return res_str
}


func new_account () {
	var app_name string
	var account_name string
	var passwd string
	fmt.Println("App Name:\n")
	fmt.Scanln(&app_name)
	fmt.Println("Account name:\n")
	fmt.Scanln(&account_name)
	fmt.Println("Password:\n")
	fmt.Scanln(&passwd)
	account_str := supersecure_hashing(app_name, account_name, passwd)
	write_account(account_str)
}


func get_account (accounts []string) {
	fmt.Println("what app are you looking for?\n")
	var app_name string
	fmt.Scanln(&app_name)
	
	for _,val := range accounts {
		if strings.Contains(val, app_name) {
			fmt.Println("I Found it!")
			
			name_pass := de_hashify(val)
			fmt.Println("Account Name: ",name_pass[0])
			fmt.Println("Account Password: ", name_pass[1])
			fmt.Println("\n\n")
		}
	}
}


func password_manage(global_password string, accounts []string) {
	var passwd string
	fmt.Println(":")
	fmt.Scanln(&passwd)
	if passwd == global_password {
		fmt.Println("WARNING: do not put real creds in here, and if you want to you NEED to update the password 'hash' function")
		for {
			fmt.Println("type 'get' to get password \ntype 'new' to insert new account\ntype 'end' to leave")
			var cmd string
			fmt.Scanln(&cmd)
			switch cmd {
			case "new":
				new_account()
			case "get":
				get_account(accounts)
			case "end":
				os.Exit(3)
			}
		}
	}
}


func main(){
	global_password := "pwd"
	accounts := read_file()
	password_manage(global_password, accounts)	

}
