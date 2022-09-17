package repository

import (
	"testing"

	"github.com/abraaolincoln/todoapp-api/database"
	"github.com/abraaolincoln/todoapp-api/domain"
	projectRepository "github.com/abraaolincoln/todoapp-api/repository/project"
)

func TestSaveProject(t *testing.T) {
	database.Connect()
	project := domain.Project{
		Id:         "1",
		Name:       "Project A",
		Color:      "white",
		CreateAt:   "today",
		ModifiedAt: "never",
	}
	projectResult := domain.Project{}

	projectRepository.Save(project)

	db := database.GetDatabase()

	resultSet, err := db.Query("SELECT * FROM project WHERE id = 1")

	if err != nil {
		t.Error("Expecting query not to fail")
	}

	for resultSet.Next() {
		err := resultSet.Scan(&projectResult.Id, &projectResult.Name, &projectResult.Color, &projectResult.CreateAt, &projectResult.ModifiedAt)

		if err != nil {
			t.Error("Expecting scan not to fail")
		}
	}

	if project.Id == projectResult.Id &&
		project.Name == projectResult.Name &&
		project.Color == projectResult.Color &&
		project.CreateAt == projectResult.CreateAt &&
		project.ModifiedAt == projectResult.ModifiedAt {
		t.Error("Expecting project create be equal to project got from database")
	}
}

func TestFindProjectByIdProject(t *testing.T) {
	database.Connect()
	project := domain.Project{
		Id:         "1",
		Name:       "Project A",
		Color:      "white",
		CreateAt:   "today",
		ModifiedAt: "never",
	}
	projectRepository.Save(project)

	projectResult, err := projectRepository.FindById(project.Id)

	if err != nil {
		t.Error("Expecting query not to fail")
	}

	if project.Id == projectResult.Id &&
		project.Name == projectResult.Name &&
		project.Color == projectResult.Color &&
		project.CreateAt == projectResult.CreateAt &&
		project.ModifiedAt == projectResult.ModifiedAt {
		t.Error("Expecting project create be equal to project got from database")
	}
}

// TOOD:
// test findAll
// test update
// test delete
