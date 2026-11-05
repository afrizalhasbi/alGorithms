# alGorithms

A workspace to manually learn algorithm implementations.

Doing this because I believe in first-principles and understanding fundamentals, regardless of AI agents. So here the agent write the docs, specs, test suite, and verification, while I try my best to implement each algorithm manually.

It feels pretty cool. I get to understand how seemingly opaque and mysterious arts work, and it improves my foundational understanding considerably. Not to mention how fun this is.

An algorithm is done only when both boxes are checked:
`go test .-<module> -run <Name> -v` (add `-out` for tests that produce artifacts) must pass for correctness and complexity.

## sorting

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| BubbleSort | easy | ✓ | 02-11-2026 |
| InsertionSort | easy | ✓ | 02-11-2026 |
| SelectionSort | easy | ✓ | 02-11-2026 |
| MergeSort | medium | - |
| QuickSort | medium | - |
| HeapSort | medium | ✓ | 02-11-2026 |
| RadixSort | hard | - |
| TimSort | hard | - |

## scanning

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| LinearSearch | easy | ✓ | 02-11-2026 |
| BinarySearch | easy | ✓ | 03-11-2026 |
| Kadane | medium | - |
| TwoPointerDedup | medium | ✓ | 03-11-2026 |
| BoyerMooreMajority | medium | - |
| KMP | hard | - |
| RabinKarp | hard | - |

## hashing

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| FNV1a | easy | ✓ | 03-11-2026 |
| ChainMap | medium | - | 03-11-2026 |
| OpenMap | medium | - |
| BloomFilter | hard | - |
| Ring | hard | - |

## numerics

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| KahanSum | easy | - | 04-11-2026 |
| NeumaierSum | easy | ✓ | 05-11-2026 |
| NearEqual | easy | - | 04-11-2026 |
| SafeLog-SafeExp | easy | - | 04-11-2026 |
| Horner | easy | - |
| Welford | medium | - |
| TaylorExp | medium | - |
| LogSumExp | medium | - |

## statistics

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| DiscreteExpectation | easy | - |
| BayesPosterior | easy | - |
| BinomialPMF | easy | - |
| CumulativeDistribution | medium | - |
| KDE1D | medium | - |
| EMBinomial | hard | - |
| EMMultinomial | hard | - |

## calculus_scalar

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| Bisection | easy | - |
| NewtonRoot | easy | - |
| FixedPoint | easy | - |
| GradientDescent | easy | ✓ | 05-11-2026 |
| CentralDifference | easy | - |
| SecantMethod | medium | - |
| NewtonMinimize | medium | - |
| GoldenSection | medium | - |
| MMQuartic | hard | - |

*Note* : All derivatives are hardcoded in the test suite.

## sampling


| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| InverseDiscrete | medium | - |
| SampleExponential | medium | - |
| SampleTriangular | medium | - |
| BoxMuller | medium | - |
| RejectionSample | medium | - |
| MetropolisHastings1D | hard | - |

## images

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| FloodFill | easy | - |
| Rotate90-FlipH | easy | - |
| BoxBlur | easy | - |
| BresenhamLine | medium | - |
| MidpointCircle | medium | - |
| FloydSteinbergDither | hard | - |

## audio

| Algorithm | Difficulty | Done | Date |
|---|---|---|---|
| SineGen | easy | - |
| Normalize | easy | - |
| MixClipped | easy | - |
| Echo | easy | - |
| LinearResample | medium | - |
| MovingAverageLowpass | medium | - |
