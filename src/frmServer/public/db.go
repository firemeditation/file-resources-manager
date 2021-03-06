package public

import(
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"os"
)
	
func connDB () *sql.DB {
	db_server, _ := ServerConfig.GetString("db","server")
	db_port, _ := ServerConfig.GetString("db","port")
	db_user, _ := ServerConfig.GetString("db","user")
	db_passwd, _ := ServerConfig.GetString("db","passwd")
	db_dbname, _ := ServerConfig.GetString("db","dbname")
	connection_string := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", db_dbname, db_user, db_passwd, db_server, db_port)
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		fmt.Fprintf(os.Stderr, "链接数据库出错：", err)
		os.Exit(1)
	}
	return db
}
