package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/gen"
	"github.com/pkg/errors"
)

func main() {
	if err := generate(); err != nil {
		log.WithError(err).Fatal("error generating")
	}
}

func generate() error {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("couldn't get file from runtime")
	}
	dirname := path.Dir(file)

	vdomPath := path.Join(dirname, "..", "..", "vdom")
	log.Infof(vdomPath)

	buf, err := ioutil.ReadFile(path.Join(dirname, "inputs", "tags.json"))
	if err != nil {
		return errors.Wrapf(err, "error reading tags.json")
	}

	var data struct {
		Tags map[string]struct {
			Comment string
			Attrs   []string
		}
		Global []string
		Events []string
		Types  map[string]struct {
			Type   string
			Values []string
		}
	}

	if err := json.Unmarshal(buf, &data); err != nil {
		return errors.Wrapf(err, "error unmarshalling tags")
	}

	for name, tag := range data.Tags {
		type Attr struct {
			Key   string
			Value gen.Vartype
		}

		d := struct {
			Name    string
			Comment string
			Attrs   []Attr
		}{
			Name:    name,
			Comment: tag.Comment,
		}

		attrs := append(tag.Attrs, data.Global...)
		attrs = append(attrs, data.Events...)

		for _, attr := range attrs {
			parts := strings.Split(attr, ":")
			key := parts[0]
			kind := "string"

			if len(parts) > 2 {
				key = strings.Join(parts[0:len(parts)-2], ":")
				kind = parts[len(parts)-1]
			} else if len(parts) == 2 {
				kind = parts[1]
			}

			if _, isset := data.Types[kind]; isset {
				log.Infof("kind isset %s", kind)
			}

			d.Attrs = append(d.Attrs, Attr{
				Key: key,
				Value: gen.Vartype{
					Var:  "value",
					Type: "string",
				},
			})
		}

		code, err := gen.Generate(name, d, `
			package {{ lowercase .Name }}

			// {{ capitalize .Name }} struct
			// js:"{{ lowercase .Name }},omit"
			type {{ capitalize .Name }} struct {
				vdom.Child
				vdom.Node

				attrs    map[string]interface{}
				children []vdom.Child
			}

			// Props struct
			// js:"props,omit"
			type Props struct {
				attrs map[string]interface{}
			}

			// New {{ lowercase .Name }} element
			// 
			// {{ .Comment }}
			func New(props *Props, children ...vdom.Child) *{{ capitalize .Name }} {
				macro.Rewrite("$1('{{ .Name }}', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
				if props == nil {
					props = &Props{attrs: map[string]interface{}{}}
				}
				return &{{ capitalize .Name }}{
					attrs:    props.attrs,
					children: children,
				}
			}

			// Render fn
			func (s *{{ capitalize .Name }}) Render() vdom.Node {
				return s
			}

			func (s *{{ capitalize .Name }}) String() string {
				// props
				var props []string
				for key, val := range s.attrs {
					bytes, e := json.Marshal(val)
					// TODO: skip over errors?
					if e != nil {
						continue
					}
					props = append(props, key+"="+string(bytes))
				}
			
				// children
				var children []string
				for _, child := range s.children {
					children = append(children, child.Render().String())
				}
			
				if len(props) > 0 {
					return "<{{ lowercase .Name }} " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</{{ lowercase .Name }}>"
				}
			
				return "<{{ lowercase .Name }}>" + strings.Join(children, "") + "</{{ lowercase .Name }}>"
			}

			{{ range .Attrs }}
			// {{ capitalize .Key }} fn
			func {{ capitalize .Key }}({{ lowercase .Key }} string) *Props {
				macro.Rewrite("$1().Set('{{ lowercase .Key }}', $2)", macro.Runtime("Map", "Set", "JSON"), {{ lowercase .Key }})
				p := &Props{attrs: map[string]interface{}{}}
				return p.{{ capitalize .Key }}({{ lowercase .Key }})
			}

			// {{ capitalize .Key }} fn
			func (p *Props) {{ capitalize .Key }}({{ lowercase .Key }} string) *Props {
				macro.Rewrite("$_.Set('{{ lowercase .Key }}', $1)", {{ lowercase .Key }})
				p.attrs["{{ .Key }}"] = {{ lowercase .Key }}
				return p
			}
			{{ end }}

			// Attr fn
			func Attr(key string, value interface{}) *Props {
				macro.Rewrite("$1().Set($2, $3)", macro.Runtime("Map", "Set", "JSON"), key, value)
				p := &Props{attrs: map[string]interface{}{}}
				return p.Attr(key, value)
			}

			// Attr fn
			func (p *Props) Attr(key string, value interface{}) *Props {
				macro.Rewrite("$_.Set($1, $2)", key, value)
				p.attrs[key] = value
				return p
			}
		`)
		if err != nil {
			return errors.Wrapf(err, "error generating code")
		}

		dirpath := path.Join(vdomPath, d.Name)
		if err := os.MkdirAll(dirpath, 0755); err != nil {
			return errors.Wrapf(err, "error mkdir")
		}

		if err := ioutil.WriteFile(path.Join(dirpath, d.Name+".go"), []byte(code), 0644); err != nil {
			return errors.Wrapf(err, "error writefile")
		}
	}

	// finally, copy over some supporting files
	files := []string{"text.go", "vdom.go"}
	for _, file := range files {
		buf, err := ioutil.ReadFile(path.Join(dirname, "inputs", file))
		if err != nil {
			errors.Wrapf(err, "error reading %s", file)
		}

		if err := ioutil.WriteFile(path.Join(vdomPath, file), buf, 0644); err != nil {
			return errors.Wrapf(err, "error writing file %s", file)
		}
	}

	// format all the files now that they're written
	if err := gen.FormatAll(vdomPath); err != nil {
		return errors.Wrapf(err, "error formatting vdom/")
	}

	return nil
}