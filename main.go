// fortune is a stripped-down implementation of the classic BSD Unix
// fortune command. Unlike the BSD fortune command (or my own Python version,
// at https://github.com/bmc/fortune), this version does not use an index file.
// We have loads of memory these days, and fortunes files aren't that big, so
// it's feasible to load the whole text file in memory, parse it on the fly,
// and randomly choose a resulting fortune.
//
// See the accompanying README for more information.

package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"git.sr.ht/~adnano/go-gemini"
	"github.com/qpliu/qrencode-go/qrencode"
	"github.com/t-900-a/rss"
)

const VERSION = "1.0"

// ---------------------------------------------------------------------------
// Local error object that conforms to the Go error interface.
// ---------------------------------------------------------------------------

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// ---------------------------------------------------------------------------
// Local functions
// ---------------------------------------------------------------------------

// Convenience function to print a message (printf style) to standard error
// and exit with a non-zero exit code.
func die(format string, args ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, args...) + "\n")
	os.Exit(1)
}

// Given a path representing a fortune file, load the file, parse it,
// an return an array of fortune strings.
func readFortuneFile(fortuneFile string) ([]string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "\n%\n")
	}
	return fortunes, err
}

// Given a path representing a fortune file, load the file, parse it,
// choose a random fortune, and display it to standard output. Returns
// a Go error object on error or nil on success.
func getFortune(fortuneFile string) (string, error) {
	fortunes, err := readFortuneFile(fortuneFile)
	if err == nil {
		rand.Seed(time.Now().UTC().UnixNano())
		i := rand.Int() % len(fortunes)
		return fortunes[i], nil
	}
	return "", err
}

func getEmoji() string {
	var buf bytes.Buffer
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Int() % len(emojis)
	buf.WriteRune(emojis[i])
	return buf.String()
}

// Parse the command line arguments. For now, this is simple, because this
// program requires very few arguments. If something more complicated is
// needed, consider the Go "flag" module or github.com/docopt/docopt-go
func parseArgs() (string, string, string, string, string, error) {
	prog := path.Base(os.Args[0])
	usage := fmt.Sprintf(`%s, version %s

Usage:
  %s [/path/to/fortune/cookie/file]
  %s [payment uri]
  %s [tx hash]
  %s -h|--help

If the fortune cookie file path is omitted, the contents of environment
variable FORTUNE_FILE will be used. If neither is available, fortune will
abort.`, prog, VERSION, prog, prog)
	// TODO REDO variable handling
	var fortuneFile string
	var websiteUri string
	var txHash string
	var pmtUri string
	var pmtViewKey string
	var err error
	switch len(os.Args) {
	case 2:
		fortuneFile = os.Args[1]
	case 3:
		fortuneFile = os.Args[1]
		websiteUri = os.Args[2]
	case 4:
		fortuneFile = os.Args[1]
		websiteUri = os.Args[2]
		txHash = os.Args[3]
	case 5:
		fortuneFile = os.Args[1]
		websiteUri = os.Args[2]
		txHash = os.Args[3]
		pmtUri = os.Args[4]
	case 6:
		fortuneFile = os.Args[1]
		websiteUri = os.Args[2]
		txHash = os.Args[3]
		pmtUri = os.Args[4]
		pmtViewKey = os.Args[5]
	case 7:
		{
			if (os.Args[1] == "-h") || (os.Args[1] == "--help") {
				err = &Error{usage}
			} else {
				fortuneFile = os.Args[1]
			}
		}
	default:
		err = &Error{usage}
	}

	if (err == nil) && (fortuneFile == "") {
		err = &Error{"No fortunes parameter and no FORTUNE_FILE " +
			"environment variable"}
	}

	return fortuneFile, websiteUri, txHash, pmtUri, pmtViewKey, err
}

// ---------------------------------------------------------------------------
// Main program
// ---------------------------------------------------------------------------

