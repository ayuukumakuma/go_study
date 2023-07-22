package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=todo_app dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// psqlはテーブル作成をsqlファイルで行う
	// `psql -U todo_app -f example_psql.sql -d testdb`

	// Create
	// psqlは$1, $2のようなプレースホルダを使う
	// cmd := `INSERT INTO persons (name, age) VALUES ($1, $2)`
	// _, err = Db.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Read
	// 単体
	// cmd := `SELECT * FROM persons where age = $1`
	// row := Db.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age) // row.Scanは引数にポインタを取る
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// 複数
	// cmd := "SELECT * FROM persons"
	// rows, _ := Db.Query(cmd) // cmdに一致するデータを全て取得
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// Update
	// cmd := "UPDATE persons SET age = $1 WHERE name = $2"
	// _, err = Db.Exec(cmd, 25, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Delete
	// cmd := "DELETE FROM persons WHERE name = $1"
	// _, err = Db.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
