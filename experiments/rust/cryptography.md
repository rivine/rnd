Cryptography
===

Cryptography seems the weakest point of Rust at the moment. The core team does not integrate it in the std lib, which is probably good, but at the same time it also doesn’t seem to incentivise the community to provide a good standard cryptography library (or set of libraries). Luckily we do not require a lot of different cryptographic libraries. For Rivine purposess we would only really care about ED25519 (signatures) and Blake2 (hasing) for now. Here are some alternatives.

## Libraries

### <https://github.com/sodiumoxide/sodiumoxide>

Bindings to Sodium is the most obvious choice and is also what a big blockchain player such as Maidsafe uses. It is also what the core team seems to hint at. Alex C. also seems to indicate that it is possible to use on the WASM platform with his toy project <https://github.com/alexcrichton/wasm-sodium>.

It is the more secure option, as sodium is a well trusted implementation, but it does make us much less cross platform. Mobile platforms are not supported for now, and WASM is not supported either. Only the desktop platforms are fine for now, which does already immediately raises an issue. Still, it is a valid option that we cannot ignore just like that.

> (TODO) Try and Document

### <https://github.com/dalek-cryptography/ed25519-dalek>

This is certainly a valid choice for the ED25519 part of our cryptographic needs. It is in pure rust which makes me happy, but it is a new implementation with only a tiny fraction of the auditing that went into sodium. Still, if we can use this one without sacrificing security it would be my choice at the moment.

> CONCLUSION: While it was promising, there seems to be some concerns/issues:
> * It doesn't hash+ the Rng private key input, which is concerning;
> * There are runtime issues (panics) when importing ED25519_dalek;
> * Even if we can get past this issues, it doesn't have as much of exposure (auditing) wise as for example sodium;


### <https://github.com/RustCrypto/hashes>

This is certainly a valid choice for the Blake2B (hashing) part of our cryptographic needs. It is in pure rust which makes me happy, but it is a new implementation with only a tiny fraction of the auditing that went into sodium. Still, if we can use this one without sacrificing security it would be my choice at the moment.

> CONCLUSION: Runs fine on the web as well as on MacOS, easy to use,
> and allows for variable-sized digests.
> Playground available at: https://github.com/GlenDC/crypto-playground.rs/tree/master/blake2-web

## Conclusion

I feared that cryptography would be the weakest point of Rust from earlier research,
and the current research does seem to back that up. <https://github.com/sodiumoxide/sodiumoxide>
seems to be the only mature option we have, but haven't tried it out yet. It does mean
we rely for something very important on unsafe rust, opening ourself to problems there,
as well as making us less cross-platform.
