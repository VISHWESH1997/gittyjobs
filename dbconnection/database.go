package dbconnection
import(
	"log"
	_"mysql"
	_"mysqlserver"
	"database/sql"
)
var db *sql.DB
func Connection()*sql.DB{

	Db,err := sql.Open("mysql","root:root@tcp(localhost:3306)/gitty")
	
	if err!=nil{
		log.Fatalln(err)
	}
	return Db
}