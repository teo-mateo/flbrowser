package filelist

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	//"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

var flurl, _ = url.Parse("https://filelist.ro")
var flurlLogin, _ = url.Parse("https://filelist.ro/takelogin.php")
var flurlBrowse, _ = url.Parse("https://filelist.ro/browse.php")

var client *http.Client

func createClient() error {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	client = &http.Client{
		Jar: jar,
	}

	return nil
}

// Login to filelist
func login() error {

	if client == nil {
		err := createClient()
		if err != nil {
			return err
		}
	}

	if len(client.Jar.Cookies(flurl)) == 4 {
		return nil
	}

	fmt.Printf("Cookies before call: %d\n", len(client.Jar.Cookies(flurl)))

	_, err := client.Get(flurl.String())
	if err != nil {
		return err
	}

	fmt.Printf("Cookies after first call: %d\n", len(client.Jar.Cookies(flurl)))
	for _, c := range client.Jar.Cookies(flurl) {
		fmt.Printf("%s\t\t%s\n", c.Name, c.Value)
	}

	u, p, err := getFLAuth()
	if err != nil{
		return err
	}

	//log in
	data := url.Values{
		"username": []string{u},
		"password": []string{p},
	}

	_, err = client.PostForm(flurlLogin.String(), data)

	fmt.Printf("Cookies after login (%s): %d\n", flurl.String(), len(client.Jar.Cookies(flurl)))
	for _, c := range client.Jar.Cookies(flurl) {
		fmt.Printf("%s\t\t%s\n", c.Name, c.Value)
	}

	return err
}

// GetTorrents returns torrents
func GetTorrents(category int, page int) ([]TorrentInfo, error) {

	login()

	q := flurlBrowse.Query()
	q.Add("search", "")
	q.Add("cat", strconv.Itoa(category))
	q.Add("searchin", "1")
	q.Add("sort", "2")
	q.Add("page", strconv.Itoa(page))
	flurlBrowse.RawQuery = q.Encode()

	fmt.Printf("%s\n", flurlBrowse.String())
	fmt.Printf("%s\n", flurlBrowse.String())

	response, err := client.Get(flurlBrowse.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	torrents := make([]TorrentInfo, 0)
	doc.Find(".torrentrow").Each(func(i int, s *goquery.Selection) {

		ti := TorrentInfo{}

		// torrent name
		html, err := s.Find("b").First().Html()
		if err == nil {
			ti.Name = html
		}

		//torrent dl url
		s.Find("a").Each(func(i int, s *goquery.Selection) {

			href, ok := s.Attr("href")
			if ok {
				if strings.Index(href, "download.php") == 0 && strings.Index(href, "usetoken") == -1 {
					ti.DlURL = href
				}
			}
		})

		s.Find("div").Each(func (i int, s *goquery.Selection){
			html, err = s.Html()
			if err != nil{
				fmt.Printf("%d --> %s\n", i, err.Error())
			} else {
				fmt.Printf("%d --> %s\n", i, html)
			}

			switch i {
			case 1:
				name, err := s.Find("b").Html()
				if err == nil { ti.Name = name }
				break
			case 3:
				dlurl, exists:= s.Find("a").Attr("href")
				if exists { ti.DlURL = dlurl}
				break
			case 5:
				dateadded, err := s.Find("font").Html()
				if err == nil {
					ti.DateAded = strings.Replace(dateadded, "<br/>", " ", 1)
					}
			case 6:
				size, err := s.Find("font").Html()
				if err == nil{
					ti.Size = strings.Replace(size, "<br/>", " ", 1)
				}
			case 7:
				timesdownloaded, err := s.Find("font").Html()
				if err == nil{
					p1:=strings.Split(timesdownloaded, "<br/>")[0]
					ti.TimesDownloaded = p1
				} else {
					ti.TimesDownloaded = "N/A"
				}
			case 8:
				seeders, err := s.Find("font").Html()
				if err == nil{
					ti.Seeders = seeders
				} else {
					ti.Seeders = "N/A"
				}
				case 9:
					fmt.Println(s.Html())
					leechers, err := s.Find("b").Html()
					if err == nil{
						ti.Leechers = leechers
					} else {
						ti.Leechers = "N/A"
					}
			}

		})

		fmt.Printf("%v\n", ti)
		torrents = append(torrents, ti)

	})

	return torrents, nil

}

// TorrentInfo - torrent information
type TorrentInfo struct {
	Name  string `json:"name"`
	DlURL string `json:"dlurl"`
	DateAded string `json:"dateadded"`
	Size string `json:"size"`
	TimesDownloaded string `json:"timesdownloaded"`
	Seeders string `json:"seeders"`
	Leechers string `json:"leechers"`
}
