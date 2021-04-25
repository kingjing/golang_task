import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// DAO层

func foo() (id int, err error) {
	db, err := sql.Open("mysql", "root:user/test") //后面格式为"user:password@/dbname"
	defer db.Close()
	rows, err := db.QueryRow("SELECT id from aaaaa where id=123")
	defer rows.Close()
	return rows.Scan(&id), errors.Wrap(sql.ErrNoRows, "foo() sql: data not found")
}

// 调用层
func main() {
	_, err := foo()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
