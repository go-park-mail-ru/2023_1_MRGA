// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package info_user

import (
	json "encoding/json"
	constform "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
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

func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(in *jlexer.Lexer, out *UserRestTemp) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(out *jwriter.Writer, in UserRestTemp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserRestTemp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserRestTemp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserRestTemp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserRestTemp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(in *jlexer.Lexer, out *UserRes) {
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
		case "userId":
			out.UserId = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "age":
			out.Age = int(in.Int())
		case "avatarId":
			out.Avatar = uint(in.Uint())
		case "step":
			out.Step = constform.Step(in.Int())
		case "banned":
			out.Banned = bool(in.Bool())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(out *jwriter.Writer, in UserRes) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"avatarId\":"
		out.RawString(prefix)
		out.Uint(uint(in.Avatar))
	}
	{
		const prefix string = ",\"step\":"
		out.RawString(prefix)
		out.Int(int(in.Step))
	}
	{
		const prefix string = ",\"banned\":"
		out.RawString(prefix)
		out.Bool(bool(in.Banned))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserRes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserRes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserRes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserRes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser1(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(in *jlexer.Lexer, out *StatusInp) {
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
		case "status":
			out.Status = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(out *jwriter.Writer, in StatusInp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatusInp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatusInp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatusInp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatusInp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser2(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(in *jlexer.Lexer, out *InfoStructAnswer) {
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
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "city":
			out.City = string(in.String())
		case "sex":
			out.Sex = uint(in.Uint())
		case "description":
			out.Description = string(in.String())
		case "zodiac":
			out.Zodiac = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "education":
			out.Education = string(in.String())
		case "age":
			out.Age = int(in.Int())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]uint, 0, 8)
					} else {
						out.Photos = []uint{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v1 uint
					v1 = uint(in.Uint())
					out.Photos = append(out.Photos, v1)
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(out *jwriter.Writer, in InfoStructAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.Uint(uint(in.Sex))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"zodiac\":"
		out.RawString(prefix)
		out.String(string(in.Zodiac))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		out.String(string(in.Education))
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		if in.Photos == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Photos {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Uint(uint(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InfoStructAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InfoStructAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InfoStructAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InfoStructAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser3(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(in *jlexer.Lexer, out *InfoStruct) {
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
		case "name":
			out.Name = string(in.String())
		case "city":
			out.City = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "sex":
			out.Sex = uint(in.Uint())
		case "description":
			out.Description = string(in.String())
		case "zodiac":
			out.Zodiac = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "education":
			out.Education = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(out *jwriter.Writer, in InfoStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.Uint(uint(in.Sex))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"zodiac\":"
		out.RawString(prefix)
		out.String(string(in.Zodiac))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		out.String(string(in.Education))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InfoStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InfoStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InfoStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InfoStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser4(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(in *jlexer.Lexer, out *InfoChange) {
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
		case "name":
			out.Name = string(in.String())
		case "city":
			out.City = string(in.String())
		case "sex":
			out.Sex = uint(in.Uint())
		case "description":
			out.Description = string(in.String())
		case "zodiac":
			out.Zodiac = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "education":
			out.Education = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(out *jwriter.Writer, in InfoChange) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.Uint(uint(in.Sex))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"zodiac\":"
		out.RawString(prefix)
		out.String(string(in.Zodiac))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		out.String(string(in.Education))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InfoChange) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InfoChange) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InfoChange) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InfoChange) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser5(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(in *jlexer.Lexer, out *HashtagInp) {
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
		case "hashtag":
			if in.IsNull() {
				in.Skip()
				out.Hashtag = nil
			} else {
				in.Delim('[')
				if out.Hashtag == nil {
					if !in.IsDelim(']') {
						out.Hashtag = make([]string, 0, 4)
					} else {
						out.Hashtag = []string{}
					}
				} else {
					out.Hashtag = (out.Hashtag)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Hashtag = append(out.Hashtag, v4)
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(out *jwriter.Writer, in HashtagInp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hashtag\":"
		out.RawString(prefix[1:])
		if in.Hashtag == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Hashtag {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HashtagInp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HashtagInp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HashtagInp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HashtagInp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgInfoUser6(l, v)
}
