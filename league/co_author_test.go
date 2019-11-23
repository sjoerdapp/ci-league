package league

import "testing"

func TestExtractCoAuthor(t *testing.T) {
	tests := []struct {
		name string
		commit string
		want string
	}{
		{
			"With co-author",
			`Remove default setting of anchor by useActiveNavItem

Co-authored-by: LisaMcCormack <lisamccormack85@gmail.com>`,
			"lisamccormack85@gmail.com",
		},
		{
			"Without co-author",
			`refactoring`,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractCoAuthor(tt.commit); got != tt.want {
				t.Errorf("extractCoAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}