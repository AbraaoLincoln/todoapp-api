package controllers

import (
	"github.com/abraaolincoln/todoapp-api/domain"
	log "github.com/abraaolincoln/todoapp-api/logger"
	projectRepository "github.com/abraaolincoln/todoapp-api/repository/project"
	"github.com/abraaolincoln/todoapp-api/router"
	"github.com/abraaolincoln/todoapp-api/util"
	"net/http"
)

func RegisterProjectController(rm *router.RestMux) {
	rm.Post("/projects", saveNewProject)
	rm.Get("/projects/:id", findProjectById)
	rm.Get("/projects/", findAllProject)
	rm.Delete("/projects/:id", deleteProjectById)
}

const ContentType = "Content-type"
const ApplicationJon = "application/json"

func saveNewProject(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	var newProject domain.Project
	util.ParseBody(r, &newProject)

	err := validateProject(&newProject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = projectRepository.Save(newProject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateProject(project *domain.Project) error {
	//fmt.Println(project.Id)
	//fmt.Println(project.Name)
	//fmt.Println(project.Color)
	//fmt.Println(project.CreateAt)
	//fmt.Println(project.ModifiedAt)

	return nil
}

func findProjectById(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	w.Header().Set(ContentType, ApplicationJon)
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
	w.Header().Set(ContentType, ApplicationJon)
	project, err := projectRepository.FindAll()

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = util.PutResultOnResponse(w, &project)
}

func deleteProjectById(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	w.Header().Set(ContentType, ApplicationJon)
	projectId := extraInfo.PathVariables[":id"]
	projectWasDeleted, err := projectRepository.DeleteById(projectId)

	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if projectWasDeleted {
		log.Info("Project with id " + projectId + " was deleted")
		w.WriteHeader(http.StatusNoContent)
	} else {
		log.Info("Unable to delete project with id " + projectId)
		w.WriteHeader(http.StatusBadRequest)
	}
}
