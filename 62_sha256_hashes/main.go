// SHA256 Hashes
//
// SHA256 hashes are frequently used to compute short identities for binary or text
// blobs. For example, TLS/SSL certificates use SHA256 to compute a certificate's
// signature. Here's how to compute SHA256 hashes in Go.
//
// Run: go run ./62_sha256_hashes
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash/crc32"
	"io"
	"strings"
)

// Go implements several hash functions in various crypto/* packages.
func main() {
	s := "sha256 this string"

	// Here we start with a new hash.
	h := sha256.New()

	// Write expects bytes. If you have a string s, use []byte(s) to coerce it to
	// bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. The argument to Sum can be
	// used to append to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// For data already in memory, sha256.Sum256 is a one-liner.
	fmt.Printf("%x\n", sha256.Sum256([]byte(s)))

	// A hash is an io.Writer, so io.Copy can hash a stream, such as a file, without
	// loading it into memory.
	h.Reset()
	io.Copy(h, strings.NewReader(s))
	fmt.Printf("streamed matches: %t\n", fmt.Sprintf("%x", h.Sum(nil)) == fmt.Sprintf("%x", bs))

	// An HMAC signs a message with a secret key. Compare MACs with hmac.Equal, which
	// runs in constant time.
	mac := hmac.New(sha256.New, []byte("secret-key"))
	mac.Write([]byte("message"))
	sig := mac.Sum(nil)
	fmt.Printf("hmac: %x\n", sig)
	fmt.Println("valid:", hmac.Equal(sig, sig))

	// Other hashes work the same way. MD5 is broken for security but still used for
	// checksums, and CRC32 is a fast non-cryptographic checksum.
	fmt.Printf("md5: %x\n", md5.Sum([]byte(s)))
	fmt.Printf("crc32: %08x\n", crc32.ChecksumIEEE([]byte(s)))
}
