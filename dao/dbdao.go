package dao

import "fmt"

func DBQL(msg string) string {
	var ret string
	rows, err := DB.Query(msg)
	if err != nil {
		fmt.Println(err)
		return ret
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return ret
	}
	for i := 0; i < len(cols); i++ {
		ret += cols[i] + "\t"
	}
	ret += "\n"
	for rows.Next() {
		var s, v string
		for i := 0; i < len(cols); i++ {
			rows.Scan(&v)
			s += v + ": "
			fmt.Printf("%v\n", v)
		}
		ret += s + "\n"
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
