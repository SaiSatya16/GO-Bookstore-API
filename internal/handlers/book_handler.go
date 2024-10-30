package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "bookstore-api/internal/models"
    "bookstore-api/internal/repository"
    "bookstore-api/pkg/response"
)

type BookHandler struct {
    repo *repository.BookRepository
}

func NewBookHandler(repo *repository.BookRepository) *BookHandler {
    return &BookHandler{repo: repo}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    if err := book.Validate(); err != nil {
        response.Error(w, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.repo.Create(&book); err != nil {
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    response.JSON(w, http.StatusCreated, book)
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    books, err := h.repo.GetAll()
    if err != nil {
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    response.JSON(w, http.StatusOK, books)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    book, err := h.repo.GetByID(params["id"])
    if err != nil {
        response.Error(w, http.StatusNotFound, "Book not found")
        return
    }

    response.JSON(w, http.StatusOK, book)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    book.ID = params["id"]
    if err := book.Validate(); err != nil {
        response.Error(w, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.repo.Update(&book); err != nil {
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    response.JSON(w, http.StatusOK, book)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    if err := h.repo.Delete(params["id"]); err != nil {
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    response.JSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}