// +build cortexm3 cortexm4 cortexm4f cortexm7f cortexm7d

inline __attribute__((always_inline))
uint
bits$leadingZeros32(uint32 u) {
	uint n;
	asm volatile ("clz %0, %1":"=r" (n):"r"(u));
	return n;
}

inline __attribute__((always_inline))
uint
bits$leadingZerosPtr(uintptr u) {
	uint n;
	asm volatile ("clz %0, %1":"=r" (n):"r"(u));
	return n;
}

inline __attribute__((always_inline))
uint
bits$leadingZeros64(uint64 u) {
	uint n;
	asm volatile (
		"clz %0, %R1\n\t"
		"cmp %0, 32\n\t"
		"itt eq\n\t"
		"clzeq %0, %Q1\n\t"
		"addeq %0, 32\n\t"
		:"=&r" (n)
		:"r"(u)
		:"cc"
	);
	return n;
}
