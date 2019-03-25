TFChain Desktop App
===

Currently we do not have a user-friendly application that allows the managements
of several TFChain wallets, be it single signature or multi signature.

- The TFChain Mobile Wallet comes close, but misses support for multi-signature wallets;
- The TFChain Python client, integrated in JSX, does it all, is a light client, but has no GUI;
- The TFChain CLI App does it all as well, but has no GUI;
- The TFChain CLI Light App does a ot, but has no GUI;

The fact that we have no complete (light-client) wallet solution, controlled by a GUI,
means that a typical end-user has currently no means to utilize the full feature set that TFChain has to offer.

## Implementations

### Go

Personally I would like to do it in Go as it is the language our team is most comfortable with at the moment,
and it would allow us to deliver something within days. All blockchain specific code is also ready as
Rivine was written in Go. And as TFChain is a Rivine-go based blockchain, so is TFChain.

If we go for something like <https://github.com/fyne-io/fyne> we can immediately focus on the
wallet logic, as it is all native in Go (does use CGO) and seems more than good enough for our needs.

Another option would be to use electron (via astilectron), but as our team is not as much of
an expert in web technology, it might be a less time-efficient solution, even though it is a more powerful one.

### Python

Given we do already have a Light Client written in Python, the one that is integrated in JSX. It would
only be a small effort to copy it over into its own repo, decouple it from JSX and have a feature-complete
light client ready for use in a Python-based backend.

There are plenty of solutions available, but time-wise it would cost us a lot of time as a team to go for this
road, given that we are not that experienced with python, and the solutions are a lot more hardcore in usage.
The more easy solutions go also for a web-based frontend, but if we go for that road, we can just as well
stick with Go.

### Rust

Doing it in Rust would be the least time-efficient solution, as we would have to start from scratch.
However it would give our team some working experience with Rust, and allowing us to productively
experiment further with WASM.


## Conclusion

If we want to deliver a Desktop App within days, Go (+Fyne) is the way to go.
When time plays no role in our decision, and we are not afraid from sacrifising time
to pay for a lot of valuable learning experiences than Rust (+WASM) could be a nice way to go forward.
If the latter it would effectively be a pure and native. Web-Application.