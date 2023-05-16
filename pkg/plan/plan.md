# 学习计划 - {{ htmlDate .Date }}

## 预览

{{ range $idx, $word := .Words }}
{{ add1 $idx }}. {{$word.Name}}
{{- end}}

{{ range $idx, $word := .Words }}

## {{ add1 $idx }}. {{ .Name }}

### 读音拼写

音标：/ {{ .Phonetic }} /

释义：{{ .Meaning }}

### 记忆技巧

{{ $length := len .Roots }} {{ if gt $length 0 }}
词根：
{{ range .Roots }}
{{.}}
{{ end }}
{{ end }}

{{ .RootGraph }}

{{ .WordsGraph }}

### 语境语法

例句：
{{range .Examples}}
{{.}}
{{end}}

{{end}}
