package main

import (
	"fmt"
	"net/http"
	"strconv"
)
type people struct {
	name string
	age int
}
func main() {
	users:=make(map[int]people)
	users[1]=people{name:"Marry",age: 18}
	users[2]=people{name:"Alexander",age: 20}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if(r.Method=="GET") {
			id, err :=strconv.Atoi(r.URL.Query().Get("id"))

			if(id> len(users)||id<1||err!=nil) {
				fmt.Fprintf(w,"Entered is not correct ")
			} else{
				fmt.Fprintf(w,"name= %s age= %d",users[id].name,users[id].age)
			}
		}else if (r.Method=="POST"){
			/*name,err:=ioutil.ReadAll(r.Body)
			if err!=nil{
				_,_=w.Write([]byte(err.Error()))
				return
			}
			_,_=fmt.Fprintf(w,"HI %+v",string(name))*/
			name := r.FormValue("name")
			age,err :=strconv.Atoi(r.FormValue("age"))

			isFind:=false
			for i:=1;i< len(users);i++ {
				if(name==users[i].name&& age==users[i].age){
					fmt.Fprintf(w, "User: %s with age: %s has id = %d", name, age,i)
					isFind=true
				}
			}
			if (!isFind||err!=nil){
				fmt.Fprintf(w, "There is no user with name: %s and age: %d .", name, age)
			}
		}

	})
	fmt.Println("Server is listening...")
	err:=http.ListenAndServe(":8080", nil)
	if err!=nil{
		fmt.Println(err.Error())
	}
}
