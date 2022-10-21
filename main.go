package main

import (
	//"bufio"
	"errors"
	"fmt"
	"html"

	//"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strings"
	//"myProj/xsd"
	//"myProj/xmlfmt"
)

//////xmlfmt//////
var (
	reg = regexp.MustCompile(`<([/!]?)([^>]+?)(/?)>`)
	// NL is the newline string used in XML output.
	NL = "\n"
)

func init() {
	// define NL for Windows
	if runtime.GOOS == "windows" {
		NL = "\r\n"
	}
}

// FormatXML will (purly) reformat the XML string in a readable way, without any rewriting/altering the structure.
// If your XML Comments have nested tags in them, or you're not 100% sure otherwise, pass `true` as the third parameter to this function. But don't turn it on blindly, as the code has become ten times more complicated because of it.
func FormatXML(xmls, prefix, indent string, nestedTagsInComments ...bool) string {
	nestedTagsInComment := false
	if len(nestedTagsInComments) > 0 {
		nestedTagsInComment = nestedTagsInComments[0]
	}
	reXmlComments := regexp.MustCompile(`(?s)(<!--)(.*?)(-->)`)
	src := regexp.MustCompile(`(?s)>\s+<`).ReplaceAllString(xmls, "><")
	if nestedTagsInComment {
		src = reXmlComments.ReplaceAllStringFunc(src, func(m string) string {
			parts := reXmlComments.FindStringSubmatch(m)
			p2 := regexp.MustCompile(`\r*\n`).ReplaceAllString(parts[2], " ")
			return parts[1] + html.EscapeString(p2) + parts[3]
		})
	}
	rf := replaceTag(prefix, indent)
	r := prefix + reg.ReplaceAllStringFunc(src, rf)
	if nestedTagsInComment {
		r = reXmlComments.ReplaceAllStringFunc(r, func(m string) string {
			parts := reXmlComments.FindStringSubmatch(m)
			return parts[1] + html.UnescapeString(parts[2]) + parts[3]
		})
	}

	return r
}

// replaceTag returns a closure function to do 's/(?<=>)\s+(?=<)//g; s(<(/?)([^>]+?)(/?)>)($indent+=$3?0:$1?-1:1;"<$1$2$3>"."\n".("  "x$indent))ge' as in Perl
// and deal with comments as well
func replaceTag(prefix, indent string) func(string) string {
	indentLevel := 0
	return func(m string) string {
		// head elem
		if strings.HasPrefix(m, "<?xml") {
			return NL + prefix + strings.Repeat(indent, indentLevel) + m
		}
		// empty elem
		if strings.HasSuffix(m, "/>") {
			return NL + prefix + strings.Repeat(indent, indentLevel) + m
		}
		// comment elem
		if strings.HasPrefix(m, "<!") {
			return NL + prefix + strings.Repeat(indent, indentLevel) + m
		}
		// end elem
		if strings.HasPrefix(m, "</") {
			indentLevel--
			return NL + prefix + strings.Repeat(indent, indentLevel) + m
		}
		defer func() {
			indentLevel++
		}()

		return NL + prefix + strings.Repeat(indent, indentLevel) + m
	}
}

/////end xmlfmt///

func main() {
	/*
	   	xml1 := `<root><this><is>a</is><test /><message><!-- with comment --><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`
	   	x := xmlfmt.FormatXML(xml1, "\t", "  ")
	   	print(x)

	   	// If the XML Comments have nested tags in them
	   	xml1 = `<book> <author>Fred</author>
	   <!--
	   <price>20</price><currency>USD</currency>
	   -->
	    <isbn>23456</isbn> </book>`
	   	x = xmlfmt.FormatXML(xml1, "", "  ", true)
	   	print(x)
	*/

	infile, err := os.Open("build.txt")
	//var xml1 string
	if err != nil {
		err = errors.New("reading file error")
		fmt.Println(err)
		return
	}
	outfile, err := os.Create("exiBuild.txt")
	if err != nil {
		err = errors.New("out file error")
		fmt.Println(err)
		return
	}
	defer infile.Close()
	defer outfile.Close()

	/*
		infle := bufio.NewScanner(infile)
		for infle.Scan() {
			xml1 = xml1 + infle.Text()

			//fmt.Println(infle.Text())
		}
		x := FormatXML(xml1, "\t", "  ")
		print(x)
		fmt.Println("-------------------------------")
		y, _ := ioutil.ReadFile("build.txt")
		fmt.Println(string(y))
		//fmt.Printf("%T", infile)
		//fmt.Println(xml1)
	*/
}
