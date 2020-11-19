package skills

import (
	database "api/internal/pkg/db/mysql"
	"log"
)

type Skill struct {
	ID   string
	Name string
	Desc string
}

func (skill Skill) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Skills(Name,Desciption) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(skill.Name, skill.Desc)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("New skill inserted!")
	return id
}

func GetAll() []Skill {
	stmt, err := database.Db.Prepare("select id, Name, Desciption from Skills")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var skills []Skill
	for rows.Next() {
		var skill Skill
		err := rows.Scan(&skill.ID, &skill.Name, &skill.Desc)
		if err != nil {
			log.Fatal(err)
		}
		skills = append(skills, skill)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return skills
}

func GetListSkillByMemberId(memberId string) []Skill {
	stmt, err := database.Db.Prepare("select Skills.ID, Skills.Name, Skills.Desciption from Skills LEFT JOIN Member_Skill ON Member_Skill.skill_id=Skills.ID WHERE Member_Skill.member_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(memberId)
	var skills []Skill

	for rows.Next() {
		var skill Skill
		err := rows.Scan(&skill.ID, &skill.Name, &skill.Desc)
		if err != nil {
			log.Fatal(err)
		}
		skills = append(skills, skill)
	}

	return skills
}
