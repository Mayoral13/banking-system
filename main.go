package main
import(
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
	"math/rand"
)

var counter int;
var maps = make(map[int]Customer);
type Customer struct{
	name string;
	balance int;
	accountNo int;
}
var customers []Customer;

func ShowDetail(){
	if(len(customers) < 1){
		fmt.Println("No Records")
	}else{
	for i:=0; i < len(customers); i++{
    fmt.Println("Customer's name is :" ,customers[i].name)
	fmt.Println("Customer's balance is :" ,customers[i].balance)
	fmt.Println("Customer's account Number is :" ,customers[i].accountNo)
	}
}
	
}
func Register(){
	fmt.Println("Enter name:")
	input:= bufio.NewReader(os.Stdin);
	text,_ := input.ReadString('\n');
	value := strings.TrimSpace(text);
	counter++;
	accountNo := rand.Int();
	balance := rand.Int()
	customers = append(customers,Customer{name:value,balance:balance,accountNo:accountNo})
	maps[accountNo] = Customer{name:value,balance:balance,accountNo:accountNo}
}
func Find(){
fmt.Println("Enter account Number")
input:= bufio.NewReader(os.Stdin);
text,_ := input.ReadString('\n');
value := strings.TrimSpace(text);
number,_ := strconv.Atoi(value);
if(maps[number].accountNo == number){
	fmt.Println("Record found")
	fmt.Println("Customer's name is :" ,maps[number].name)
	fmt.Println("Customer's balance is :" ,maps[number].balance)
	fmt.Println("Customer's account Number is :" ,maps[number].accountNo)
}else{
	fmt.Println("Record not found")
}
}
func Delete(){
	fmt.Println("Enter account Number")
	input:= bufio.NewReader(os.Stdin);
	text,_ := input.ReadString('\n');
	value := strings.TrimSpace(text);
	number,_ := strconv.Atoi(value);
	if(maps[number].accountNo == number){
		delete(maps,number)
		for data,item := range customers{
			if item.accountNo == number{
				customers = append(customers[:data],customers[data + 1:]... )
			}
		}
	}else{
		fmt.Println("Record not Found")

	}
	fmt.Println("Details deleted")
}
func Update(){
	fmt.Println("Enter account Number")
	input:= bufio.NewReader(os.Stdin);
	text,_ := input.ReadString('\n');
	value := strings.TrimSpace(text);
	number,_ := strconv.Atoi(value);
	if(maps[number].accountNo == number){
    fmt.Println("Enter name:")
	input:= bufio.NewReader(os.Stdin);
	text,_ := input.ReadString('\n');
	value := strings.TrimSpace(text);
	for data,item := range customers{
		if item.accountNo == number{
			customers = append(customers[:data],customers[data + 1:]... )
	      customers = append(customers,Customer{name:value,balance:maps[number].balance,accountNo:maps[number].accountNo})
	      maps[number] = Customer{name:value,balance:maps[number].balance,accountNo:maps[number].accountNo}
		}
	}
	
	}else{
		fmt.Println("Record not Found")

	}
}
func main(){
	fmt.Println("Welcome to the banking system")
	fmt.Println("Enter 1 to register user")
	fmt.Println("Enter 2 to find user")
	fmt.Println("Enter 3 to show all user details")
	fmt.Println("Enter 4 to delete user details")
	fmt.Println("Enter 5 to update user details")
	for true{
	input:= bufio.NewReader(os.Stdin);
	text,_ := input.ReadString('\n');
	value := strings.TrimSpace(text);
	if(value != "1" && value != "2" && value != "3" && value != "4" && value != "5"){
     fmt.Println("INVALID INPUT")
	 main()
	}else if(value == "1"){
		fmt.Println("Registration in progress")
		Register()
		fmt.Println("Registration successful")
	}else if(value == "2"){
		Find()
	}else if(value == "3"){
       ShowDetail()
	}else if(value == "4"){
		Delete()

	}else if(value == "5"){
		Update()
	}
}
}