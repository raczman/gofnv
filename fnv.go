package fnv

var prime32 = uint32(16777619)
var offset32 = uint32(2166136261)
var prime64 = uint64(1099511628211)
var offset64 = uint64(14695981039346656037)

// Calculates 64bit FNV1 hash of byte array k
func Hash1(k []byte) uint64 {
	offset := offset64

	for i := 0; i < len(k); i++ {
		offset = offset * prime64
		offset = offset ^ uint64(k[i])
	}
	return offset
}

// Calculates 32bit FNV1 hash of byte array k
func Hash1_32(k []byte) uint32 {
	offset := offset32

	for i := 0; i < len(k); i++ {
		offset = offset * prime32
		offset = offset ^ uint32(k[i])
	}
	return offset
}

// Calculates 64bit FNV1a hash of byte array k
func Hash1a(k []byte) uint64 {
	offset := offset64

	for i := 0; i < len(k); i++ {
		offset = offset ^ uint64(k[i])
		offset = offset * prime64
	}
	return offset
}

// Calculates 32bit FNV1a hash of byte array k
func Hash1a_32(k []byte) uint32 {
	offset := offset32

	for i := 0; i < len(k); i++ {
		offset = offset ^ uint32(k[i])
		offset = offset * prime32
	}
	return offset
}

func hashMod(hash uint64, size uint64) uint64 {
	size += 1
	retry_level := (uint64(0xffffffffffffffff) / size) * size
	for hash >= retry_level {
		hash = (hash * prime64) + offset64
	}
	return hash % size
}

func hashMod32(hash uint32, size uint32) uint32 {
	size += 1
	retry_level := (uint32(0xffffffff) / size) * size
	for hash >= retry_level {
		hash = (hash * prime32) + offset32
	}
	return hash % size
}

// Calculates 32bit FNV1a hash modulo
func Hash1aMod_32(k []byte, size uint32) uint32 {
	hash := Hash1a_32(k)
	return hashMod32(hash, size)
}

// Calculates 32bit FNV1a hash modulo
func Hash1Mod_32(k []byte, size uint32) uint32 {
	hash := Hash1_32(k)
	return hashMod32(hash, size)
}

// Calculates 64bit FNV1a hash modulo
func Hash1aMod(k []byte, size uint64) uint64 {
	hash := Hash1a(k)
	return hashMod(hash, size)
}

// Calculates 64bit FNV1 hash modulo
func Hash1Mod(k []byte, size uint64) uint64 {
	hash := Hash1(k)
	return hashMod(hash, size)
}