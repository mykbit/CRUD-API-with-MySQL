package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mykbit/CRUD-API-with-MySQL/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	_, err := app.vinyls.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Displaying all vinyls in the database..."))
}

func (app *application) vinylView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	_, err = app.vinyls.GetById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Write([]byte("Displaying a specific vinyl..."))
}

func (app *application) vinylCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying the form for adding a new vinyl..."))
}

func (app *application) vinylCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	description := r.PostForm.Get("description")
	releaseDate := r.PostForm.Get("releaseDate")

	id, err := app.vinyls.Insert(title, description, releaseDate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/vinyl/view/%d", id), http.StatusSeeOther)
}

func (app *application) vinylUpdate(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	_, err = app.vinyls.GetById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Write([]byte("Displaying the form for updating the vinyl..."))
}

func (app *application) vinylUpdatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PostForm.Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	title := r.PostForm.Get("title")
	description := r.PostForm.Get("description")
	releaseDate := r.PostForm.Get("releaseDate")

	err = app.vinyls.Update(id, title, description, releaseDate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/vinyl/view/%d", id), http.StatusSeeOther)
}

func (app *application) vinylDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	err = app.vinyls.Delete(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
