package handler

import (
	"golang.org/x/net/context"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	proto "github.com/Rakanixu/factorial-micro-service/server/proto"
	"github.com/micro/go-micro/client"
	"encoding/json"
	"strconv"
)

type Factorial struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

func (f *Factorial) Read(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var cr *proto.FactorialRequest
	var err error

	if d := extractValue(req.Get["Number"]); len(d) > 0 {
		num, _ := strconv.ParseInt(d, 10, 64)
		request := client.NewRequest(
			"go.micro.srv.factorial",
			"Factorial.CalcFactorial",
			&proto.FactorialRequest{
				Number: num,
			},
		)

		response := &proto.FactorialResponse{}

		if err = client.Call(ctx, request, response); err != nil {
			return errors.InternalServerError("go.micro.srv.factorial", err.Error())
		}

		b, _ := json.Marshal(response.Result)

		rsp.StatusCode = 200
		rsp.Body = string(b)

	} else {
		err = json.Unmarshal([]byte(req.Body), &cr)
	}

	if err != nil {
		return errors.BadRequest("go.micro.api.factorial", "invalid data")
	}

	return nil
}
