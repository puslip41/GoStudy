package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
	"strings"
	"errors"
	"github.com/puslip41/GoStudy/restapi"
	"encoding/json"
)

func Test_UserHandler_Get_Notfound(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(UserHandler))
	defer server.Close()

	if response, err := http.Get(server.URL + "/users/puslip41/"); err != nil {
		t.Errorf("unexpacted exception: %s", err.Error())
	} else {
		if response.StatusCode != http.StatusNotFound {
			t.Errorf("expacted: %d, got: %d", http.StatusNotFound, response.StatusCode)
		}
	}
}

func Test_UserHandler_Get_Notfound_Request(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(UserHandler))
	defer server.Close()

	values := url.Values{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users/puslip41", server.URL), strings.NewReader(values.Encode()))
	client := http.Client{}

	if response, err := client.Do(req); err != nil {
		t.Errorf("unexpacted exception: %s", err.Error())
	} else {
		if response.StatusCode != http.StatusNotFound {
			t.Errorf("expacted: %d, got: %d", http.StatusNotFound, response.StatusCode)
		}
	}
}

func Test_UserHandler_Get(t *testing.T) {

}

func Test_UserHandler_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(UserHandler))
	defer server.Close()

	values := url.Values{}
	values.Add("password", "puslip41dkagh")

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/users/puslip41", server.URL), strings.NewReader(values.Encode()))
	client := http.Client{}

	if response, err := client.Do(req); err != nil {
		t.Errorf("unexpacted exception: %s", err.Error())
	} else {
		if response.StatusCode != http.StatusNotFound {
			b, _ := ioutil.ReadAll(response.Body)

			t.Errorf("expacted: %d, got: %d / %s", http.StatusCreated, response.StatusCode, string(b))
		}
	}
}

func Test_UserHandler_Put(t *testing.T) {

}

func Test_UserHandler_Post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(UserHandler))
	defer server.Close()

	values := url.Values{}
	values.Add("password", "puslip41dkagh")
	values.Add("name", "kim pilseop")
	values.Add("email", "puslip41@igloosec.com")

	if response, err := http.PostForm(fmt.Sprintf("%s/users/puslip41", server.URL), url.Values{"password":{"puslip41dkagh"}, "name":{"kim pilseop"}, "email":{"puslip41@igloosec.com"}}); err != nil {
		t.Errorf("unexpacted exception: %s", err.Error())
	} else {
		if response.StatusCode != http.StatusCreated {
			b, _ := ioutil.ReadAll(response.Body)
			t.Errorf("expacted: %d, got: %d / %s", http.StatusCreated, response.StatusCode, string(b))
		} else {
			if id, password, name, email, err := getMember(server, "puslip41"); err != nil {
				t.Errorf("unexpacted exception: %s", err.Error())
			} else {
				if id != "puslip41" {
					t.Errorf("expacted: puslip41, got: %s", id)
				}
				if password != "puslip41" {
					t.Errorf("expacted: puslip41, got: %s", password)
				}
				if name != "puslip41" {
					t.Errorf("expacted: puslip41, got: %s", name)
				}
				if email != "puslip41" {
					t.Errorf("expacted: puslip41, got: %s", email)
				}
			}
		}
	}
}

func getMember(server *httptest.Server, member_id string) (id, password, name, email string, func_err error) {
	if response, err := http.Get(fmt.Sprintf("%s/users/%s", server.URL, member_id)); err != nil {
		func_err = err
		return
	} else {
		if response.StatusCode == http.StatusFound {
			defer response.Body.Close()

			if b, err := ioutil.ReadAll(response.Body); err != nil {
				func_err = err
				return
			} else {
				item := restapi.StorageItem{}
				if err := json.Unmarshal(b, &item); err != nil {
					func_err = err
					return
				} else {
					id = item.ID
					password = item.Password
					name = item.Name
					email = item.EMail
					return
				}
				func_err = errors.New(string(b))
			}
		} else {
			func_err = errors.New(fmt.Sprintf("not found info: %d <= %s", response.StatusCode, fmt.Sprintf("%s/users/%s", server.URL, member_id)))
		}
	}

	return
}
