package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/repositories"
	"github.com/gorilla/mux"
)

type ProblemService struct {
	problemRepo *repositories.ProblemRepository
}

func NewProblemService(problemRepo *repositories.ProblemRepository) *ProblemService {
	return &ProblemService{
		problemRepo: problemRepo,
	}
}

func (p *ProblemService) CreateProblem(w http.ResponseWriter, r *http.Request) {
	newProblem := models.Problem{}
	if err := json.NewDecoder(r.Body).Decode(&newProblem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.problemRepo.CreateProblem(&newProblem); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *ProblemService) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	var filter repositories.ProblemFilter

	w.Header().Set("content-type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		log.Println("Failed while validating problem filter:" + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	problems, err := p.problemRepo.GetAllProblems(filter)
	if err != nil {
		log.Fatal("getting all problems by filter failed: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(problems); err != nil {
		log.Println("Failed while transferring date into response: " + err.Error())
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
}

func (p *ProblemService) GetProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemId := vars["id"]

	w.Header().Set("content-type", "application/json")

	problem, err := p.problemRepo.GetProblem(problemId)
	if err != nil {
		log.Println("Getting problem from database failed:" + err.Error())
		http.Error(w, "Problem not found"+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&problem); err != nil {
		log.Println("Writing to response failed:" + err.Error())
		http.Error(w, "Writing to response body failed:"+err.Error(), http.StatusBadRequest)
	}
}

func (p *ProblemService) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateFilter repositories.UpdateProblem

	if err := json.NewDecoder(r.Body).Decode(&updateFilter); err != nil {
		http.Error(w, "decoding to update filter failed: "+err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	if err := p.problemRepo.UpdateProblem(id, updateFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Failed updating problem: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Problem successfully updated"))
}

func (p *ProblemService) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := p.problemRepo.DeleteProblem(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
}
