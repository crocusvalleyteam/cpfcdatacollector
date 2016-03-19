package main

import (


	"net/http"
	"fmt"
	//"io/ioutil"
	"log"
	"golang.org/x/net/html"
	"io"


)

//main -entrypoint
func main() {


	resp, err := http.Get("http://www.crystalpalace-mad.co.uk/results_fixtures/2015_2016/crystal_palace/index.shtml")

	dealwitherrors(err)

	y := html.NewTokenizer(resp.Body)


	//first find all table cells
	for {
		yy := y.Next()

		switch {
		case yy == html.ErrorToken:
			// End of the document, we're done
			return
		case yy == html.StartTagToken:
			t := y.Token()

			istablecell := t.Data == "td"
			if istablecell {
				fmt.Println("table cell found")

				//if cell found work on the content
				contenttokniser(resp.Body)

				defer resp.Body.Close()
			}
		}
	}




}

//contenttokniser takes as an urgument an io.Reader (the body html page to be processed)

func contenttokniser(htmlpage io.Reader) {


	z := html.NewTokenizer(htmlpage)



	for {


			tt := z.Next()


			switch {

			case tt == html.ErrorToken:
			// End of the document, we're done
			return

			//html page content
			case tt == html.TextToken:


				t := z.Text()

				fmt.Print(string(t))

			}


		}


}

func dealwitherrors(err error) {


	if err != nil {

		log.Fatal(err)

	}
}

