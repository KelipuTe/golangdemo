package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pdb, err := sql.Open("mysql", makeDBConn()) //创建数据库操作句柄
	checkErr(err)
	defer pdb.Close() //关闭操作句柄，提前写
	Query(pdb)
	Insert(pdb)
	Query(pdb)
	Update(pdb)
	Query(pdb)
	Delete(pdb)
	Query(pdb)
}

//构造连接
func makeDBConn() string {
	var host, port, user, pass, name, charset string
	host = "rm-m5e8ei6f32ky6oxwwho.mysql.rds.aliyuncs.com"
	port = "3306"
	user = "testotaplus"
	pass = "Miot21t**test"
	name = "testotaplus"
	charset = "utf8"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pass, host, port, name, charset)
}

//处理error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//查询，直接执行sql
func Query(pdb *sql.DB) {
	rows, err := pdb.Query("SELECT * FROM test")
	checkErr(err)

	//rows.Close()是幂等的，如果rows.Next()返回false，rows也会自动关闭
	defer rows.Close() //关闭rows，提前写

	//rows.Next()迭代查询数据，获取不到数据时返回false，配合rows.Scan()获取数据
	for rows.Next() {
		var id int
		var testInt int
		var testVarchar string
		var testFloat float64
		var testText string
		var testDate string
		var testDatetime string
		var testTimpstamp string

		//rows.Scan()获取一行数据，这里的参数顺序和sql语句里的顺序要一致，要不然会错位
		err = rows.Scan(&id, &testInt, &testVarchar, &testFloat, &testText, &testDate, &testDatetime, &testTimpstamp)
		checkErr(err)

		fmt.Printf("id=%d,test_int=%d,test_varchar=%s,test_float=%f,test_text=%s,test_date=%s,test_datetime=%s,test_timpstamp=%s\r\n",
			id, testInt, testVarchar, testFloat, testText, testDate, testDatetime, testTimpstamp)
	}
}

//插入，预编译sql
func Insert(pdb *sql.DB) {
	stmt, _ := pdb.Prepare(`INSERT INTO test (test_int,test_varchar,test_float,test_text,test_date,test_datetime,test_timpstamp) VALUES (?,?,?,?,?,?,?)`)
	defer stmt.Close()

	result, err := stmt.Exec(2, "bbb", 2.22, "bbb", "2020-04-02", "2020-04-02 12:00:00", "2020-04-02 12:00:00")
	checkErr(err)

	id, err := result.LastInsertId()
	checkErr(err)

	fmt.Printf("insert id=%v\r\n", id)
}

//修改
func Update(pdb *sql.DB) {
	stmt, _ := pdb.Prepare(`UPDATE test SET test_int=?,test_float=? WHERE test_varchar=?`)
	defer stmt.Close()

	result, err := stmt.Exec(3, 3.33, "bbb")
	checkErr(err)

	nums, err := result.RowsAffected()
	checkErr(err)

	fmt.Printf("rows affected=%v\r\n", nums)
}

//删除
func Delete(pdb *sql.DB) {
	stmt, _ := pdb.Prepare(`DELETE FROM test WHERE id>?`)
	defer stmt.Close()

	result, err := stmt.Exec(1)
	checkErr(err)

	nums, err := result.RowsAffected()
	checkErr(err)

	fmt.Printf("rows affected=%v\r\n", nums)
}
