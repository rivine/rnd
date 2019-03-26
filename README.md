# rnd

Research &amp; Development 

## Prototypes

### TFChain Block Explorer

A small Go+Fyne prototype to see how easy it is develop a Pure Go GUI App.
Conclusion is that it is really easy.

The prototype is a simple block explorer, showing the current block
height, ID and time for both standard and test network.
It is not an advanced application, more like a test.

Dependencies can be get as follows:

```
go get -u github.com/threefoldtech/rivine
go get -u fyne.io/fyne
```

The prototype can be run as follows:

```
go run prototypes/tfchain-block-explorer/main.go
```

Screenshot:

![screenshot of TFChain Block Explorer prototype](/assets/screenshots/tfchain-block-explorer.png)

Prototype isn't super advanced in features, but does show that it is really simple to build a Pure Go GUI App.
More discipline, in architecture, is required for a more serious App though.

Fyne showcases a lot more examples on how to do stuff at <https://github.com/fyne-io/examples/>.

## Experiments

### Rust

You can find some documents under [/experiments/rust](/experiments/rust) that describe the Rust-focussed experiments that have taken place.
