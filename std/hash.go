package std

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func HashTest() {
	data := "Hello World"

	// 1. MD5 哈希（128位，不推荐用于安全场景）
	fmt.Println("1. MD5 Hash:")
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))                                 // Feeds the input data (converted to bytes) into the hash.
	fmt.Printf("   MD5(\"%s\") = %x\n", data, md5Hash.Sum(nil)) // md5Hash.Sum(nil): Finalizes the hash and returns the raw 128-bit (16-byte) digest

	// 2. SHA1 哈希（160位，已被认为不够安全）
	fmt.Println("\n2. SHA1 哈希:")
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(data))
	fmt.Printf("   SHA1(\"%s\") = %x\n", data, sha1Hash.Sum(nil))

	// 3. SHA256 哈希（256位，常用且安全）
	fmt.Println("\n3. SHA256 哈希:")
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	fmt.Printf("   SHA256(\"%s\") = %x\n", data, sha256Hash.Sum(nil))

	// 4. SHA512 哈希（512位，更安全但更长）
	fmt.Println("\n4. SHA512 哈希:")
	sha512Hash := sha512.New()
	sha512Hash.Write([]byte(data))
	fmt.Printf("   SHA512(\"%s\") = %x\n", data, sha512Hash.Sum(nil))

	// 5. 使用 hex 编码展示不同的输出格式
	fmt.Println("\n5. 不同编码格式:")
	checksum := sha256.Sum256([]byte(data))
	fmt.Printf("   十六进制: %x\n", checksum)
	fmt.Printf("   十六进制(大写): %X\n", checksum)
	fmt.Printf("   hex.EncodeToString: %s\n", hex.EncodeToString(checksum[:]))

	// 6. 演示相同数据产生相同哈希
	fmt.Println("\n6. 哈希一致性验证:")
	data1 := "password123"
	data2 := "password123"
	data3 := "password124" // 只有一个字符不同

	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))
	hash3 := sha256.Sum256([]byte(data3))

	fmt.Printf("   \"%s\" -> %x\n", data1, hash1)
	fmt.Printf("   \"%s\" -> %x\n", data2, hash2)
	fmt.Printf("   \"%s\" -> %x\n", data3, hash3)
	fmt.Printf("   data1 == data2: %v (哈希相同)\n", hash1 == hash2)
	fmt.Printf("   data1 == data3: %v (哈希不同，即使只差一个字符)\n", hash1 == hash3)

	// 7. HMAC (Hash-based Message Authentication Code) - 带密钥的哈希消息认证码
	// which is used to prove both message integrity and authenticity — meaning the message was not modified and was created by someone who knows the secret key.
	// Sender and receiver both know the same secret key before any message
	// Sender:   HMAC(K, message) ──┐
	//                              ├─ compare → valid?
	// Verifier: HMAC(K, message) ──┘
	// Real-world examples
	// 1. Webhooks (GitHub, Stripe)
	// - You configure a secret once
	// - Server signs payload with HMAC
	// - Your app verifies with the same secret
	// 2. API request signing
	// - Mobile app and backend share a secret
	// - Requests are signed with HMAC
	fmt.Println("\n7. HMAC 消息认证:")
	secret := []byte("my-secret-key")
	message := "Important message"

	// 使用 HMAC-SHA256
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	fmt.Printf("   消息: \"%s\"\n", message)
	fmt.Printf("   密钥: \"%s\"\n", string(secret))
	fmt.Printf("   HMAC-SHA256: %x\n", signature)

	// 验证 HMAC
	h2 := hmac.New(sha256.New, secret)
	h2.Write([]byte(message))
	expectedSignature := h2.Sum(nil)
	equal := hmac.Equal(signature, expectedSignature)
	fmt.Printf("   正确密钥HMAC验证: %v\n", equal)

	// 使用错误的密钥验证
	wrongSecret := []byte("wrong-key")
	h3 := hmac.New(sha256.New, wrongSecret)
	h3.Write([]byte(message))
	wrongSignature := h3.Sum(nil)
	equal = hmac.Equal(signature, wrongSignature)
	fmt.Printf("   错误密钥HMAC验证: %v\n", equal)

	// 8. 文件内容哈希计算
	fmt.Println("\n8. 文件内容哈希:")
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic("无法获取当前文件路径")
	}
	dir := filepath.Dir(path)
	file := filepath.Join(dir, "go-std-hash.txt")

	// 先确保文件存在
	if err := os.WriteFile(file, []byte("File content for hashing"), 0o644); err != nil {
		fmt.Printf("   无法创建测试文件: %v\n", err)
	} else {
		fileHash, err := hashFile(file)
		if err != nil {
			fmt.Printf("   计算文件哈希失败: %v\n", err)
		} else {
			fmt.Printf("   文件: %s\n", filepath.Base(file))
			fmt.Printf("   SHA256: %s\n", fileHash)
		}
	}

	// 9. 多次写入累积哈希
	fmt.Println("\n9. 流式哈希计算（多次写入）:")
	streamHash := sha256.New()
	chunks := []string{"Hello", " ", "World", "!"}
	for i, chunk := range chunks {
		streamHash.Write([]byte(chunk))
		fmt.Printf("   写入第 %d 块: \"%s\"\n", i+1, chunk)
	}
	finalHash := streamHash.Sum(nil)
	fmt.Printf("   最终哈希: %x\n", finalHash)

	// 对比一次性哈希
	oneTimeHash := sha256.Sum256([]byte("Hello World!"))
	fmt.Printf("   一次性哈希: %x\n", oneTimeHash)
	fmt.Printf("   两种方式结果相同: %v\n", hex.EncodeToString(finalHash) == hex.EncodeToString(oneTimeHash[:]))

	fmt.Println()
}

// hashFile 计算文件的 SHA256 哈希值
// File ──> io.Copy ──> SHA256 machine ──> final digest
func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	// How does io.Copy(hash, file) works?
	// Feed these file content in bytes into the hash algorithm
	// - Streams the entire file (No need to load the whole file into memory)
	// - Chunk by chunk (efficient for large files)
	// - Each chunk updates the SHA-256 internal state
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
