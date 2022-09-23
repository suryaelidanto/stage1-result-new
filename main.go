package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{index}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")
	route.HandleFunc("/delete-blog/{index}", deleteBlog).Methods("GET")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))

	// var start = "2022-09-20"
	// var end = "2022-10-20"

	// t1, _ := time.Parse("2006-01-02", end)
	// t2, _ := time.Parse("2006-01-02", start)
	// hs := t1.Sub(t2).Hours()

	// hs, mf := math.Modf(hs)
	// ms := mf * 60

	// ms, sf := math.Modf(ms)
	// ss := sf * 60

	// fmt.Println(hs, "hours", ms, "minutes", ss, "seconds")
}

func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/add-blog.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

// var dataBlog = []
type Blog struct {
	Title     string
	Content   string
	Author    string
	Post_date string
}

var dataBlog = []Blog{
	{
		Title:   "Hallo Title",
		Content: "Hallo Content",
	},
	{
		Title:   "Hallo Title 2",
		Content: "Hallo Content 2",
	},
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var title = r.PostForm.Get("inputTitle")
	var content = r.PostForm.Get("inputContent")
	// var start = r.PostForm.Get("start")
	// var end = r.PostForm.Get("end")

	// fmt.Println(start)
	// fmt.Println(end)

	// t1 := time.Date(start)
	// t2 := time.Date(1984, time.November, 3, 10, 23, 34, 0, time.UTC)

	// hs := t1.Sub(t2).Hours()

	// hs, mf := math.Modf(hs)
	// ms := mf * 60

	// ms, sf := math.Modf(ms)
	// ss := sf * 60

	// fmt.Println(hs, "hours", ms, "minutes", ss, "seconds")

	// let blog = {
	// 	title,
	// 	content
	// }

	var newBlog = Blog{
		Title:     title,
		Content:   content,
		Author:    "Samsul Rijal",
		Post_date: time.Now().Format("2 January 2006 15:04"),
	}

	// dataBlog.push(blog)
	dataBlog = append(dataBlog, newBlog)
	// fmt.Println(dataBlog)

	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	response := map[string]interface{}{
		"Blogs": dataBlog,
	}

	tmpl.Execute(w, response)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var BlogDetail = Blog{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range dataBlog {
		if index == i {
			BlogDetail = Blog{
				Title:     data.Title,
				Content:   data.Content,
				Post_date: data.Post_date,
				Author:    data.Author,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}
	// fmt.Println(data)
	tmpl.Execute(w, data)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	// fmt.Println(index)

	dataBlog = append(dataBlog[:index], dataBlog[index+1:]...)
	// fmt.Println(dataBlog)

	http.Redirect(w, r, "/blog", http.StatusFound)
}
