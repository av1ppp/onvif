// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var mainTemplate = `// Code generated : DO NOT EDIT.

package {{.Package}}

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/{{.StructPackage}}"
	"github.com/av1ppp/onvif/errors"
)

// Call_{{.TypeRequest}} forwards the call to dev.CallMethod() then parses the payload of the reply as a {{.TypeReply}}.
func Call_{{.TypeRequest}}(ctx context.Context, dev *onvif.Device, request {{.StructPackage}}.{{.TypeRequest}}) ({{.StructPackage}}.{{.TypeReply}}, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			{{.TypeReply}} {{.StructPackage}}.{{.TypeReply}}
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.{{.TypeReply}}, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "{{.TypeRequest}}")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.{{.TypeReply}}, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "{{.TypeRequest}}")
	}
	return reply.Body.{{.TypeReply}}, nil
}

// CallWithLogging_{{.TypeRequest}} works like Call_{{.TypeRequest}} but also logs the response body.
func CallWithLogging_{{.TypeRequest}}(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request {{.StructPackage}}.{{.TypeRequest}}) ({{.StructPackage}}.{{.TypeReply}}, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			{{.TypeReply}} {{.StructPackage}}.{{.TypeReply}}
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.{{.TypeReply}}, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "{{.TypeRequest}}")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "{{.TypeRequest}}")
	if err != nil {
		return reply.Body.{{.TypeReply}}, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "{{.TypeRequest}}")
	}
	return reply.Body.{{.TypeReply}}, nil
}
`

type parserEnv struct {
	Package       string
	StructPackage string
	TypeReply     string
	TypeRequest   string
}

func main() {
	flag.Parse()
	env := parserEnv{
		Package:       flag.Arg(0),
		StructPackage: flag.Arg(1),
		TypeRequest:   flag.Arg(2),
		TypeReply:     flag.Arg(2) + "Response",
	}

	log.Println(env)

	body, err := template.New("body").Parse(mainTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	fout, err := os.Create(env.TypeRequest + "_auto.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()

	err = body.Execute(fout, &env)
	if err != nil {
		log.Fatalln(err)
	}
}
