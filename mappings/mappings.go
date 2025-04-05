package mappings

var TypeMappings = map[string]map[string]string{
	"cpp": {
		"integer":    "int",
		"string":     "string",
		"array":      "vector<${subtype}>",
		"linkedlist": "ListNode*",
		"char":       "char",
		"boolean":    "bool",
		"float":      "float",
		"double":     "double",
	},
	"python": {
		"integer":    "int",
		"string":     "str",
		"array":      "List[${subtype}]",
		"linkedlist": "ListNode",
		"char":       "str",
		"boolean":    "bool",
		"float":      "float",
		"double":     "float",
	},
	"java": {
		"integer":    "int",
		"string":     "String",
		"array":      "${subtype}[]",
		"linkedlist": "ListNode",
		"char":       "char",
		"boolean":    "boolean",
		"float":      "float",
		"double":     "double",
	},
}
