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

type __service__ struct{}

type List__service__Req struct{}

type __service__List struct{}

func (cs *__service__) List__service__(c context.Context, r *List__service__Req) (*__service__List, error) {
  list := new(__service__List)
  return list, nil
}
