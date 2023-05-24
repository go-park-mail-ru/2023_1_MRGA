// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package filter

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(in *jlexer.Lexer, out *FilterInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "minAge":
			out.MinAge = int(in.Int())
		case "maxAge":
			out.MaxAge = int(in.Int())
		case "sexSearch":
			out.SearchSex = uint(in.Uint())
		case "reason":
			if in.IsNull() {
				in.Skip()
				out.Reason = nil
			} else {
				in.Delim('[')
				if out.Reason == nil {
					if !in.IsDelim(']') {
						out.Reason = make([]string, 0, 4)
					} else {
						out.Reason = []string{}
					}
				} else {
					out.Reason = (out.Reason)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Reason = append(out.Reason, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(out *jwriter.Writer, in FilterInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"minAge\":"
		out.RawString(prefix[1:])
		out.Int(int(in.MinAge))
	}
	{
		const prefix string = ",\"maxAge\":"
		out.RawString(prefix)
		out.Int(int(in.MaxAge))
	}
	{
		const prefix string = ",\"sexSearch\":"
		out.RawString(prefix)
		out.Uint(uint(in.SearchSex))
	}
	{
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		if in.Reason == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Reason {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilterInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilterInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilterInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilterInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgFilter(l, v)
}
