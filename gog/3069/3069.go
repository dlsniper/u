package main

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/flatbuffers/go"
)

type dem struct {
	o kithttp.ErrorEncoder
}

func main() {
	_ = dem{
		o: func(ctx context.Context, err error, w http.ResponseWriter) {},
	}
}
