package repository

import (
	"github.com/abraaolincoln/todoapp-api/database"
	"github.com/abraaolincoln/todoapp-api/domain"
	log "github.com/abraaolincoln/todoapp-api/logger"
)

func FindAll() ([]domain.Project, error) {
	var projects []domain.Project
	db := database.GetDatabase()

	resultSet, err := db.Query("SELECT * FROM project")

	if err != nil {
		log.Error("Counld't do the query")
		return projects, err
	}

	for resultSet.Next() {
		var project domain.Project
		err := resultSet.Scan(&project.Id, &project.Name, &project.Color, &project.CreateAt, &project.ModifiedAt)

		if err != nil {
			log.Error("Couldn't scan result set to struct")
			return projects, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func FindById(id string) (domain.Project, error) {
	var project domain.Project
	db := database.GetDatabase()
	selectQuery := "SELECT * FROM project WHERE id = ?"

	resultSet, err := db.Query(selectQuery, id)

	if err != nil {
		log.Error("Couldn't do the query")
		return project, err
	}

	for resultSet.Next() {
		err := resultSet.Scan(&project.Id, &project.Name, &project.Color, &project.CreateAt, &project.ModifiedAt)

		if err != nil {
			log.Error("Couldn't scan result set to struct")
			return project, err
		}
	}

	return project, nil
}

func Save(project domain.Project) error {
	db := database.GetDatabase()
	insert, err := db.Prepare("INSERT INTO project (id, name, color, created_at, modified_at) values (?,?,?,?,?)")

	if err != nil {
		log.Error(err.Error())
		log.Error("Couldn't prepare insert")
		return err
	}

	_, err = insert.Exec(project.Id, project.Name, project.Color, project.CreateAt, project.ModifiedAt)

	if err != nil {
		log.Error("Couldn't insert project")
		return err
	}

	return nil
}

func DeleteById(projectId string) (bool, error) {
	db := database.GetDatabase()
	delete, err := db.Prepare("DELETE FROM project WHERE id = ?")

	if err != nil {
		log.Error("Couldn't prepare insert")
		return false, err
	}

	result, err := delete.Exec(projectId)

	if err != nil {
		log.Error("Couldn't delete project")
		return false, err
	}

	rowsAffect, _ := result.RowsAffected()
	if rowsAffect == 1 {
		return true, nil
	}

	return false, nil
}
