package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (d *Deps) InsertHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actorExists, err := d.ActorExists(r.Context(), event.Actor)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !actorExists {
		err = d.InsertEvent(r.Context(), event)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err = d.UpdateEvent(r.Context(), event)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (d *Deps) GetHandler(w http.ResponseWriter, r *http.Request) {
	actor := r.URL.Query().Get("actor")
	if actor == "" {
		// get all
		events, err := d.GetAllEvents(r.Context())
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(events)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	// get by actor
	events, err := d.GetEventsByActor(r.Context(), actor)
	if err != nil {
		if errors.Is(err, ErrNoRecordsFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
