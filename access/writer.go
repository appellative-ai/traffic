package access

import (
	"fmt"
	"strings"
)

const (
	markupNull   = "\"%v\":null"
	markupString = "\"%v\":\"%v\""
	markupValue  = "\"%v\":%v"
)

func writeJson(items []Operator, data *event) string {
	if len(items) == 0 || data == nil {
		return "{}"
	}
	sb := strings.Builder{}
	for _, op := range items {
		writeMarkup(&sb, op.Name, data.value(op.Value), isStringValue(op))
	}
	sb.WriteString("}")
	return sb.String()
}

func writeMarkup(sb *strings.Builder, name, value string, stringValue bool) {
	if sb.Len() == 0 {
		sb.WriteString("{")
	} else {
		sb.WriteString(",")
	}
	if len(value) == 0 {
		sb.WriteString(fmt.Sprintf(markupNull, name))
	} else {
		format := markupString
		if !stringValue {
			format = markupValue
		}
		sb.WriteString(fmt.Sprintf(format, name, value))
	}
}

func writeText(items []Operator, data *event) string {
	if len(items) == 0 || data == nil {
		return ""
	}
	sb := strings.Builder{}
	for i, op := range items {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(data.value(op.Value))
	}

	return sb.String()
}
