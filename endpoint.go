package __projectname__

import (
  "github.com/GoogleCloudPlatform/go-endpoints/endpoints"
  "golang.org/x/net/context"
  "log"
)

var (
  scopes          = []string{endpoints.EmailScope}
  clientIds       = []string{endpoints.APIExplorerClientID}
  audiences       = []string{}
)

type __service__Service struct{}

type List__service__Req struct{}

type __service__List struct{
  Name string `json:"name"`
}

func (cs *__service__Service) List__service__(c context.Context, r *List__service__Req) (*__service__List, error) {
  list := new(__service__List)
  list.Name = "foo"
  return list, nil
}

func init() {
  api, err := endpoints.RegisterService(&__service__Service{}, "__projectname__", "v1", "__service__ API", true)
  if err != nil {
    log.Fatalf("Register service: %v", err)
  }

  list := api.MethodByName("List__service__").Info()
  // list.Scopes = scopes
  // list.ClientIds = clientIds
  list.HTTPMethod = "GET"
  list.Path = "__projectname__/list"
  list.Name = "list"
  list.Desc = "List __projectname__."

  endpoints.HandleHTTP()
}
