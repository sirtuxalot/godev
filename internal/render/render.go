/* internal/render/render.go */

package render

import (
  "bytes"
  "github.com/justinas/nosurf"
  "godev/internal/config"
  "godev/internal/models"
  "html/template"
  "log"
  "net/http"
  "path/filepath"
)

// access application config
var app *config.AppConfig

// sets config for render package
func NewTemplates(config *config.AppConfig) {
  app = config
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
  templateData.CSRFToken = nosurf.Token(request)
  return templateData
}

// function to render templates to browser
func Template(writer http.ResponseWriter, request *http.Request, tmpl string, templateData *models.TemplateData) {

  // create local templateCache variable
  templateCache := map[string]*template.Template{}

  // enable rebuilding of template cache while in development but not production
  if app.UseCache {
    // retrieve template cache from application config
    templateCache = app.TemplateCache
  } else {
    // rebuild template cache
    templateCache, _ = CreateTemplateCache()
  }

  // get requested template from cache
  pageTemplate, ok := templateCache[tmpl]
  if !ok {
    log.Fatal("Unable to retrieve template from cache")
  }

  // create bytes buffer to check for issues with the templates
  templateBuffer := new(bytes.Buffer)
  templateData = AddDefaultData(templateData, request)
  _ = pageTemplate.Execute(templateBuffer, templateData)

  // render the template
  _, err := templateBuffer.WriteTo(writer)
  if err != nil {
    log.Println(err)
  }
}

func CreateTemplateCache() (map[string]*template.Template, error) {

  // create local templateCache variable
  templateCache := map[string]*template.Template{}

  // get filenames of html templates (*.html) from the templates folder
  pages, err := filepath.Glob("./templates/*/*.html")
  if err != nil {
    return templateCache, err
  }

  // range through html templates the parse the templates
  for _, page := range pages {
    // parse out the name of the template
    fileName := filepath.Base(page)

    // add the template called fileName that is parsed from the page, and add it to the templateSet
    templateSet, err := template.New(fileName).ParseFiles(page)
    if err != nil {
      return templateCache, err
    }

    // get filenames of layout templates (*.layout.tmpl) from the templates folder
    layouts, err := filepath.Glob("./templates/*.layout.tmpl")
    if err != nil {
      return templateCache, err
    }

    if len(layouts) > 0 {
      // add the layout templates to the templateSet
      templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
      if err != nil {
        return templateCache, err
      }
    }

    // add the template set at the key of fileName to the templateCache map
    templateCache[fileName] = templateSet
  }

  log.Println("Template cache generated!")
  return templateCache, err
}
