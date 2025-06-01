package structs

import "encoding/binary"

const (
	n                 = 312
	m                 = 156
	not_seeded        = n + 1
	upper_mask uint64 = 0xffffffff80000000
	lower_mask uint64 = 0x000000007fffffff
	matrix_a   uint64 = 0xB5026F5AA96619E9
)

type Randomizer struct {
	state []uint64
	index int
}

func NewRandomizer(seed int64) Randomizer {

	var randomizer Randomizer

	randomizer.state = make([]uint64, 312)
	randomizer.index = not_seeded

	return randomizer

}

func (randomizer *Randomizer) SetSeed(seed int64) {

	state := randomizer.state
	state[0] = uint64(seed)

	for i := uint64(1); i < n; i++ {
		state[i] = 6364136223846793005*(state[i-1]^(state[i-1]>>62)) + i
	}

	randomizer.state = state
	randomizer.index = n

}

func (randomizer *Randomizer) SetSeedFromBytes(bytes []byte) {

	var uint64_seed []uint64
	var b = 0

	for b = 0; b < len(bytes); b += 8 {
		uint64_seed = append(uint64_seed, binary.LittleEndian.Uint64(bytes[b:b+8]))
	}

	if b < len(bytes) {
		uint64_seed = append(uint64_seed, binary.LittleEndian.Uint64(bytes[b:]))
	}

	randomizer.SetSeedFromSlice(uint64_seed)

}

func (randomizer *Randomizer) SetSeedFromSlice(key []uint64) {

	randomizer.SetSeed(19650218)

	state := randomizer.state
	length := len(key)
	i := uint64(1)
	j := 0

	if n > length {
		length = n
	}

	for length > 0 {

		state[i] = (state[i] ^ ((state[i-1] ^ (state[i-1] >> 62)) * 3935559000370003845) + key[j] + uint64(j))
		i++

		if i >= n {
			state[0] = state[n-1]
			i = 1
		}

		j++

		if j >= len(key) {
			j = 0
		}

		length--

	}

	for j := uint64(0); j < n-1; j++ {

		state[i] = state[i] ^ ((state[i-1] ^ (state[i-1] >> 62)) * 2862933555777941757) - i
		i++

		if i >= n {
			state[0] = state[n-1]
			i = 1
		}

	}

	state[0] = 1 << 63

	randomizer.state = state

}

func (randomizer *Randomizer) RandomUint64() uint64 {

	state := randomizer.state

	if randomizer.index >= n {

		if randomizer.index == not_seeded {
			randomizer.SetSeed(5489)
		}

		for i := 0; i < n-m; i++ {
			y := (state[i] & upper_mask) | (state[i+1] & lower_mask)
			state[i] = state[i+m] ^ (y >> 1) ^ ((y & 1) * matrix_a)
		}

		for i := n - m; i < n-1; i++ {
			y := (state[i] & upper_mask) | (state[i+1] & lower_mask)
			state[i] = state[i+(m-n)] ^ (y >> 1) ^ ((y & 1) * matrix_a)
		}

		y := (state[n-1] & upper_mask) | (state[0] & lower_mask)
		state[n-1] = state[m-1] ^ (y >> 1) ^ ((y & 1) * matrix_a)

		randomizer.index = 0

	}

	result := state[randomizer.index]
	result ^= (result >> 29) & 0x5555555555555555
	result ^= (result << 17) & 0x71D67FFFEDA60000
	result ^= (result << 37) & 0xFFF7EEE000000000
	result ^= (result >> 43)

	randomizer.state = state
	randomizer.index++

	return result

}

func (randomizer *Randomizer) RandomClamp() float64 {
	return float64(randomizer.RandomUint64()>>11) / 9007199254740992.0
}

func (randomizer *Randomizer) Read(pseudo []byte) (int, error) {

	var length = len(pseudo)

	for len(pseudo) >= 8 {

		value := randomizer.RandomUint64()

		pseudo[0] = byte(value)
		pseudo[1] = byte(value >> 8)
		pseudo[2] = byte(value >> 16)
		pseudo[3] = byte(value >> 24)
		pseudo[4] = byte(value >> 32)
		pseudo[5] = byte(value >> 40)
		pseudo[6] = byte(value >> 48)
		pseudo[7] = byte(value >> 56)

		pseudo = pseudo[8:]

	}

	if len(pseudo) > 0 {

		value := randomizer.RandomUint64()

		for p := 0; p < len(pseudo); p++ {
			pseudo[p] = byte(value)
			value >>= 8
		}

	}

	return length, nil

}
