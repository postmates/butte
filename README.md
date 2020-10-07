Butte
=====

Performs useless calculation to make the CPU warm / start fires. Or as a very
simple CPU benchmark. Specifically, it calculates bcrypt hashes of the word
"fire" in an endless loop.

Butte provides a Prometheus endpoint on port 2112. The
`butte_hashes_generated_total` metric can be used to gauge how fast the hashes
can be calculated.
