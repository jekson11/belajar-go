-- FindAllUserData
SELECT user_id, name, username, email, created_at
FROM learngo.td_user
WHERE 1=1
{{ if .Name }}
    AND LOWER(name) LIKE {{ .Name }}
{{ end }}
{{ if .Email }}
    AND email = {{ .Email }}
{{ end }}
{{ if .Username }}
    AND username = {{ .Username }}
{{ end }}
__ORDER_BY__
LIMIT {{ .Limit }}
OFFSET {{ .Page }}

-- CountDataUser
SELECT COUNT(user_id)
FROM learngo.td_user
WHERE 1=1
    {{ if .Name }}
    AND LOWER(name) LIKE {{ .Name }}
{{ end }}
{{ if .Email }}
    AND email = {{ .Email }}
{{ end }}
{{ if .Username }}
    AND username = {{ .Username }}
{{ end }}

-- CreateUser
INSERT INTO learngo.td_user (name, username, password, email)
VALUES
{{- range $i, $v := . }}
    {{- if $i}},{{ end }}
    ({{$v.Name}}, {{$v.Username}}, {{$v.Password}}, {{$v.Email}})
{{- end }}