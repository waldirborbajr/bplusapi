package services

// time for db process any transaction
// const dbTimeout = time.Second * 3

// var db *sql.DB

// // create a new db pool to cache data
// func New(dbPool *sql.DB) Models {
// 	db = dbPool
// 	return Models{
// 		User: User{},
// 	}
// }

type Model struct {
	// User         User
	JsonResponse JsonResponse
	// Report        Report
	// Field         Field
	// Profile       Profile
	// Member        Member
	// Group         Group
	TokenResponse TokenResponse
	// Game          Game
}
