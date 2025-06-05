package server

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"muxHello/data"
	"muxHello/security"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type Server struct {
	DB       *os.File
	Writer   *csv.Writer
	Reader   *csv.Reader
	Records  []data.Record
	NewestID int
}

func NewServer(db *os.File) *Server {

	writer := csv.NewWriter(db)

	var rec []data.Record

	if err := gocsv.UnmarshalFile(db, &rec); err != nil {
		log.Fatalf("Failed unmarshalling file %s", db.Name())
		return nil
	}

	db.Seek(0, io.SeekStart)

	reader := csv.NewReader(db)

	newestID, err := getNewestID(db)
	if err != nil {
		log.Fatal("An error occurred getting the newest ID: ", err)
		return nil
	}

	return &Server{
		DB:       db,
		Writer:   writer,
		Reader:   reader,
		Records:  rec,
		NewestID: newestID,
	}

}

func getNewestID(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	fmt.Println("Number of lines:", lineCount)

	return lineCount, nil
}

func (s *Server) CloseServer() {

	s.Writer.Flush()
	s.DB.Close()

}

func (s *Server) FindEmail(email string) (*data.Record, error) {
	// update records
	s.DB.Seek(0, io.SeekStart)
	if err := gocsv.UnmarshalFile(s.DB, &s.Records); err != nil {
		return nil, err
	}

	for _, user := range s.Records {
		if user.Email == email {
			log.Printf("Found email: %s\n", user.Email)

			return &user, nil
		}
	}

	return nil, fmt.Errorf("email %s was not found in the database ", email)
}

func (s *Server) CheckLogin(data data.AccountData) bool {
	user, err := s.FindEmail(data.Email)
	if err != nil {
		return false
	}

	correctPassword := security.CheckPasswordHash(data.Password, user.PasswordHash)

	if correctPassword {
		log.Println("Password was correct")

		return true
	} else {
		log.Println("Password was incorrect.")

		return false
	}
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "/api/signup/ is a POST.", http.StatusMethodNotAllowed)
	}

	var data data.AccountData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Print("an error occurred parsing the JSON.", err)
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}

	loginWorked := s.CheckLogin(data)

	if loginWorked {
		log.Println("Password was correct")

		w.WriteHeader(http.StatusOK)
		return
	} else {
		log.Println("Username or password was not found.")

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "/api/signup/ is a POST.", http.StatusMethodNotAllowed)
	}

	var data data.AccountData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Print("an error occurred parsing the JSON.", err)
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}
	hash, err := security.HashPassword(data.Password)
	if err != nil {
		log.Fatal("error hashing password")
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	rec, err := s.FindEmail(data.Email)
	if err != nil {
		line := []string{fmt.Sprintf("%d", s.NewestID+1), data.Username, data.Email, hash}

		s.Writer.Write(line)
		s.Writer.Flush()

		http.Redirect(w, r, "./static/dashboard/dashboard.html", http.StatusFound)
		return
	} else if rec != nil && rec.Email == data.Email {
		http.Error(w, "Email already in use.", http.StatusConflict)
		return
	}

}
