package contest

type CtstInfo struct {
	Remarks      string `json:"remarks"`
	CtstPassword string `json:"ctst_password"`
	ID           string `json:"id" binding:"required"`
}

type CtstResponse struct {
	Submissions  [][]int             `json:"submissions"`
	Participants map[string][]string `json:"participants"`
	Title        string              `json:"title"`
	Begin        int64               `json:"begin"`
	Length       int64               `json:"length"`
	ID           int                 `json:"id"`
}
