package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// only post
type postPage struct {
	HeaderPost []headerpostdata
	PostInfo   []postinfodata
	Footer     []footerdata
}

type headerpostdata struct {
	Escape string
	Nav    []navdata
}

type postinfodata struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Image    string `db:"image_post"`
	Content  string `db:"contant"`
}

// only post but footer in index

// only main page
type indexPage struct {
	Header     []headerdata
	Menu       []menudata
	TitleBpost string
	Bposts     []*bpostsdata
	TitleSpost string
	Sposts     []*spostsdata
	Footer     []footerdata
}

// start header
type headerdata struct {
	Escape   string
	Nav      []navdata
	TopTitle []toptitledata
}

type navdata struct {
	First    string
	FirstURL string
	Second   string
	Third    string
	Fourth   string
}

type toptitledata struct {
	Title  string
	Text   string
	Button string
}

// end header

// start menu
type menudata struct {
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
}

// end menu

// start posts
type bpostsdata struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	AuthorName  string `db:"author_name"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
	ImagePost   string `db:"image_post"`
	PostURL     string
}

type spostsdata struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	AuthorName  string `db:"author_name"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
	ImagePost   string `db:"image_post"`
	PostURL     string
}

//end posts

// start footer
type footerdata struct {
	FormFound    []dataformfound
	NavbarFooter []datanavbarfooter
}

type dataformfound struct {
	Title  string
	Button string
}

type datanavbarfooter struct {
	Escape string
	Nav    []navdata
}

// end footer

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bposts, err := bposts(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		sposts, err := sposts(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			Header:     header(),
			Menu:       Menu(),
			TitleBpost: "Feaurted Post",
			Bposts:     bposts,
			TitleSpost: "Most Recent",
			Sposts:     sposts,
			Footer:     footer(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Mega Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error 1111", 500)
			log.Println(err.Error())
			return
		}

		data := postPage{
			HeaderPost: headerpost(),
			PostInfo:   post,
			Footer:     footer(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error 2222", 500)
			log.Println(err.Error())
			return
		}
	}
}

func postByID(db *sqlx.DB, postID int) ([]postinfodata, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  image_post,
		  contant
		FROM
		  post
	    WHERE
		  post_id = ?
	`

	var post []postinfodata

	err := db.Select(&post, query, postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func headerpost() []headerpostdata {
	return []headerpostdata{
		{
			Escape: "../static/image/Escape_Black.png",
			Nav:    nav(),
		},
	}
}

func header() []headerdata {
	return []headerdata{
		{
			Escape:   "static/image/Escape..svg",
			Nav:      nav(),
			TopTitle: toptitle(),
		},
	}
}

func toptitle() []toptitledata {
	return []toptitledata{
		{
			Title:  "Let's do it together.",
			Text:   "We travel the world in search of stories. Come along for the ride.",
			Button: "View Latest Posts",
		},
	}
}

func nav() []navdata {
	return []navdata{
		{
			First:    "HOME",
			FirstURL: "/home",
			Second:   "CATEGORIES",
			Third:    "ABOUT",
			Fourth:   "CONTACT",
		},
	}
}

func Menu() []menudata {
	return []menudata{
		{
			One:   "Nature",
			Two:   "Photography",
			Three: "Relaxation",
			Four:  "Vacation",
			Five:  "Travel",
			Six:   "Adventure",
		},
	}
}

func bposts(db *sqlx.DB) ([]*bpostsdata, error) {
	const query = `
	SELECT
		  post_id,
		  title,
		  subtitle,
		  author_name,
		  author_img,
		  publish_date,
		  image_post
	FROM
		  post
	WHERE featured = 1
	`

	var bposts []*bpostsdata

	err := db.Select(&bposts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range bposts {
		post.PostURL = "/post/" + post.PostID
	}

	log.Println(bposts)

	return bposts, nil
}

func sposts(db *sqlx.DB) ([]*spostsdata, error) {
	const query = `
	SELECT
		  post_id,
		  title,
		  subtitle,
		  author_name,
		  author_img,
		  publish_date,
		  image_post
	FROM
		  post
	WHERE featured = 0
	`

	var sposts []*spostsdata

	err := db.Select(&sposts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range sposts {
		post.PostURL = "/post/" + post.PostID
	}

	log.Println(sposts)

	return sposts, nil
}

func footer() []footerdata {
	return []footerdata{
		{
			FormFound:    formfounder(),
			NavbarFooter: navbarfooter(),
		},
	}
}

func formfounder() []dataformfound {
	return []dataformfound{
		{
			Title:  "Stay in Touch",
			Button: "Sumbit",
		},
	}
}

func navbarfooter() []datanavbarfooter {
	return []datanavbarfooter{
		{
			Escape: "../static/image/Escape..svg",
			Nav:    nav(),
		},
	}
}
