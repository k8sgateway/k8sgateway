// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/cors/cors.proto

package cors

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *CorsPolicy) Clone() proto.Message {
	var target *CorsPolicy
	if m == nil {
		return target
	}
	target = &CorsPolicy{}

	if m.GetAllowOrigin() != nil {
		target.AllowOrigin = make([]string, len(m.GetAllowOrigin()))
		for idx, v := range m.GetAllowOrigin() {

			target.AllowOrigin[idx] = v

		}
	}

	if m.GetAllowOriginRegex() != nil {
		target.AllowOriginRegex = make([]string, len(m.GetAllowOriginRegex()))
		for idx, v := range m.GetAllowOriginRegex() {

			target.AllowOriginRegex[idx] = v

		}
	}

	if m.GetAllowMethods() != nil {
		target.AllowMethods = make([]string, len(m.GetAllowMethods()))
		for idx, v := range m.GetAllowMethods() {

			target.AllowMethods[idx] = v

		}
	}

	if m.GetAllowHeaders() != nil {
		target.AllowHeaders = make([]string, len(m.GetAllowHeaders()))
		for idx, v := range m.GetAllowHeaders() {

			target.AllowHeaders[idx] = v

		}
	}

	if m.GetExposeHeaders() != nil {
		target.ExposeHeaders = make([]string, len(m.GetExposeHeaders()))
		for idx, v := range m.GetExposeHeaders() {

			target.ExposeHeaders[idx] = v

		}
	}

	target.MaxAge = m.GetMaxAge()

	target.AllowCredentials = m.GetAllowCredentials()

	target.DisableForRoute = m.GetDisableForRoute()

	return target
}
