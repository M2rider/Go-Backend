package main
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Name string `json:"name"`
}

func main(){
	fmt.Println("GO MySQL Beginning")

	db,err := sql.Open("mysql","root:GetSQL@tcp(127.0.0.1:3306)/testdb")

	if err!= nil{
		panic(err.Error())
	}
	
	defer db.Close()
	fmt.Println("Successfully started the database")
	/*
	// inserting
	insert,err := db.Query("INSERT INTO users VALUES('MAYANK')")
	if err!= nil{
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted into user tables")
	//inserting
	insertt,err := db.Query("INSERT INTO users VALUES('ASHISH')")
	if err!= nil{
		panic(err.Error())
	}
	defer insertt.Close()
	fmt.Println("Successfully inserted into user tables")
	*/

	/*
	// Truncate
	delete,err := db.Query("TRUNCATE users")
	if err!=nil {
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Println("Successfully Truncated users table")
	*/

	results,err := db.Query("SELECT name FROM users")
	if err!=nil{
		panic(err.Error())
	}

	for results.Next(){
		var user User
		err = results.Scan(&user.Name)
		if err!=nil{
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

}