package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	connection.DatabaseConnect()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")
	route.HandleFunc("/delete-blog/{id}", deleteBlog).Methods("GET")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
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
	ID          int
	Title       string
	Content     string
	Author      string
	Post_date   time.Time
	Format_date string
}

// var dataBlog = []Blog{
// 	{
// 		Title:   "Hallo Title",
// 		Content: "Hallo Content",
// 	},
// 	{
// 		Title:   "Hallo Title 2",
// 		Content: "Hallo Content 2",
// 	},
// }

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var title = r.PostForm.Get("inputTitle")
	var content = r.PostForm.Get("inputContent")
	var author = r.PostForm.Get("inputAuthor")

	// var newBlog = Blog{
	// 	Title:     title,
	// 	Content:   content,
	// 	Author:    "Samsul Rijal",
	// 	Post_date: time.Now().Format("2 January 2006 15:04"),
	// }

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog(title, content, author) VALUES ($1, $2, $3)", title, content, author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content, post_date FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.Post_date)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Author = "Abel Dustin"
		each.Format_date = each.Post_date.Format("2 January 2006")

		result = append(result, each)
	}

	// connection.Conn.QueryRow(context.Background(), "SELECT id, title, content FROM tb_blog").Scan(&result)

	resData := map[string]interface{}{
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resData)
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

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, post_date FROM tb_blog WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Content, &BlogDetail.Post_date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
	}

	BlogDetail.Author = "Abel Dustin"
	BlogDetail.Format_date = BlogDetail.Post_date.Format("2 January 2006")

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}
	// fmt.Println(data)
	tmpl.Execute(w, data)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// fmt.Println(index)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
	}

	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}
