WASM
===

Conclusion: _WASM_ seems fully supported and is probably fine for our needs. I only have scratched the surface but in this document you'll find some distilled information. I think that if we plan to ever support the _WASM_ platform that it would be best that we try to keep pure-rust with our dependencies as it will make it much easier, in the cryptographic area that might however be a bit hard.

Get started:

* Using rustup: install rust, or make sure that you are on the latest stable version (we don’t need nightly for our purposes);
* Install the `wasm32-unknown-unknown` target for stable using `rustup`;
* Install cargo-web: `cargo install cargo-web` ( https://github.com/koute/cargo-web )

> NOTE: there are other ways to build for the web than using `cargo-web`,
> but to me `cargo-web` seems the most complete and easy solution there is at the moment, and is therefore the tool I use at the moment.

That’s it. Creating a project you would still do using `cargo init <name> --bin`,
once in there and you have code ready to be run in your browser you can serve it directly
into your browser using the `cargo web start` command
(changes to the file will make it reload your code in the browser on manual refresh).

> NOTE: while `extern create` definitions are no longer required in Rust since 2018,
> they are still required when building for the web due to dependency upon [`proc-macro2`](https://github.com/alexcrichton/proc-macro2)
>> confirmation: https://doc.rust-lang.org/nightly/edition-guide/rust-2018/module-system/path-clarity.html#an-exception

Rust allows you to interact with JS in both directions, do full manipulations on your DOM and use WebGL as well should you need it. A lot of information about this can be found at https://rustwasm.github.io/docs/wasm-bindgen/introduction.html

> NOTE:
> I noticed that resources such as <https://rustwasm.github.io/docs/book/introduction.html>
> as well as the above resource do not use `cargo-web` and instead build using other tools,
> or do it in a different way. Will need to check what the best way is, not sure about that.

Frameworks for frontend (as in DOM in WASM):
* https://github.com/DenisKolodin/yew (has most stars, have been playing with this frontend framework only for now, seems easy enough to useed

Frameworks for frontend (as in API)
* https://actix.rs (recommended?)
* https://github.com/SergioBenitez/Rocket/tree/v0.4 (lots of stars)

> NOTE:
> Make sure to have no panics,
> otherwise your web-app will no longer work until you refresh the web-page!

Open Questions:
* How to print to console? `println!` does not seem to work...
* What tooling is recommended by the core team?
