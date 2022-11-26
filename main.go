package main // 

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter() // variabelnya

	//route untuk menginisialisai folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))) // ini utk manggil di folder publik

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/addblog", addBlog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/myProjec", myProjec).Methods("GET")
	route.HandleFunc("/form-myProjec", procesmyProjec).Methods("POST")

	fmt.Println("Server berjalan pada port 5000") // sama seperti console
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")// utk memanggil index.html

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil) // utk kirim
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/add-blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func myProjec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/myProjeck.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func procesmyProjec(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()


	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	// menampilkan data projec ke console terminal
	fmt.Println("Project Name: "+ r.PostForm.Get("project_name"))
	fmt.Println("Start Date: "+ r.PostForm.Get("start_date"))
	fmt.Println("End Date: "+ r.PostForm.Get("end_date"))
	fmt.Println("Description: "+ r.PostForm.Get("description"))
	fmt.Println("Technologies: ", r.Form["technologies"])

	http.Redirect(w, r, "/myProjec", http.StatusMovedPermanently)

}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// object golang
	data := map[string]interface{}{
		"Title":   "Jasa PEmotong Rumput Dari Dumbways",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan.",
		"Id":      id,
	}

	tmpt.Execute(w, data)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog.html")

	// if condition
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