func main() {
	ti := time.Now()
	fortuneFile, websiteUri, txHash, pmtUri, pmtViewKey, err := parseArgs()
	if err != nil {
		die(err.Error())
	}

	fortune, err := getFortune(fortuneFile)
	if err != nil {
		die(err.Error())
	}

	// write fortune to gemini text
	var t gemini.Text
	title := getEmoji() + " New Fortune " + getEmoji()
	heading := gemini.LineHeading1(title)
	t = append(t, heading)

	newLine := gemini.LineText("\n")
	t = append(t, newLine)

	if len(txHash) > 0 {
		link := "https://xmrchain.net/tx/" + txHash
		txt := "Fortune requested, associated transaction found here"
		blockExplorer := &gemini.LineLink{URL: link, Name: txt}
		t = append(t, blockExplorer)
	}

	newLine = gemini.LineText("\n")
	t = append(t, newLine)
	body := "```" + botsay(fortune) + "```"
	scanner := bufio.NewScanner(strings.NewReader(body))
	for scanner.Scan() {
		fortuneLine := gemini.LineText(scanner.Text())
		t = append(t, fortuneLine)
	}

	dateLine := gemini.LineQuote(ti.Format("2006-01-02T15:04:05Z"))
	t = append(t, dateLine)

	newLine = gemini.LineText("\n")
	t = append(t, newLine)

	if len(pmtUri) > 0 {
		txt := "Curious about today's fortune? Send any amount of Monero and you will get your very own."
		pmtAddress := &gemini.LineLink{URL: pmtUri, Name: txt}
		t = append(t, pmtAddress)
		// generate qr code to scan
		// only works for gemini cli clients

		newLine = gemini.LineText("\n")
		t = append(t, newLine)

		grid, err := qrencode.Encode(pmtUri, qrencode.ECLevelL)
		if err != nil {
			log.Fatal(err)
		}
		var b bytes.Buffer
		grid.TerminalOutput(&b)
		qrCode := "```" + b.String() + "```"
		scanner := bufio.NewScanner(strings.NewReader(qrCode))
		for scanner.Scan() {
			fortuneLine := gemini.LineText(scanner.Text())
			t = append(t, fortuneLine)
		}

		newLine = gemini.LineText("\n")
		t = append(t, newLine)

		txt = "Don't have any Monero? Learn more here"
		link := "https://getmonero.org/"
		info := &gemini.LineLink{URL: link, Name: txt}
		t = append(t, info)
	}

	fileName := "fortune_" + ti.Format("2006_01_02_15_04_05") + ".gmi"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range t {
		file.WriteString(line.String() + "\n")
	}

	// write ATOM FEED
	// TODO read and append to existing feed if there
	// payment data
	var pmtData []rss.AtomLink
	pmtLink := rss.AtomLink{
		Href: pmtUri,
		Rel:  "payment",
		Type: "application/monero-paymentrequest",
	}
	pmtData = append(pmtData, pmtLink)
	pmtLink = rss.AtomLink{
		Href: pmtViewKey,
		Rel:  "payment",
		Type: "application/monero-viewkey",
	}
	pmtData = append(pmtData, pmtLink)
	// feed website
	var websiteLink []rss.AtomLink
	wbLink := rss.AtomLink{
		Href: websiteUri,
	}
	websiteLink = append(websiteLink, wbLink)
	// this item
	var itemLink []rss.AtomLink
	thisLink := websiteUri + "/gemlog/" + fileName
	itmLink := rss.AtomLink{
		Href: thisLink,
	}
	itemLink = append(itemLink, itmLink)
	// feed items
	var feedItems []rss.AtomItem
	feedItem := rss.AtomItem{
		Title:   title,
		Content: rss.RAWContent{RAWContent: "Anon was provided with their fortune, ready for yours?"},
		Links:   itemLink,
		Date:    ti.Format("2006-01-02T15:04:05Z"),
	}
	feedItems = append(feedItems, feedItem)

	feed := &rss.AtomFeed{
		Title:       "Fortune Bot",
		Description: "Your seer in the Gemini Space",
		Author: rss.AtomAuthor{
			Name:       "Anon",
			URI:        websiteUri,
			Extensions: pmtData,
		},
		Link:    websiteLink,
		Items:   feedItems,
		Updated: ti.Format("2006-01-02T15:04:05Z"),
	}
	atomFile, err := xml.MarshalIndent(feed, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("atom.xml", atomFile, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
