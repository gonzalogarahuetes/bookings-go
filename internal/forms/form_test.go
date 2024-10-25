package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid when should have been valid.")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields are missing.")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows invalid when required fields are present.")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has2 := form.Has("e")
	if has2 {
		t.Error("Says it has a field when it does not have it.")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has := form.Has("a")
	if !has {
		t.Error("Says it has not a field when it has it.")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	var minLength = 3

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("It shows min length for non-existent field.")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have an error but did not get one.")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "aaaaaa")
	form = New(postedData)

	hasMinLength := form.MinLength("a", minLength)
	if hasMinLength {
		t.Error("It approves the minimum length when it should not.")
	}

	hasMinLength2 := form.MinLength("b", minLength)
	if !hasMinLength2 {
		t.Error("It rejects the minimum length when it should not.")
	}

	isError = form.Errors.Get("b")
	if isError != "" {
		t.Error("Should not have an error but got one.")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("It takes as a valid email a non-existent field.")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("It takes as a valid email something it is not.")
	}

	postedData = url.Values{}
	postedData.Add("b", "realemail@yes.com")
	form = New(postedData)

	form.IsEmail("b")
	if !form.Valid() {
		t.Error("It takes as a wrong email something it is not.")
	}
}
