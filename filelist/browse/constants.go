package browse

const (
	CAT_ALL       = 0
	CAT_MOV_SD    = 1
	CAT_MOV_HD    = 4
	CAT_SERIES_SD = 23
	CAT_XXX       = 7
)

var Categories = map[string]int{
	"CAT_ALL": 0,
	"CAT_MOV_SD": 1,
	"CAT_MOV_HD": 4,
	"CAT_SERIES_SD": 23,
	"CAT_XXX": 7,

}

// IsCategory ...
func IsCategory(category int) bool {
	for _,v := range Categories{
		if v == category{
			return true
		}
	}
	return false
}