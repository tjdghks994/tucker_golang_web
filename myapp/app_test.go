package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("No Users", string(data))
}

func TestGetUserList(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sunghwan", "last_name":"park", "email":"tjdghks994@gmail.com"}`),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"jihong", "last_name":"park", "email":"tjdghks994@gmail.com"}`),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:56")
}

func TestCreatUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sunghwan", "last_name":"park", "email":"tjdghks994@gmail.com"}`),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)

}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "first_name":"update", "last_name":"update", "email":"update@gamil.com"}`),
	)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"id":1, "first_name":"update", "last_name":"update", "email":"update@gamil.com"}`),
	)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "first_nama":"update"}`, user.ID)

	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr),
	)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("update", updateUser.FirstName)
	assert.Equal(user.LastName, updateUser.LastName)
	assert.Equal(user.Email, updateUser.Email)
}
