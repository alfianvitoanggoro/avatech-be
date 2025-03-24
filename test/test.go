package test

func Test() {
	// Test change color text
	// c := NewDebug("Hello world")
	// c.Color()

	// Test get toml document
	// Toml()

	// Test get yaml document
	// Yaml()

	// Test use viper for read document
	// Viper()

	// Test use cron job
	// CronJob()

	// Test user cron job without library
	// CronJobWithoutLibrary()

	// Test use cobra
	// Cobra()

	// Test use resty
	r := NewResty()
	// r.GetData()
	// r.GetDataWithPathParam(100)
	// Create Post
	// newPost := Post{
	// 	UserID: 1,
	// 	Title:  "Belajar Go dengan Resty",
	// 	Body:   "Ini adalah contoh request POST dengan Go dan Resty.",
	// }

	// r.CreatePost(newPost)

	// Update Post
	updatePost := Post{
		UserID: 1,
		Title:  "Belajar Go dengan Resty Update",
		Body:   "Ini adalah contoh request POST dengan Go dan Resty Update.",
	}
	r.UpdatePost(100, updatePost)
}
