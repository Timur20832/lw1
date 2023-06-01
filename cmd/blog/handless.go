package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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

// start admin
type AdminPage struct {
	AdminHeader []adminheaderdata
	MainInfo    []maininfodata
	FullPage    []fullpagedata
}

type adminheaderdata struct {
	ImageLogo1    string
	ImageLogo2    string
	FirstCharName string
	ImageExit     string
}

type maininfodata struct {
	Title    string
	Subtitle string
	Button   string
}

//end admin

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

func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := AdminPage{
		AdminHeader: adminheader(),
		MainInfo:    maininfo(),
		FullPage:    fullpage(),
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func adminheader() []adminheaderdata {
	return []adminheaderdata{
		{
			ImageLogo1:    "../static/image/escapeforadmin.svg",
			ImageLogo2:    "../static/image/authorforadmin.svg",
			FirstCharName: "K",
			ImageExit:     "../static/image/log-out.svg",
		},
	}
}

func maininfo() []maininfodata {
	return []maininfodata{
		{
			Title:    "New Post",
			Subtitle: "Fill out the form bellow and publish your article",
			Button:   "Publish",
		},
	}
}

type fullpagedata struct {
	MainTitle       string
	Title           string
	Description     string
	AuthorName      string
	TextAuthorphoto string
	Authorphoto     string
	TextDate        string
	ImageCamera     string
	Upload          string
	TextImage       string

	PreviewArticle1 string
	PreviewImage1   string

	PreviewArticle2 string
	PreviewImage2   string

	TitleInPreview           string
	SubtitleInPreview        string
	PreviewImageChange       string
	PreviewImageChangeAuthor string
	RandomName               string
	RandomDate               string

	FooterTitle    string
	FooterSubtitle string

	ImageMegaCamera string
	ImageMegaTrash  string
}

func fullpage() []fullpagedata {
	return []fullpagedata{
		{
			MainTitle:                "Main Information",
			Title:                    "Title",
			Description:              "Short description",
			AuthorName:               "Author name",
			TextAuthorphoto:          "Author Photo",
			Authorphoto:              "../static/image/camera-without.svg",
			TextDate:                 "Publish Date",
			ImageCamera:              "../static/image/camera.svg",
			Upload:                   "Upload",
			TextImage:                "Hero Image",
			PreviewArticle1:          "Article preview",
			PreviewImage1:            "../static/image/megaphone.jpg",
			PreviewArticle2:          "Post card preview",
			PreviewImage2:            "../static/image/smallphone.jpg",
			TitleInPreview:           "New Post",
			SubtitleInPreview:        "Please, enter any description",
			PreviewImageChangeAuthor: "../static/image/blankLogo.png",
			PreviewImageChange:       "../static/image/kek.jpg",
			RandomName:               "Enter author name",
			RandomDate:               "4/19/2023",
			FooterTitle:              "Content",
			FooterSubtitle:           "Post content (plain text)",
			ImageMegaCamera:          "../static/image/camera.png",
			ImageMegaTrash:           "../static/image/trash-2.png",
		},
	}
}

type createPostRequest struct {
	Title           string `json:"title_g"`
	Description     string `json:"subtitle_g"`
	AuthorName      string `json:"author_name_g"`
	AuthorPhoto     string `json:"author_url_name"`
	AuthorPhotoName string `json:"author_url_name_base64"`
	Date            string `json:"date_g"`
	BigImage        string `json:"big_image_name"`
	BigImageName    string `json:"big_image_name_base64"`
	SmallImage      string `json:"small_image_name"`
	SmallImageName  string `json:"small_image_name_base64"`
	ContentPost     string `json:"text_area_content_g"`
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "1Error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "2Error", 500)
			log.Println(err.Error())
			return
		}

		b64Author := req.AuthorPhotoName[strings.IndexByte(req.AuthorPhotoName, ',')+1:]
		authorImg, err := base64.StdEncoding.DecodeString(b64Author)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileAuthor, err := os.Create("static/image/" + req.AuthorPhoto)
		if err != nil {
			http.Error(w, "file", 500)
			fmt.Println(err.Error())
			return
		}

		_, err = fileAuthor.Write(authorImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		b64Big := req.BigImageName[strings.IndexByte(req.BigImageName, ',')+1:]
		bigImg, err := base64.StdEncoding.DecodeString(b64Big)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileBig, err := os.Create("static/image/" + req.BigImage)
		if err != nil {
			http.Error(w, "file", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileBig.Write(bigImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		b64Small := req.SmallImageName[strings.IndexByte(req.SmallImageName, ',')+1:]
		smallImg, err := base64.StdEncoding.DecodeString(b64Small)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileSmall, err := os.Create("static/image/" + req.SmallImage)
		if err != nil {
			http.Error(w, "file", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileSmall.Write(smallImg)
		if err != nil {
			http.Error(w, "write", 500)
			log.Println(err.Error())
			return
		}

		err = saveOrder(db, req)
		if err != nil {
			http.Error(w, "bd", 500)
			log.Println(err.Error())
			return
		}
		return
	}
}

func saveOrder(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			subtitle,
			author_name,
			author_img,
			publish_date,
			image_post,
			contant,
			featured
		)
		VALUES
		(
			?,
			?,
			?,
			CONCAT('../static/image/', ?),
			?,
			CONCAT('../static/image/', ?),
			?,
			?
		)
	`

	_, err := db.Exec(query, req.Title, req.Description, req.AuthorName, req.AuthorPhoto, req.Date, req.SmallImage, req.ContentPost, 0)
	return err
}

/*
func formatDate(oldDate string) string {
	dateStr := strings.Split(oldDate, "-")
	newDateStr := dateStr[2] + "/" + dateStr[1] + "/" + dateStr[0]
	return newDateStr
}
*/
