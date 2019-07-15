package parseme

import "io"

type Link struct{
	href string
	text string
}


func ParsePage(r io.Reader) []Link{

	a := []Link{
		{href:"http://google.com",
			text:"salam bache ha"},
			{href:"http://facebook.com",
				text:"inam link facebook"},
	}
	return a
}