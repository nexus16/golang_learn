package members

import (
	database "api/internal/pkg/db/mysql"
	"log"
)

type Member struct {
	ID      string
	Name    string
	Age     string
	SkillID []int64
}

func (member Member) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Members(Name,Age) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(member.Name, member.Age)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("New member inserted!")
	return id
}

func GetAll() []Member {
	stmt, err := database.Db.Prepare("select ID, Name, Age from Members")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Age)
		if err != nil {
			log.Fatal(err)
		}
		members = append(members, member)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return members
}
