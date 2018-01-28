package mysql

import (
	"database/sql"
	_ "mysql-master"
	"fmt"
	"../tools"
)

var db *sql.DB
var err error

func InitInfo() {
	db,err = sql.Open("mysql", "root:@/blog?charset=utf8")
	if err != nil  {
		fmt.Println("Connet mysql faild.....")
	}


	db.SetMaxOpenConns(200)   //设置打开数据库的最大连接数  如果请求达到最大值 还有请求到来 即block(阻塞) 等待中
	db.SetMaxIdleConns(100)  //设置连接池中的保持连接的最大连接数
	err = db.Ping()  // 调用即刻把链接返回到链接池

	if err != nil {
		tools.LogInfo(err.Error())
	}
}

func FetchRow(sql string, args...interface{})(*map[string]string, error)  {
	stmtOut, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()  //读出查询出的列字段名
	if err != nil {
		panic(err.Error())
	}

	values := make([][]byte, len(columns))  // values是每个列的值，这里获取到byte里
	scanArgs := make([]interface{}, len(values)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next(){
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(values)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break
	}
	return &ret, nil
}

func UpdateRow(sql string, args...interface{})(int64, error)  {
	stmtIns, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		//panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result.RowsAffected()
}

func InsertRow(sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.LastInsertId()
}

func FetchRows(sql string, args...interface{}) (map[int]map[string]string, error)  {
	results := make(map[int]map[string]string)
	stmtOut, err := db.Prepare(sql)
	if err != nil {
		tools.LogError(err.Error())
		return results, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		tools.LogError(err.Error())
		return results, err
	}

	column, _ := rows.Columns()
	values := make([][]byte, len(column))
	scans  := make([]interface{}, len(column))

	for i := range values {
		scans[i] = &values[i]
	}

	i := 0

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			scans[i] = &values[i]
			tools.LogError(err.Error())
			return results, err
		}
		row := make(map[string]string) //每行数据

		for k, v := range values{  // 每行数据都在values里面，现把它放到row
			key := column[k]
			row[key] = string(v)
		}

		results[i] = row  //装入结果集中
		i++
	}
	tools.LogInfo("hello")
	return results,nil
}

func DbClose()  {
	defer db.Close()
}


