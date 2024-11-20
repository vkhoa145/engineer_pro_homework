package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vkhoa145/engineer_pro_homework/day3/application/models"
	"github.com/vkhoa145/engineer_pro_homework/day3/application/utils"
)

func main() {
	userStorage := "users.txt"
	_, err := os.Stat(userStorage)
	if os.IsNotExist(err) {
		file, err := os.Create(userStorage)
		if err != nil {
			return
		}
		defer file.Close()
	}

	router := mux.NewRouter()
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	router.HandleFunc("/signup", signUpHandler).Methods("POST")
	router.HandleFunc("/signin", signInHandler).Methods("POST")
	router.HandleFunc("/users/{id}", getUserDetailHanlder).Methods("GET")
	router.HandleFunc("/users/{id}", editUserDetailHandler).Methods("PUT")
	http.ListenAndServe(":3000", router)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	existedUser := utils.FindUser(1, r.FormValue("username"))
	if existedUser != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User already existed")
		return
	}

	userFile := utils.OpenFile()
	defer userFile.Close()

	var lastline string

	scanner := bufio.NewScanner(userFile)
	for scanner.Scan() {
		lastline = scanner.Text()
	}

	var UserID int
	if lastline == "" {
		UserID = 1
	} else {
		parts := strings.Split(lastline, ",")
		intValue, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error converting '%s' to int: %v\n", parts[0], err)
		}

		UserID = intValue + 1
	}

	filename := utils.UploadImage(UserID, r)
	user := &models.User{
		Username:    r.FormValue("username"),
		Password:    r.FormValue("password"),
		ID:          UserID,
		UserProfile: filename,
	}
	userInfo := fmt.Sprintf("%d,%s,%s,%s\n", user.ID, user.Username, user.Password, filename)
	if _, err := userFile.WriteString(userInfo); err != nil {
		log.Printf("Failed to write user: %v", err)
	}
	json.NewEncoder(w).Encode(user)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	existedUser := utils.FindUser(1, r.FormValue("username"))
	if existedUser == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User not existed")
		return
	}

	if existedUser.Password != r.FormValue("password") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong Password")
		return
	}

	json.NewEncoder(w).Encode("Login successfull")
}

func getUserDetailHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	existedUser := utils.FindUser(0, id)
	if existedUser == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not existed")
		return
	}

	imagePath := utils.GetImage(existedUser.UserProfile, r)
	userDetail := &models.User{
		ID:          existedUser.ID,
		Username:    existedUser.Username,
		UserProfile: imagePath,
	}

	json.NewEncoder(w).Encode(userDetail)
}

func editUserDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	existedUser := utils.FindUser(0, id)
	if existedUser == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not existed")
		return
	}

	userFile := utils.OpenFile()
	defer userFile.Close()

	var lines []string

	scanner := bufio.NewScanner(userFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var newUsername string
	if existedUser.Username != username {
		existedUsername := utils.FindUser(1, username)
		if existedUsername != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Username is existed")
			return
		}
		newUsername = username
	} else {
		newUsername = existedUser.Username
	}

	var newPassword string
	if existedUser.Password != password {
		newPassword = password
	} else {
		newPassword = existedUser.Password
	}

	err := utils.DeleteImage(existedUser.UserProfile)
	if err != nil {
		fmt.Printf("error delete image: %v", err)
		return
	}

	var newUserProfile string
	filename := utils.UploadImage(existedUser.ID, r)
	if filename == "" {
		newUserProfile = existedUser.UserProfile
	} else {
		newUserProfile = filename
	}

	updatedUser := &models.User{
		ID:          existedUser.ID,
		Username:    newUsername,
		Password:    newPassword,
		UserProfile: newUserProfile,
	}

	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) > 0 && parts[0] == strconv.Itoa(existedUser.ID) {
			lines[i] = fmt.Sprintf("%d,%s,%s,%s", updatedUser.ID, updatedUser.Username, updatedUser.Password, updatedUser.UserProfile)
			break
		}
	}

	file, err := os.OpenFile("users.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("error opening file for writing: %v", err)
		return
	}
	defer file.Close()

	writter := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writter.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	writter.Flush()

	json.NewEncoder(w).Encode(updatedUser)
}
