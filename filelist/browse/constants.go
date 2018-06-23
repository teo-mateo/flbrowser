package browse

const (
	CAT_ALL       = 0
	CAT_MOV_SD    = 1
	CAT_MOV_HD    = 4
	CAT_SERIES_SD = 23
	CAT_XXX       = 7
	CAT_AUDIO	  = 11
	CAT_DOCS      =16
)

var Categories = map[string]Category{
	"CAT_ALL": {0, "All"},
	"CAT_MOV_SD": {1, "Movies - SD"},
	"CAT_MOV_HD": {4, "Movies - HD"},
	"CAT_SERIES_SD": {23, "Series - SD"},
	"CAT_XXX": {7, "XXX"},
	"CAT_AUDIO": {11, "Audio"},
	"CAT_DOCS": {16, "Docs"},
}

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// IsCategory ...
func IsCategory(category int) bool {
	for _,v := range Categories{
		if v.ID == category{
			return true
		}
	}
	return false
}