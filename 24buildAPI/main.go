package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//controllers
// serve home route

func servehome(w http.ResponseWriter, r *http.Request) { //order of writer and reader is imp
	w.Write([]byte("<h1>Home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses!!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course!!")
	w.Header().Set("Content-Type", "application/json")
	//garb id from req(r)
	params := mux.Vars(r)

	// loop through courses ,, and find matching id and return res

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// if course not found
	json.NewEncoder(w).Encode("NO COURSE FOUND")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One course!!")
	w.Header().Set("Content-Type", "application/json")

	// if request is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("EMPTY REQUEST")
	}

	// data sent and {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("PLEASE SEND PROPER DATA")
		return
	}

	//generate userid, string
	// append course to courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update One Course!!")
	w.Header().Set("Content-Type", "application/json")

	// first get id from req
	params := mux.Vars(r)
	// loop and find course and remove and then again add the same course with same id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		} else {
			// TODO when id not found
		}

	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete One Course!!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("COURSE DELETED")
			break
		}
	}
}

// fake DB
var courses []Course

// middleware checking id courseid and name is null of not
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API-Learnign")
	r := mux.NewRouter()

	//seeding data
	// ?? we need to pass address of author as its a pointer in the structs ??
	courses = append(courses, Course{CourseId: "20", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Adarsh Kumar", Website: "aksite"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "ElectronJS", CoursePrice: 400, Author: &Author{Fullname: "Adarsh Kumar", Website: "aksite2"}})

	//routing
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") //params put in { } and must be same name as the one used in function called
	r.HandleFunc("/createCourse", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listeing
	log.Fatal(http.ListenAndServe(":4000", r))

}
