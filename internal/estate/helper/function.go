package helper

import "github.com/SawitProRecruitment/UserService/generated"

func GetTreeHeight(treeHeights map[[2]int]int, x, y int) int {
	if height, exists := treeHeights[[2]int{x, y}]; exists {
		return height
	}
	return 0
}

func CheckMaxDistance(totalDistance float32, maxDistance *int, landingX, landingY int) (*generated.GetEstateDronePlanResponse, bool) {
	if maxDistance != nil && totalDistance >= float32(*maxDistance) {
		resp := &generated.GetEstateDronePlanResponse{
			Distance: &totalDistance,
			Rest: &struct {
				X *int `json:"x,omitempty"`
				Y *int `json:"y,omitempty"`
			}{
				X: &landingX,
				Y: &landingY,
			},
		}
		return resp, true
	}
	return nil, false
}
