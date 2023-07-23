package dao

import "fmt"

func DBQL(msg string) string {
	var ret string = ""
	rows, err := DB.Query(msg)
	if err != nil {
		fmt.Println(err)
		return ret
	}
	defer rows.Close()
	cols, err := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	if err != nil {
		fmt.Println(err)
		return ret
	}
	for i := 0; i < len(cols); i++ {
		ret += cols[i] + "\t"
	}
	ret += "\n"
	i := 0
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
		}
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			if key != "_row" {
				ret += string(v) + "\t"
			}
			i++
		}
		ret += "\n"
	}
	return ret
}
func DBML(msg string) bool {
	stmt, err := DB.Prepare(msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}
	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(id)
	return true
}
func DBDL(msg string) bool {
	stmt, err := DB.Prepare(msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}
	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(id)
	return true
}
