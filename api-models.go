package main

//functions to create models rendered by the API
import (
  "net/http"
  "github.com/xogeny/go-siren"
  "strconv"
)

type API struct {
  entity *gosiren.SirenEntity
}

func (a *API) SetTitle(title string) {
  a.entity.Title = title
}

func (a *API) SetClasses(classes []string) {
  a.entity.Class = classes
}

func (a *API) SetProperties(properties map[string]interface{} ) {
  a.entity.Properties = properties
}

func (a *API) SetEntities(entities []gosiren.SirenEmbed) {
  a.entity.Entities = entities
}

func (a *API) SetActions(actions []gosiren.SirenAction) {
  a.entity.Actions = actions
}

func (a *API) SetLinks(links []gosiren.SirenLink) {
  a.entity.Links = links
}

func (a *API) AddLink(rel []string, href string, l gosiren.SirenLink) {
  a.entity.AddLink(rel, href, l)
}

func (a *API) AddAction(name string, title string, method string,
	href string, ctype string, fields ...gosiren.SirenField) {
  a.entity.AddAction(name, title, method, href, ctype, fields...)
}

func CreateRootModel(req *http.Request) *gosiren.SirenEntity {

  a := new(API)
  a.entity = gosiren.NewSirenEntity()
  upLink := "http://" + req.Host + "/"
  selfLink := "http://" + req.Host + req.URL.Path
  tickerLink := "ws://" + req.Host + "/ticker"
  postLink := selfLink + "create"

  // root := gosiren.NewSirenEntity()
  // root.Title = "Root Document"
  // root.Class = []string{"root"}
  // root.Properties = map[string]interface{}{
  //   "world": "game-01",
  // }


  a.SetTitle("Root Document")
  a.SetClasses([]string{"root"})
  a.SetProperties(map[string]interface{}{
    "world": "game-01",
  })

  a.AddLink([]string{"self"}, selfLink, gosiren.SirenLink{})
  a.AddLink([]string{"up"}, upLink, gosiren.SirenLink{})
  a.AddLink([]string{"monitor"}, tickerLink, gosiren.SirenLink{})
  a.AddAction("start", "Create Game", "POST", postLink, "application/x-www-form-urlencoded", gosiren.MakeField("name", "string", ""))
  return a.entity
}

func CreatePlayerModel(p Player, req *http.Request) *gosiren.SirenEntity {

  a := new(API)
  a.entity = gosiren.NewSirenEntity()
  upLink := "http://" + req.Host + "/"
  selfLink := "http://" + req.Host + req.URL.Path
  tickerLink := "ws://" + req.Host + "/ticker"

  // root := gosiren.NewSirenEntity()
  // root.Title = "Root Document"
  // root.Class = []string{"root"}
  // root.Properties = map[string]interface{}{
  //   "world": "game-01",
  // }


  a.SetTitle("Player Document")
  a.SetClasses([]string{"player"})
  a.SetProperties(map[string]interface{}{
    "name": p.name,
    "id": p.id,
    "money": p.money,
    "shares": p.shares,
  })

  a.AddLink([]string{"up"}, upLink, gosiren.SirenLink{})
  a.AddLink([]string{"self"}, selfLink, gosiren.SirenLink{})
  a.AddLink([]string{"monitor"}, tickerLink, gosiren.SirenLink{})

  return a.entity
}

func CreateCompanyList(m Market, req *http.Request) *gosiren.SirenEntity {
  a := new(API)
  a.entity = gosiren.NewSirenEntity()
  upLink := "http://" + req.Host + "/"
  selfLink := "http://" + req.Host + req.URL.Path
  tickerLink := "ws://" + req.Host + "/ticker"

  // root := gosiren.NewSirenEntity()
  // root.Title = "Root Document"
  // root.Class = []string{"root"}
  // root.Properties = map[string]interface{}{
  //   "world": "game-01",
  // }


  a.SetTitle("Companies Document")
  a.SetClasses([]string{"company", "collection"})
  a.SetProperties(map[string]interface{}{})

  for _, company := range m {
    id := company.id
    url := upLink + "companies/" + strconv.FormatInt(int64(id), 10)
    embeddedEntity := gosiren.SirenEmbed{
    	Rel: []string{"item"},
    	Href: url,
    	Type: "company",
    	Title: company.name + " Document",
    	Class: []string{"company"},
    	Properties:map[string]interface{}{
        "id": company.id,
        "name": company.name,
        "money": company.money,
        "shares": company.shares,
        "pricePerShare": company.pricePerShare,
      },
    	Entities:[]gosiren.SirenEmbed{},
    	Actions:[]gosiren.SirenAction{},
    	Links:[]gosiren.SirenLink{},
    }
    a.entity.AddEmbed([]string{"item"}, embeddedEntity)
  }

  a.AddLink([]string{"up"}, upLink, gosiren.SirenLink{})
  a.AddLink([]string{"self"}, selfLink, gosiren.SirenLink{})
  a.AddLink([]string{"monitor"}, tickerLink, gosiren.SirenLink{})

  return a.entity
}
