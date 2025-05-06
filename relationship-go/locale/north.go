package locale

var North = Locale{
	Name: "北方",
	Data: map[string][]string{
		"m,m":      {"姥姥"},
		"m,m,xs":   {"姨姥姥"},
		"m,m,xs,h": {"姨姥爷"},
		"m,m,xb":   {"舅姥爷"},
		"m,m,xb,w": {"舅姥姥"},
		"m,f":      {"姥爷"},
		"m,f,xs":   {"姑姥姥"},
		"m,f,xs,h": {"姑姥爷"},
		"m,f,ob":   {"大姥爷"},
		"m,f,ob,w": {"大姥姥"},
		"m,f,lb":   {"小姥爷"},
		"m,f,lb,w": {"小姥姥"},
		"f,ob":     {"大爷"},
		"f,ob,w":   {"大娘"},
	},
}
