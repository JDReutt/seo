package seo

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"
)

// MicroProduct micro product definition, ref: https://developers.google.com/structured-data/rich-snippets/products
type MicroProduct struct {
	Name            string
	Image           string
	Description     string
	BrandName       string
	SKU             string
	RatingValue     float32
	ReviewCount     int
	PriceCurrency   string
	Price           float64
	PriceValidUntil string
	SellerName      string
}

// MicroSearch micro search definition, ref: https://developers.google.com/structured-data/slsb-overview
// e.g.
//   Target: https://query.example-petstore.com/search?q={keyword}
type MicroSearch struct {
	URL        string
	Target     string
	QueryInput string
}

// MicroContact micro search definition, ref: https://developers.google.com/structured-data/customize/contact-points
type MicroContact struct {
	URL         string
	Telephone   string
	ContactType string
}

// Render render micro product structured data
func (product MicroProduct) Render() template.HTML {
	return renderTemplate("product.tmpl", product)
}

// Render render micro search structured data
func (search MicroSearch) Render() template.HTML {
	return renderTemplate("search.tmpl", search)
}

// FormattedQueryInput format query input
func (search MicroSearch) FormattedQueryInput() string {
	if search.QueryInput == "" {
		return "required name=keyword"
	}
	return search.QueryInput
}

// Render render micro contact structured data
func (contact MicroContact) Render() template.HTML {
	return renderTemplate("contact.tmpl", contact)
}

func renderTemplate(templateName string, obj interface{}) template.HTML {
	var templatePath string
	for _, gopath := range strings.Split(os.Getenv("GOPATH"), ":") {
		templatePath = path.Join(gopath, "src/github.com/qor/seo/views/microdata/", templateName)
	}
	if tmpl, err := template.ParseFiles(templatePath); err == nil {
		var datas bytes.Buffer
		if err = tmpl.Execute(&datas, obj); err != nil {
			fmt.Println(err)
		} else {
			return template.HTML(datas.String())
		}
	} else {
		panic(err)
	}
	return template.HTML("")
}
