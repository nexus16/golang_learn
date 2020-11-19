package memberSkill

import (
	database "api/internal/pkg/db/mysql"
	"log"
)

type MemberSkill struct {
	ID       int64
	MemberID string
	SkillID  string
}

func (ms MemberSkill) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Member_Skill(member_id,skill_id) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(ms.MemberID, ms.SkillID)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("New member skill inserted!")
	return id
}
