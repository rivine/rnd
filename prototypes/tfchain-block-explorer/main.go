package main

import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"encoding/json"

	"github.com/threefoldtech/rivine/pkg/api"

	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

type Context struct {
	Network string
	NetworkRadio *widget.Radio

	UpdateCh chan struct{}

	BlockHashLabel		*widget.Label
	BlockHeightLabel	*widget.Label
	BlockTimeLabel		*widget.Label
}

const (
	networkStandard = "standard"
	networkTestnet = "testnet"
)

func (ctx *Context) SetNetwork(s string) {
	if s == "" {
		ctx.NetworkRadio.Selected = ctx.Network
		widget.Renderer(ctx.NetworkRadio).Refresh()
		return
	}
	if s != networkStandard && s != networkTestnet {
		panic(fmt.Sprintf("invalid network: %s", s))
	}
	ctx.Network = s
	ctx.UpdateCh <- struct{}{}
}

func NewRadioWithSelected(options []string, selected string, f func(s string)) *widget.Radio {
	r := new(widget.Radio)
	r.Options = options
	r.Selected = selected
	r.OnChanged = f
	widget.Renderer(r).Layout(r.MinSize())
	widget.Renderer(r).Refresh()
	return r
}

func main() {
	app := app.New()

	ctx := &Context{
		Network: networkStandard,
		UpdateCh: make(chan struct{}),
	}
	ctx.NetworkRadio = NewRadioWithSelected(
		[]string{networkTestnet, networkStandard}, ctx.Network, ctx.SetNetwork)
	ctx.BlockHashLabel = widget.NewLabel("?")
	ctx.BlockHeightLabel = widget.NewLabel(strings.Repeat("?", 80))
	ctx.BlockTimeLabel = widget.NewLabel("??? ???")

	w := app.NewWindow("TFChain Chain Explorer")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Threefold Chain Explorer"),
		widget.NewHBox(
			widget.NewLabel("Network:"),
			ctx.NetworkRadio,
		),
		widget.NewVBox(
			widget.NewHBox(
				widget.NewLabel("Current Block Height:"),
				ctx.BlockHeightLabel,
			),
			widget.NewHBox(
				widget.NewLabel("Current Block:"),
				ctx.BlockHashLabel,
			),
			widget.NewHBox(
				widget.NewLabel("Current Chain Time:"),
				ctx.BlockTimeLabel,
			),
		),
	))

	// spawn refresh logic
	go func() {
		refreshCh := time.Tick(30 * time.Second)
		for {
			var chainInfo api.ExplorerGET
			func() {
				url := "https://explorer.threefoldtoken.com"
				if ctx.Network == networkTestnet {
					url = "https://explorer.testnet.threefoldtoken.com"
				}
				url += "/explorer"
				resp, err := http.Get(url)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
				}
				err = json.NewDecoder(resp.Body).Decode(&chainInfo)
				if err != nil {
					panic(err)
				}
			}()
			ctx.BlockHashLabel.SetText(chainInfo.BlockID.String())
			ctx.BlockHeightLabel.SetText(fmt.Sprintf("%d", chainInfo.Height))

			var blockGET api.ExplorerBlockGET
			func() {
				url := "https://explorer.threefoldtoken.com"
				if ctx.Network == networkTestnet {
					url = "https://explorer.testnet.threefoldtoken.com"
				}
				url += "/explorer/hashes/" + chainInfo.BlockID.String()
				resp, err := http.Get(url)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
				}
				err = json.NewDecoder(resp.Body).Decode(&blockGET)
				if err != nil {
					panic(err)
				}
			}()
			ctx.BlockTimeLabel.SetText(
				fmt.Sprint(time.Unix(int64(blockGET.Block.RawBlock.Timestamp), 0)))

			select {
			case <-refreshCh:
			case <-ctx.UpdateCh:
			}
		}
	}()

	w.ShowAndRun()
}
