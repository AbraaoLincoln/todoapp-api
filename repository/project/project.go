package repository

import (
	"github.com/abraaolincoln/todoapp-api/database"
	"github.com/abraaolincoln/todoapp-api/domain"
	log "github.com/abraaolincoln/todoapp-api/logger"
)

func FindById(id string) (domain.Project, error) {
	var project domain.Project
	db := database.GetDatabase()
	selectQuery := "SELECT * FROM project WHERE id = ?"

	resultSet, err := db.Query(selectQuery, id)

	if err != nil {
		log.Error("Counld't do the query")
		return project, err
	}

	for resultSet.Next() {
		err := resultSet.Scan(&project.Id, &project.Name, &project.Color, &project.CreateAt, &project.ModifiedAt)

		if err != nil {
			log.Error("Couldn't scan result set to sctruc")
		}
	}

	return project, nil
}

func Save(project domain.Project) error {
	db := database.GetDatabase()
	insert, err := db.Prepare("INSERT INTO project (id, name, color, create_at, modified_at) value (?,?,?,?,?)")

	if err != nil {
		log.Error("Counld't prepare insert")
		return err
	}

	_, err = insert.Exec(project.Id, project.Name, project.Color, project.CreateAt, project.ModifiedAt)

	if err != nil {
		log.Error("Counld't insert project")
		return err
	}

	return nil
}

func DeleteById(projectId string) error {
	db := database.GetDatabase()
	delete, err := db.Prepare("DELETE FROM project WHERE id = ?")

	if err != nil {
		log.Error("Counld't prepare insert")
		return err
	}

	_, err = delete.Exec(projectId)

	if err != nil {
		log.Error("Counld't delete project")
		return err
	}

	return nil
}
