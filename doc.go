// Package sutil
//
// Sutil is slice util. It's spatula in Indonesian/Javanese language.
// a personal collection of slice snipplet
//
// Example:
//	input := []string{"7892141641500", "7892141600279", "7892141600422", "7892141640145", "7892141650236", "7892141650274", "7892141650311"}
//	limit := 2
//
//	pages, err := sutil.SplitStrings(input, limit)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(pages)
//
// Will return:
//
// 	[[7892141641500 7892141600279] [7892141600422 7892141640145] [7892141650236 7892141650274] [7892141650311]]
package sutil
