package repository

import (
  "fmt"
  "context"
  "google.golang.org/grpc/metadata"
)

// IBiz dashboard service business logic
type IBiz interface {}

func extractAuthInfoFromContext(ctx context.Context) (string, string, []string, error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return "", "", nil, fmt.Errorf("fail to extract auth info from ctx")
  }
  if len(md["uid"]) < 1 {
    return "", "", nil, fmt.Errorf("uid empty")
  }
  if len(md["nid"]) < 1 {
    return "", "", nil, fmt.Errorf("nid empty")
  }

  if len(md["roles"]) < 1 {
    return "", "", nil, fmt.Errorf("roles empty")
  }

  return md["uid"][0], md["nid"][0], md["roles"], nil
}
