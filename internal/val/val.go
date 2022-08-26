package val

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func retrieveIDFromURL(urlStr string, model string) (int64, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return 0, err
	}

	var id string
	urlSlice := strings.Split(url.Path, "/")
	for i, v := range urlSlice {
		if v == model {
			if len(urlSlice) > i+1 {
				id = urlSlice[i+1]
			}
		}
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return int64(i), nil
}

func formToMap(r *http.Request, params []string) map[string]string {
	formMap := make(map[string]string)

	for _, k := range params {
		formMap[k] = r.PostFormValue(k)
	}

	return formMap
}

func checkFormParams(r *http.Request, params []string) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	for _, k := range params {
		_, e := r.PostForm[k]
		if !e {
			return fmt.Errorf("%s is missing", k)
		}
		if r.PostFormValue(k) == "" {
			return fmt.Errorf("%s cannot be empty", k)
		}
	}

	return nil
}

func checkPasswordConfirmation(password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return fmt.Errorf("passwords mismatch")
	}
	return nil
}

func checkPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	if err != nil {
		return err
	}

	return nil
}

func checkDateAfter(date time.Time) error {
	if time.Now().After(date) {
		return fmt.Errorf("date already passed")
	}

	return nil
}
