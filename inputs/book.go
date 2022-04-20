package inputs

type CreateBook struct {
	Title       string `json:"title"        binding:"required"`
	Author      string `json:"author"       binding:"required"`
	Published   string `json:"published"    binding:"required"`
	IsCompleted bool   `json:"is_completed"`
}

type SearchBook struct {
	Search string `json:"search" binding:"required"`
}
