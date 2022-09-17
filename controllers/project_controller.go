package controllers

import (
	"fmt"
	"github.com/abraaolincoln/todoapp-api/domain"
	log "github.com/abraaolincoln/todoapp-api/logger"
	projectRepository "github.com/abraaolincoln/todoapp-api/repository/project"
	"github.com/abraaolincoln/todoapp-api/router"
	"github.com/abraaolincoln/todoapp-api/util"
	"net/http"
)

func RegisterProjectController(rm *router.RestMux) {
	rm.Post("/project", saveNewProject)
	rm.Get("/project/:id", findProjectById)
	rm.Get("/project", findAllProject)
}

func saveNewProject(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	var newProject domain.Project
	util.ParseBody(r, &newProject)

	err := validateProject(&newProject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = projectRepository.Save(newProject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func validateProject(project *domain.Project) error {
	fmt.Println(project.Id)
	fmt.Println(project.Name)
	fmt.Println(project.Color)
	fmt.Println(project.CreateAt)
	fmt.Println(project.ModifiedAt)

	return nil
}

func findProjectById(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	projectId := extraInfo.PathVariables[":id"]
	project, err := projectRepository.FindById(projectId)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if project.IsEmpty() {
		log.Info("Project with id " + projectId + " not found")
		w.WriteHeader(http.StatusNotFound)

		err = util.PutResultOnResponse(w, domain.Empty{})
		return
	} else {
		err = util.PutResultOnResponse(w, &project)
	}
}

func findAllProject(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	project, err := projectRepository.FindAll()

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = util.PutResultOnResponse(w, &project)
}
